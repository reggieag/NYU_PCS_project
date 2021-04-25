import requests
import uuid

from functools import partial

from .lib.lab_6 import generate_api_calls


ATTACK_TEMPLATES = [
    # "' union select '{uuid}' --",
    "1 union select '{uuid}' --",
]


class UnionAttackVulnerabilityException(Exception):
    """
    Detected a union attack vulnerability.
    """


def generate_api_calls_openapi(base_url, attack_param):
    # TODO: Implement this for reals.
    url = base_url + '/data/' + attack_param
    return [(url, None, 'get')]


def union_attack_test(base_url, api_call_generator):
    """
    This test attempts a union attack against an endpoint.

    :param base_url: The API base_url. Unused for this particular test.
    :param grammar: Restler grammar. Unused for this particular test.
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

            # print(res.status_code)
            # print(res.text)

            # No we check if we were able to insert the uuid into the result. If we were
            # we know that a sql injection attack happened since there is almost no chance
            # that the uuid we just generated is going to be constructed by the server naturally.
            if res.status_code == 200 and str(attack_uuid) in res.text:
                print(res.text)
                raise UnionAttackVulnerabilityException(f"Union attack vulnerability detected for url {url} and data {data}")


lab_6_union_attack_test = partial(union_attack_test, api_call_generator=generate_api_calls)
openapi_union_attack_test = partial(union_attack_test, api_call_generator=generate_api_calls_openapi)
