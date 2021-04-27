import logging
import requests
import uuid

from functools import partial

from .lib.postgres import get_tables, drop_table, DBCONFIG


def openapi_api_call_generator(base_url, attack_param):
    # TODO: Implement this for reals
    data = {'name': attack_param, 'quantity': 2}
    return [('http://127.0.0.1:8080/data/1', data, 'post')]


def create_table_attack_test(base_url, api_call_generator):
    gucci = True
    attack_uuid = uuid.uuid4()
    attack = f"' where 1=2 ; create table {attack_uuid} (id int); -- "
    for url, body, method in api_call_generator(base_url, attack):
        logging.info(f"Testing generated {url} method {method} body {body}")
        if method == 'post':
            requests.post(url, data=body)
        elif method == 'get':
            requests.get(url, data=body)
        elif method == 'delete':
            requests.delete(url, data=body)
        elif method == 'patch':
            requests.patch(url, data=body)

        tables = get_tables(DBCONFIG)

        if any([True for table in tables if str(attack_uuid) == table[0]]):
            logging.warning(f"Vulnerability found for {url} method {method} body {body}")
            gucci = False

        drop_table(DBCONFIG, attack_uuid)

    return gucci


openapi_create_table_attack_test = partial(create_table_attack_test, api_call_generator=openapi_api_call_generator)