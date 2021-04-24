import requests
import uuid

from lib.lab_6 import generate_api_calls


class UnionAttackVulnerabilityException:
    """
    Detected a union attack vulnerability.
    """


def union_attack_test(host, grammar):
    """
    This test is run against lab6 as a proof of concept that we can automate
    detecting a particular SQL injection attack that exploits string interpolation.

    :param host: The API host. Unused for this particular test.
    :param grammar: Restler grammar. Unused for this particular test.
    :return: Boolean denoting whether or not the test has passed
    """
    attack_uuid = uuid.uuid4()
    attack_param = "' union select {uuid} --".format(uuid=attack_uuid)
    for url, data in generate_api_calls(attack_param):
        res = requests.post(url, data=data)

        # No we check if we were able to insert the uuid into the result. If we were
        # we know that a sql injection attack happened since there is almost no chance
        # that the uuid we just generated is going to be constructed by the server naturally.
        if str(attack_uuid) in res.text:
            raise UnionAttackVulnerabilityException
