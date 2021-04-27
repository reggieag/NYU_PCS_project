import os
import logging
import sys

from strategies.create_table_strategy import openapi_create_table_attack_test


def run_tests(url):
    gucci = True
    gucci = openapi_create_table_attack_test(url) and gucci
    if gucci:
        sys.exit(0)
    sys.exit(1)


if __name__ == "__main__":
    schema = os.getenv('API_SCHEMA')
    # TODO: Pull and construct from environment variables
    url = 'http://127.0.0.1:8080'
    # url = os.getenv('API_URL')

    log_level = os.getenv('LOG_LEVEL')
    logging_format = '%(levelname)s:module_oauth2_scopes:%(message)s'
    if log_level == 'DEBUG':
        logging.basicConfig(format=logging_format, level=logging.DEBUG)
    elif log_level == 'INFO':
        logging.basicConfig(format=logging_format, level=logging.INFO)
    elif log_level == 'WARNING':
        logging.basicConfig(format=logging_format, level=logging.WARNING)
    elif log_level == 'ERROR':
        logging.basicConfig(format=logging_format, level=logging.ERROR)
    else:
        logging.basicConfig(format=logging_format, level=logging.INFO)

    run_tests(url)
