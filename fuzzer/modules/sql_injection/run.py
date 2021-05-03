import os
import logging
import sys

from strategies.create_table_strategy import Run


if __name__ == "__main__":
    schema = os.getenv('API_SCHEMA')
    clients_list = os.getenv('API_CLIENTS')
    url = os.getenv('API_URL')

    db_config = {
        "host": os.environ['DB_HOST'],
        "port": os.environ['DB_PORT'],
        "user": os.environ['DB_USERNAME'],
        "password": os.environ['DB_PASSWORD'],
        "dbname": os.environ['DB_NAME'],
    }

    exhaustive = (os.getenv('EXHAUSTIVE') == 'true')
    force_http = (os.getenv('FORCE_HTTP') == 'true')
    log_level = os.getenv('LOG_LEVEL')
    logging_format = '%(levelname)s:sql_injection_module:%(message)s'
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

    if force_http:
        logging.info('Forcing HTTP mode')
        os.environ['OAUTHLIB_INSECURE_TRANSPORT'] = '1'

    gucci = Run(
        schema=schema,
        clients=clients_list,
        api_url=url,
        exhaustive=exhaustive,
        db_config=db_config).run()
    if gucci:
        sys.exit(0)
    sys.exit(1)
