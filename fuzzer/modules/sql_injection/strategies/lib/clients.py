import yaml


class BadClientConfig(Exception):
    def __init__(self, message='Input clients data is improperly formatted'):
        super().__init__(message)


class Client:
    """
    Object to represent to hold data about a client:
    client id, client secret, and client scopes
    """

    def __init__(self, id, secret, scopes):
        self._id = id
        self._secret = secret
        self._scopes = scopes

    @property
    def id(self):
        return self._id

    @property
    def secret(self):
        return self._secret

    @property
    def scopes(self):
        return self._scopes


class Clients:
    """
    Meta class to manipulate lists of clients, such as creating clients from
    configuration files, etc
    """

    def __init__(self, clients_list):
        self._clients = []
        if clients_list is not None:
            self.parse_clients_list(clients_list)

    def parse_clients_list(self, clients_list):
        """
        Given a list of client configurations, parses the data
        and creates a list of Clients
        """
        clients = yaml.safe_load(clients_list)
        if clients is None or 'clients' not in clients:
            raise BadClientConfig
        for client in clients['clients']:
            try:
                self._clients.append(
                    Client(client['client_id'], client['client_secret'], client['scopes']))
            except KeyError as e:
                raise BadClientConfig(
                    message='{} is missing information'.format(client)) from e

    @property
    def clients(self):
        return self._clients
