from oauthlib.oauth2 import BackendApplicationClient
from requests_oauthlib import OAuth2Session


class AuthRequestException(Exception):
    def __init__(message=""):
        super().__init__(message)


class AuthRequest:
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
            oauth = OAuth2Session(client=oauth_client, scope=client.scopes)
            # Not mentioned in documents: include_client_id must be set to true
            # Otherwise it looks like the library just assumes Basic Auth for us
            # Which is completely contrary to the documentation which expects us
            # To explicitly provide HTTPBasicAuth
            # Auth check:
            # https://github.com/requests/requests-oauthlib/blob/master/requests_oauthlib/oauth2_session.py#L286
            # Use Basic auth if include_client_id is not set:
            # https://github.com/requests/requests-oauthlib/blob/master/requests_oauthlib/oauth2_session.py#L301
            # Meanwhile documentation:
            # https://requests-oauthlib.readthedocs.io/en/latest/oauth2_workflow.html#backend-application-flow
            # This took me a few hours and digging through their source code to
            # figure out...
            # Also as it turns out, passing in the scopes to OAuth2Session doesn't do anything
            # Basically, this documentation is garbage
            request_scopes = ','.join(client.scopes)
            token = oauth.fetch_token(
                client_credentials.token_url,
                client_id=client.id,
                client_secret=client.secret,
                include_client_id=True,
                scope=request_scopes)
            return (token, oauth)
        except KeyError as e:
            raise AuthRequestException from e
