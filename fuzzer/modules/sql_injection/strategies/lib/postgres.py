import os

import psycopg2

from contextlib import closing

SQL_FIXTURE_FILE = os.path.abspath(os.path.join(os.path.dirname(__file__), 'api_fixture_data.sql'))

# TODO replace with env variables
TARGET_DBCONFIG = {
    "host": "127.0.0.1",
    "port": 5432,
    "user": 'api_user',
    "password": 'password',
    "dbname": 'toy_api',
}


def read_sql_file(sql_filename, sql_dir=None):
    with open(sql_filename, "r") as sql_file:
        return sql_file.read()


def execute_postgres_sql(sql_filename, db_config, sql_dir=None):
    sql = read_sql_file(sql_filename, sql_dir)

    postgres_conn = psycopg2.connect(**db_config)
    with closing(postgres_conn) as conn:
        with closing(conn.cursor()) as cur:
            cur.execute(sql)
            conn.commit()


def get_tables(db_config):
    sql = 'select schema_name, table_name from information_schema.tables'
    postgres_conn = psycopg2.connect(**db_config)
    with closing(postgres_conn) as conn:
        with closing(conn.cursor()) as cur:
            cur.execute(sql)
            return cur.fetch_all()


def truncate_tables(db_config):
    tables = get_tables(db_config)
    for schema_name, table_name in tables:
        sql = f'truncate table "{schema_name}"."{table_name}";'
        postgres_conn = psycopg2.connect(**db_config)
        with closing(postgres_conn) as conn:
            with closing(conn.cursor()) as cur:
                cur.execute(sql)


def reset_fixture_data():
    truncate_tables(TARGET_DBCONFIG)
    execute_postgres_sql(SQL_FIXTURE_FILE, TARGET_DBCONFIG)