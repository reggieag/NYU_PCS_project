import os

import psycopg2

from contextlib import closing

SQL_FIXTURE_FILE = os.path.abspath(os.path.join(os.path.dirname(__file__), 'api_fixture_data.sql'))


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
    sql = 'select distinct table_name from information_schema.tables'
    postgres_conn = psycopg2.connect(**db_config)
    with closing(postgres_conn) as conn:
        with closing(conn.cursor()) as cur:
            cur.execute(sql)
            return cur.fetchall()


def truncate_tables(db_config):
    tables = get_tables(db_config)
    for schema_name, table_name in tables:
        sql = f'truncate table "{schema_name}"."{table_name}";'
        postgres_conn = psycopg2.connect(**db_config)
        with closing(postgres_conn) as conn:
            with closing(conn.cursor()) as cur:
                cur.execute(sql)


def drop_table(db_config, table):
    postgres_conn = psycopg2.connect(**db_config)
    with closing(postgres_conn) as conn:
        with closing(conn.cursor()) as cur:
            cur.execute(f'drop table if exists "{table}"')
