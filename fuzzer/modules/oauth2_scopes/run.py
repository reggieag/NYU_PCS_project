import os
import sys
from clients import Clients
from security import SecuritySchemes
from run_generator import Generator
from oauth import OAuth2Requests


class Run:
    """
    Run contains the functionality to parse the provided variables, and run the actual module
    as described in the README
    """

    def __init__(self, schema='', clients='', api_url='', exhaustive=False):
        self._security_schemes = SecuritySchemes(schema=schema)
        parsed_clients = Clients(clients_list=clients)
        self._runs = Generator(
            clients=parsed_clients.clients,
            exhaustive=exhaustive).generate()
        self._exahustive = exhaustive
        self._api_url = api_url

    def run(self):
        """
        Start the module and fuzz the api
        """
        print(
            'Starting oauth2_scoeps fuzzer on exhaustive mode: {}'.format(
                self._exahustive))
        print('Using url {}'.format(self._api_url))
        print(
            'Using security schemes: {}'.format(
                self._security_schemes.schemes))
        print('Starting')
        for run in self._runs:
            print(
                'Starting run with client id: {} secret: {}'.format(
                    run.id, run.secret))
            print('Using scopes {}'.format(run.scopes))
            oauth_session = OAuth2Requests(self._security_schemes)
            try:
                # TODO: Read from tree, and fill in security scheme dynamically
                request = oauth_session.create_request(run, 'standard')
                print('token: {}'.format(request[0]))
                # TODO: Read this from RESTler grammar
                r = request[1].get('http://127.0.0.1:8080/data')
                print('response: {}'.r.json())
                # TODO: Validate scope and check response
            except Exception as e:
                print(e)


if __name__ == "__main__":
    schema = os.getenv('API_SCHEMA')
    clients_list = os.getenv('API_CLIENTS')
    url = os.getenv('API_URL')

    exhaustive = (os.getenv('EXHAUSTIVE') == 'true')
    force_http = (os.getenv('FORCE_HTTP') == 'true')
    if force_http:
        os.environ['OAUTHLIB_INSECURE_TRANSPORT'] = '1'
    run = Run(
        schema=schema,
        clients=clients_list,
        api_url=url,
        exhaustive=exhaustive)
    run.run()
