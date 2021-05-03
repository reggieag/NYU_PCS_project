class BadSecurityScheme(Exception):
    def __init__(
            self, message='Security schema section is improperly formatted'):
        super().__init__(message)


class SecurityType:
    """
    SecurityType represents a security schema definition
    specified by the schema.
    This is a base class for actual security types
    that OpenAPI supports
    This class is not meant to be instantiated directly
    """

    @staticmethod
    def type():
        raise NotImplementedError("Not implemented in base class")

    @staticmethod
    def parse(security_type):
        raise NotImplementedError("Not implemented in base class")


class OAuth2SecurityType(SecurityType):
    """
    OAUth2SecurityType represents the Oauth2 security type
    supported by OpenAPI
    """

    def __init__(self, implicit, password,
                 client_credentials, authorization_code):
        self._implicit = implicit
        self._password = password
        self._client_credentials = client_credentials
        self._authorization_code = authorization_code

    @staticmethod
    def type():
        return 'oauth2'

    @property
    def implicit(self):
        return self._implicit

    @property
    def password(self):
        return self._implicit

    @property
    def client_credentials(self):
        return self._client_credentials

    @property
    def authorization_code(self):
        return self._authorization_code

    @staticmethod
    def parse(security_type):
        implicit = None
        password = None
        client_credentials = None
        authorization_code = None
        for flow in security_type['flows'].keys():
            flow_info = security_type['flows'][flow]
            token_url = flow_info['tokenUrl']
            refresh_url = ''
            if 'refreshUrl' in flow_info:
                refresh_url = flow_info['refreshUrl']
            scopes = [x for x in flow_info['scopes'].keys()]
            if flow == 'clientCredentials':
                client_credentials = ClientCredentialsFlow(
                    scopes=scopes, token_url=token_url, refresh_url=refresh_url)
            else:
                raise BadSecurityScheme('{} is not yet supported'.format(flow))
        return OAuth2SecurityType(
            implicit, password, client_credentials, authorization_code)


class OAuth2Flow:
    """
    OAuth2Flow is a base class for various OAuth2 flows
    This class is not meant to be instantiated directly
    """

    def __init__(self, scopes=[], refresh_url="",
                 token_url="", authorization_url=""):
        self._scopes = scopes
        self._refresh_url = refresh_url
        self._token_url = token_url
        self._authorization_url = authorization_url

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


class ClientCredentialsFlow(OAuth2Flow):
    """
    ClientCredentialsFlow contains information for
    OAuth2 client credentials flow
    """

    def __init__(self, scopes=[], refresh_url="", token_url=""):
        super().__init__(scopes, refresh_url, token_url, None)

    @property
    def authorization_url(self):
        raise TypeError(
            "Authorization URL Not appliable to Client Credentials")


class SecuritySchemes:
    """
    SecuritySchemes contains information about the types of security
    an OpenAPI schema provides
    """

    def __init__(self, schema):
        self._schemes = {}
        self._parsers = {
            OAuth2SecurityType.type(): OAuth2SecurityType.parse
        }
        if schema is not None:
            self.parse_schema(schema)

    def parse_schema(self, schema):
        """
        Parses the securitySchemes section of an OpenAPI schema
        """
        if schema is None:
            raise BadSecurityScheme
        try:
            security_schemes = schema['components']['securitySchemes']
            for name in security_schemes.keys():
                security_type = security_schemes[name]
                if security_type['type'] not in self._parsers:
                    raise BadSecurityScheme(
                        message='{} type is not supported'.format(name))
                self._schemes[name] = self._parsers[security_type['type']](
                    security_type)

        except KeyError as e:
            raise BadSecurityScheme from e

    @property
    def schemes(self):
        return self._schemes
