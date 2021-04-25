import os
import sys
from clients import Clients
from security import SecuritySchemes
from schema import Schema
from run_generator import Generator
from auth_request import AuthRequest


class Run:
    """
    Run contains the functionality to parse the provided variables, and run the actual module
    as described in the README
    """

    def __init__(self, schema='', clients='', api_url='', exhaustive=False):
        temp_file = self._write_temporary_file(schema=schema)
        self._schema = Schema(temp_file, api_url)
        self._security_schemes = self._schema.security_schemes
        parsed_clients = Clients(clients_list=clients)
        self._runs = Generator(
            clients=parsed_clients.clients,
            exhaustive=exhaustive).generate()
        self._exahustive = exhaustive
        self._api_url = api_url

    def _write_temporary_file(self, schema=''):
        name = 'temporary_schema_file.yaml'
        with open(name, 'w') as temp:
            temp.write(schema)
        return name

    def run(self):
        """
        Start the module and fuzz the api
        """
        print(
            'Starting oauth2_scopes fuzzer on exhaustive mode: {}'.format(
                self._exahustive))
        print('Using  base url {}'.format(self._api_url))
        print(
            'Using security schemes: {}'.format(
                self._security_schemes.schemes))
        print('Starting')
        oauth_session = AuthRequest(self._security_schemes)
        for run in self._runs:
            print('Starting run with client id: {}'.format(run.id))
            print('Using available client scopes {}'.format(run.scopes))
            for path in self._schema.paths:
                url = path.generate_path()
                body = None
                if path.body_required:
                    body = path.generate_request_body(
                        application_type=path.application_types[0])
                print(
                    'Testing path: {} with generated url {} method {} body {}'.format(
                        path.path, url, path.request_method, body))
                for security in path.security.keys():
                    print(
                        'Using security scheme {} requiring scopes: {}'.format(
                            security, path.security[security]))
                    try:
                        request = oauth_session.create_request(run, security)
                        print('Acquired token: {}'.format(request[0]))
                        request_transport = request[1]
                        response = None
                        if path.request_method == 'get':
                            response = request_transport.get(url)
                        elif path.request_method == 'post':
                            response = request_transport.post(url, data=body)
                        elif path.request_method == 'delete':
                            response = request_transport.delete(url)
                        elif path.request_method == 'patch':
                            response = request_transport.patch(url, data=body)
                        print('response code: {}'.format(response.status_code))
                        if self._validate_response(
                                response, run.scopes, path.security[security]):
                            print('Response is authorized')
                        else:
                            print('Response should not be authorized')
                    except Exception as e:
                        print(e)

    def _validate_response(self, response, client_scopes, required_scopes):
        required_set = set(required_scopes)
        client_set = set(client_scopes)
        should_auth = client_set.issuperset(required_set)
        # We're counting 401, 403 as auth denied
        response_denied = response.status_code == 401 or response.status_code == 403
        if should_auth:
            return not response_denied
        else:
            return response_denied


if __name__ == "__main__":
    schema = os.getenv('API_SCHEMA')
    clients_list = os.getenv('API_CLIENTS')
    url = os.getenv('API_URL')

    exhaustive = (os.getenv('EXHAUSTIVE') == 'true')
    force_http = (os.getenv('FORCE_HTTP') == 'true')
    if force_http:
        print('Forcing HTTP mode')
        os.environ['OAUTHLIB_INSECURE_TRANSPORT'] = '1'
    run = Run(
        schema=schema,
        clients=clients_list,
        api_url=url,
        exhaustive=exhaustive)
    run.run()
