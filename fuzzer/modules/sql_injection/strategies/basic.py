from .lib.api import build_requests_from_restler_grammar


def basic_test(grammar_file, host):
    # generate API calls.
    for request in build_requests_from_restler_grammar(grammar_file):
        if request.has_path_parameters:
            print(request.dynamic_object_count)
            print(request.concretize_api_string(host, [1]))
    # execute API calls and introspect.
    pass