import requests
import uuid

from functools import partial

from .lib import lab_6

ATTACK_TEMPLATES = [
    "' union select '{uuid}' --",
]


# In this library we attempt to implement an automated test using the same attack in lab 6.
# What we found though is that the detection method is difficult to automate since API responses
# can be fairly arbitrary. You may have to generate a parse_response_fn for each api_call outputted by
# the api_call_generator

class UnionAttackVulnerabilityException(Exception):
    """
    Detected a union attack vulnerability.
    """


def union_attack_test(base_url, api_call_generator, parse_response_fn):
    """
    This test attempts a union attack against an endpoint.

    :param base_url: The API base_url. Unused for this particular test.
    :param api_call_generator: Restler grammar. Unused for this particular test.
    :param parse_response_fn: Function to parse the UUID from the response
    :return: Boolean denoting whether or not the test has passed
    """
    attack_uuid = uuid.uuid4()
    for attack_template in ATTACK_TEMPLATES:
        attack_param = attack_template.format(uuid=attack_uuid)
        for url, data, method in api_call_generator(base_url, attack_param):
            if method == 'post':
                print(f"testing {url}")
                res = requests.post(url, data=data)
            elif method == 'get':
                print(f"testing {url}")
                res = requests.get(url, data=data)
            else:
                continue

            # No we check if we were able to insert the UUID into the result. We need to define
            # a parser that can extract the UUID from the response text.
            if res.status_code == 200 and str(attack_uuid) == parse_response_fn(res.text):
                raise UnionAttackVulnerabilityException(f"Union attack vulnerability detected for url {url} and data {data}")


lab_6_union_attack_test = partial(union_attack_test, api_call_generator=lab_6.generate_api_calls, parse_response_fn=lab_6.parse_response)

def generate_api_calls_openapi(base_url, attack_param):
    url = base_url + '/data/' + attack_param
    return [(url, None, 'get')]


def parse_response_openapi(response):
    return None

openapi_union_attack_test = partial(union_attack_test, api_call_generator=generate_api_calls_openapi, parse_response_fn=parse_response_openapi)
