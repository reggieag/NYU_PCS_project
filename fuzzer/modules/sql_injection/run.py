import time
import os
import sys

from strategies.blind import blind_sql_injection_test

# TODO: Pull and construct from environment variables
HOST_ADDRESS = 'http://127.0.0.1:8080'

# TODO: Dynamically generate in Restler
GRAMMAR_FILE = os.path.join(os.path.dirname(__file__), 'Compile/grammar.json')


print(os.environ)

# TODO: Poll DB to make sure that its up and running
# TODO: Poll API to make sure that its up and running (In this case, the API will be running the moment the DB is ready)
# TODO: Compile openAPI spec with RESTler
# TODO: Read in grammar.json
# TODO: Generate SQL injection calls
# TODO: Verify DB changes and report error appropriately

blind_sql_injection_test(GRAMMAR_FILE, HOST_ADDRESS)

time.sleep(4)
print("pretend we're running")
time.sleep(4)
sys.exit("Will this report error back to Go?")
