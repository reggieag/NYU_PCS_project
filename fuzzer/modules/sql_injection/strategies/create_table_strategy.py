import logging
import uuid

from api.run_generator import Generator
from api.request_generator import DefaultRandomGenerator
from api.clients import Clients
from api.schema import Schema
from api.auth_request import AuthRequest

from strategies.lib.postgres import get_tables, drop_table


class CreateTableGenerator(DefaultRandomGenerator):
    def __init__(self, uuid):
        super().__init__()
        self.uuid = uuid

    def string(self):
        return f"""' where 1=2; create table "{self.uuid}" (id int); -- """


class Run:
    def __init__(self, schema='', clients='', api_url='', exhaustive=False, db_config={}):
        temp_file = self._write_temporary_file(schema=schema)
        logging.debug('generating temporary schema file {}'.format(temp_file))
        self._schema = Schema(temp_file, api_url)
        self._security_schemes = self._schema.security_schemes
        parsed_clients = Clients(clients_list=clients)
        self._runs = Generator(
            clients=parsed_clients.clients,
            exhaustive=exhaustive).generate()
        self._exahustive = exhaustive
        self._api_url = api_url
        self.db_config = db_config

    def _write_temporary_file(self, schema=''):
        name = 'temporary_schema_file.yaml'
        with open(name, 'w') as temp:
            temp.write(schema)
        return name

    def run(self):
        """
        Start the module and fuzz the api
        """
        logging.info(self._schema.paths)
        logging.info(self._runs)
        oauth_session = AuthRequest(self._security_schemes)
        attack_uuid = uuid.uuid4()
        gucci = True
        for run in self._runs:
            logging.info('Starting run with client id: {}'.format(run.id))
            logging.info('Using available client scopes {}'.format(run.scopes))
            for path in self._schema.paths:
                url = path.generate_path(generator=CreateTableGenerator(uuid=attack_uuid))
                body = None
                if path.body_required:
                    body = path.generate_request_body(
                        application_type=path.application_types[0],
                        generator=CreateTableGenerator(uuid=attack_uuid))
                    logging.info(
                        'Testing path: {} with generated url {} method {} body {}'.format(
                            path.path, url, path.request_method, body))
                for security in path.security.keys():
                    logging.info(
                        'Using security scheme {} requiring scopes: {}'.format(
                            security, path.security[security]))
                    try:
                        request = oauth_session.create_request(run, security)
                        logging.debug('Acquired token: {}'.format(request[0]))
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
                        logging.debug(
                            'response code: {}'.format(
                                response.status_code))
                    except Exception as e:
                        logging.error(e)

                    tables = get_tables(self.db_config)
                    if any([True for table in tables if str(attack_uuid) == table[0]]):
                        logging.warning(f"Vulnerability found for {url} method {path.request_method} body {body}")
                        gucci = False

                        drop_table(self.db_config, attack_uuid)
        return gucci
