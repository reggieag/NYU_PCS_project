import requests
import uuid

from .lib.api import build_requests_from_restler_grammar


# TODO: take lab6_string_interpolation strategy and hook it into RESTler

def generate_api_calls_restler(grammar_file, host, params):
    for request in build_requests_from_restler_grammar(grammar_file):
        if request.has_path_parameters:
            yield request.concretize_api_string(host, params)


def string_interpolation_test(host, grammar):
    # Generate api call
    # Make API call
    # Check result
    assert True