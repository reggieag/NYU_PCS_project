import requests

from .lib.api import build_requests_from_restler_grammar
from .lib.postgres_setup import reset_fixture_data

# https://owasp.org/www-community/attacks/Blind_SQL_Injection
# insert data into table. Need to come up with a way to manage state in the DB.
# Send valid SQL string
# interpret result
# Send SQL string with or 1=2 and see if the value changes


def blind_sql_injection_test(grammar_file, host):
    for request in build_requests_from_restler_grammar(grammar_file):
        if request.has_path_parameters and request.method == 'Get':
            reset_fixture_data()
            request_str = request.concretize_api_string(host, [1])
            res = requests.get(request_str)
            # print(request.dynamic_object_count)
            # print(request.concretize_api_string(host, [1]))
            print(request_str)
            print(res.status_code)
            print(res.text)
    # execute API calls and introspect.
    pass