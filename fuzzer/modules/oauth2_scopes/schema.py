from prance import ResolvingParser
from security import SecuritySchemes
from path import Path


class Schema:
    """
    Schema represents an OpenAPI schema, and the various section that it composes
    Aside from the security schema, it allows users to traverse the api paths
    """

    def __init__(self, parse_file='', base_url=''):
        api = ResolvingParser(parse_file)
        self._security_schemes = SecuritySchemes(api.specification)
        self._base_url = base_url
        self._paths = self._parse_paths(api.specification['paths'])

    @property
    def security_schemes(self):
        """
        Get the security schemas defined in the schema
        """
        return self._security_schemes

    @property
    def paths(self):
        """
        Get all the paths defined in the schema
        """
        return self._paths

    def _parse_paths(self, paths):
        """
        Utility function to parse through paths defined in schema
        """
        parsed_paths = []
        for path in paths.keys():
            path_methods = paths[path]
            for method in path_methods.keys():
                method_info = path_methods[method]
                application_types = {}
                body_required = False
                if 'requestBody' in method_info:
                    application_types = method_info['requestBody']['content']
                    body_required = method_info['requestBody']['required']
                security = []
                if 'security' in method_info:
                    security = method_info['security']
                parameters = []
                if 'parameters' in method_info:
                    parameters = method_info['parameters']
                new_path = Path(
                    base_url=self._base_url,
                    path=path,
                    method=method,
                    application_types=application_types,
                    body_required=body_required,
                    security=security,
                    parameters=parameters)
                parsed_paths.append(new_path)
        return parsed_paths
