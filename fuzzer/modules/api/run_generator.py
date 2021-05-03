from clients import Client
from itertools import combinations


class Generator:
    """
    Generator takes a list of clients, and generates the scopes for each tree execution
    as defined in the README
    """

    def __init__(self, clients=[], exhaustive=False):
        self._clients = clients
        self._exhaustive = exhaustive

    def generate(self):
        """
        Generates list of clients with required scopes for a tree execution
        """
        if self._exhaustive:
            return self._generate_exhaustive()
        else:
            return self._generate_non_exhaustive()

    def _generate_non_exhaustive(self):
        """
        Generates non exhaustive run list. If a client doesn't have any scopes associated, then
        only a single run is generated
        """
        runs = []
        for i in self._clients:
            if len(i.scopes) == 0:
                runs.append(i)
            else:
                no_scope = Client(i.id, i.secret, [])
                runs.append(no_scope)
                runs.append(Client(i.id, i.secret, i.scopes))
        return runs

    def _generate_exhaustive(self):
        """
        Generates an exhaustive list for all runs according to README
        """
        runs = []
        for client in self._clients:
            # Generate all scope combinations
            for i in range(0, len(client.scopes) + 1):
                scopes = [
                    list(x) for x in list(
                        combinations(
                            client.scopes, i))]
                clients = [Client(client.id, client.secret, x) for x in scopes]
                runs.extend(clients)
        return runs
