import json


class UnexpectedPathParameters(Exception):
    pass


class Request:
    """
    Object representing the different attributes of a request.
    """
    def __init__(self):
        self.endpoint_format = None
        self.method = None
        self.path_components = []

    @property
    def dynamic_object_count(self):
        return sum([1 if pc.is_dynamic else 0 for pc in self.path_components])

    @property
    def has_path_parameters(self):
        return any([pc.is_dynamic for pc in self.path_components])

    def _concretize_endpoint_with_path_params(self, path_params):
        """
        Builds a concrete API string using the path parameters passed in.

        :param path_params: A list of strings to be interpolated into the dynamic path components.
        :return: concrete_endpoint_string
        """
        concrete_endpoint_string = ''
        if len(path_params) != self.dynamic_object_count:
            raise UnexpectedPathParameters(f"Expected param count of {self.dynamic_object_count} but received {len(path_params)}")

        i = 0
        for path_component in self.path_components:
            if path_component.is_dynamic:
                concrete_endpoint_string += '/' + str(path_params[i])
                i += 1
            else:
                concrete_endpoint_string += '/' + path_component.string

        return concrete_endpoint_string

    def concretize_api_string(self, host, path_params=None):
        """
        Builds a concrete API string that can be executed against an API.

        :param host: The host address of an API
        :param path_params: The path parameters (optional)
        :return: concrete api string
        """
        if not self.has_path_parameters:
            return host + self.endpoint_format

        return host + self._concretize_endpoint_with_path_params(path_params)


class PathComponent:
    """
    Represents the different path components and their individual characteristics.
    """
    def __init__(self, string=None, is_dynamic=None):
        self.string = str(string)
        self.is_dynamic = is_dynamic or False


def _parse_path(path):
    """
    Translates a RESTler grammar path into a Request.path attribute.
    :param path: A RESTler grammar path. A list
    :return: list
    """
    parsed_path = []
    for component in path:
        if 'Constant' in component:
            parsed_path.append(PathComponent(component['Constant'][1]))
        if 'DynamicObject' in component:
            component_string = '_'.join(component['DynamicObject'].split('_')[3:])
            parsed_path.append(PathComponent(component_string, is_dynamic=True))

    return parsed_path


def build_requests_from_restler_grammar(grammar_file):
    """
    A factory that translates a RESTler grammar into the Request object. The reason
    we do this translation is to make the request attributes easier to work with when
    generating sql injection attacks.

    :param grammar_file:
    :return:
    """
    requests = []
    with open(grammar_file, 'r') as fh:
        grammar = json.load(fh)
    for request_grammar in grammar['Requests']:
        request = Request()

        request.endpoint_format = request_grammar['id']['endpoint']
        request.method = request_grammar['id']['method']
        request.path_components = _parse_path(request_grammar['path'])

        requests.append(request)

    return requests
