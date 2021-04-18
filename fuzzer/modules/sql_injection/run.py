import time
import os
import sys

print(os.environ)

# TODO: Poll DB to make sure that its up and running
# TODO: Poll API to make sure that its up and running (In this case, the API will be running the moment the DB is ready)
# TODO: Compile openAPI spec with RESTler
# TODO: Read in grammar.json
# TODO: Generate SQL injection calls
# TODO: Verify DB changes and report error appropriately


time.sleep(4)
print("pretend we're running")
time.sleep(4)
sys.exit("Will this report error back to Go?")
