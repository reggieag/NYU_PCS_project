
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
    pass

class OAuth2SecurityType(SecurityType):
    """
    OAUth2SecurityType represents the Oauth2 security type
    supported by OpenAPI
    """
    pass

class OAuth2Flow:
    """
    OAuth2Flow is a base class for various OAuth2 flows
    """
    pass

class ClientCredentialFlow(OAuth2Flow):
    """
    ClientCredentialFlow contains information for 
    OAuth2 client credentials flow
    """
    pass

class SecuritySchema:
    """
    SecuritySchema contains information about the types of security
    an OpenAPI schema provides
    """
    pass
