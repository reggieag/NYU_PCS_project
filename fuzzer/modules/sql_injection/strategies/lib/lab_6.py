VULNERABLE_ENDPOINT = 'http://127.0.0.1:4000/secret'


def generate_api_calls(param):
    return [(VULNERABLE_ENDPOINT, {'name': param}, 'Post')]