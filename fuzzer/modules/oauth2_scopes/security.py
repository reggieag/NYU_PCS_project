import yaml


class BadSecuritySchema(Exception):
    def __init__(self, message='Security schema section is improperly formatted'):
        super().__init__(message)


class SecurityType:
    """
    SecurityType represents a security schema definition
    specified by the schema.
    This is a base class for actual security types
    that OpenAPI supports
    """

    def __init__(self, type):
        self._type = type

    @property
    def type(self):
        return self._type

    def __new__(self, *args, **kwargs):
        if self is SecurityType:
            raise TypeError("Cannot instantiate SecurityType base class")
        return object.__new__(self, *args, **kwargs)


class OAuth2SecurityType(SecurityType):
    """
    OAUth2SecurityType represents the Oauth2 security type
    supported by OpenAPI
    """

    def __init__(self):
        super().__init__('oauth2')


class OAuth2Flow:
    """
    OAuth2Flow is a base class for various OAuth2 flows
    """

    def __init__(self, scopes=[], refresh_url="", token_url="", authorization_url=""):
        self._scopes = []
        self._refresh_url = ""
        self._token_url = ""
        self._authorization_url = ""

    @property
    def scopes(self):
        return self._scopes

    @property
    def refresh_url(self):
        return self._refresh_url

    @property
    def token_url(self):
        return self._token_url

    @property
    def authorization_url(self):
        return self._authorization_url

    def __new__(self, *args, **kwargs):
        if self is OAuth2Flow:
            raise TypeError("Cannot instantiate OAuth2Flow base class")
        return object.__new__(self, *args, **kwargs)


class ClientCredentialFlow(OAuth2Flow):
    """
    ClientCredentialFlow contains information for 
    OAuth2 client credentials flow
    """
    pass


class SecuritySchemes:
    """
    SecuritySchemes contains information about the types of security
    an OpenAPI schema provides
    """

    def __init__(self):
        self._schemes = {}

    def parse_schema(self, schema):
        """
        Parses the securitySchemes section of an OpenAPI schema
        """
        schema = yaml.safe_load(schema)
        if schema is None:
            raise BadSecuritySchema
        try:
            security_schemes = schema['components']['securitySchemes']
        except KeyError:
            raise BadSecuritySchema

    @property
    def schemes(self):
        return self._schemes
