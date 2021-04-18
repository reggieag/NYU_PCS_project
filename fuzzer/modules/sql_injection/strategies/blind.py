
def test_blind_sql_injection():
    # https://owasp.org/www-community/attacks/Blind_SQL_Injection
    # insert data into table. Need to come up with a way to manage state in the DB.
    # Send valid SQL string
    # interpret result
    # Send SQL string with or 1=2 and see if the value changes
    pass

