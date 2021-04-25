from request_generator import DefaultRandomGenerator
import json


class InvalidPathException(BaseException):
    def __init__(self, message=""):
        super().__init__(message)


class Path:

    def __init__(self, base_url='', path='', method='', security=[],
                 application_types={}, body_required=False, parameters=[]):
        self._base_url = base_url
        self._path = path
        self._method = method
        self._security = self._parse_security(security)
        self._application_types = application_types
        self._body_required = body_required
        self._parameters = parameters
        self._path_parameters = self._filter_parameters(parameters, 'path')

    def _filter_parameters(self, parameters, param_in):
        params = {}
        filtered = [x for x in parameters if x['in'] == param_in]
        for x in filtered:
            params[x['name']] = x
        return params

    def _parse_security(self, security):
        parsed_security = {}
        for s_type in security:
            for data in s_type.keys():
                parsed_security[data] = s_type[data]
        return parsed_security

    @property
    def security(self):
        return self._security

    @property
    def application_types(self):
        if not self._application_types:
            return []
        return list(self._application_types.keys())

    @property
    def request_method(self):
        return self._method

    @property
    def body_required(self):
        return self._body_required

    @property
    def path(self):
        return self._path

    def generate_path(self, generator=DefaultRandomGenerator()):
        split_path = self._path.split('/')
        for i, part in enumerate(split_path):
            if len(part) >= 3 and part[0] == '{' and part[-1] == '}':
                param = part[1:-1]
                param_schema = self._path_parameters[param]['schema']
                split_path[i] = str(
                    self._generate_body_object(
                        param_schema, generator))
        rejoined_path = '/'.join(split_path)
        return self._base_url + rejoined_path

    def generate_request_body(self, application_type,
                              generator=DefaultRandomGenerator()):
        if application_type not in self._application_types:
            raise InvalidPathException(
                message='{} is not valid for this endpoint'.format(application_type))
        schema = self._application_types[application_type]['schema']
        body = self._generate_body_object(schema, generator)
        if application_type == 'application/json':
            return json.dumps(body)
        return body

    def _generate_body_object(self, schema, generator):
        def recurse_generate(schema, generator):
            schema_type = schema['type']
            body = None
            if schema_type == 'integer':
                body = generator.integer()
            elif schema_type == 'string':
                body = generator.string()
            elif schema_type == 'boolean':
                body = generator.boolean()
            elif schema_type == 'number':
                body = generator.number()
            elif schema_type == 'object':
                body = {}
                for sub_body in schema['properties'].keys():
                    body[sub_body] = recurse_generate(
                        schema['properties'][sub_body], generator)

            return body
        return recurse_generate(schema, generator)
