from oauthlib.oauth2 import BackendApplicationClient
from requests_oauthlib import OAuth2Session


class OAuth2RequestsException(Exception):
    def __init__(message=""):
        super().__init__(message)


class OAuth2Requests:
    """
    Wrapper class to handle OAuth2 token access logic
    """

    def __init__(self, security_schemes):
        self._security_schemes = security_schemes

    def create_request(self, client, security_scheme):
        """
        Creates an OAuth2Session object with token already retrieved
        based on provided client id, secret, and scopes.
        Can be immediately used to make HTTP requests
        """
        try:
            security_scheme = self._security_schemes.schemes[security_scheme]
            # We will only support client credentials for now...
            client_credentials = security_scheme.client_credentials
            oauth_client = BackendApplicationClient(client_id=client.id)
            oauth = OAuth2Session(client=oauth_client)
            token = oauth.fetch_token(
                client_credentials.token_url,
                client_id=client.id,
                client_secret=client.secret)
            return (token, oauth)
        except KeyError as e:
            raise OAuth2RequestsException from e
