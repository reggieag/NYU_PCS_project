import requests
import uuid

from functools import partial

from .lib.postgres import get_tables, drop_table, DBCONFIG


def openapi_api_call_generator(base_url, uuid):
    # TODO: Implement this for reals
    attack = f"' create table {uuid} ; -- "
    data = {'name': attack, 'quantity': 2}
    return [('http://127.0.0.1:8080/data/1', data, 'post')]


def create_table_attack_test(base_url, api_call_generator):
    attack_uuid = uuid.uuid4()
    for url, data, method in api_call_generator(base_url, attack_uuid):
        print(f"testing {url}")
        if method == 'post':
            requests.post(url, data=data)
        elif method == 'get':
            requests.get(url, data=data)
        elif method == 'delete':
            requests.delete(url, data=data)
        elif method == 'patch':
            requests.patch(url, data=data)

        tables = get_tables(DBCONFIG)

        if any([True for table in tables if attack_uuid in table]):
            print(f"Vulnerability found for {url} {data} {method}")

        drop_table(DBCONFIG, attack_uuid)


openapi_create_table_attack_test = partial(create_table_attack_test, api_call_generator=openapi_api_call_generator)
