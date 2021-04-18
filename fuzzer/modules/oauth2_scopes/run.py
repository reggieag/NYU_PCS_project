import time
import os
import sys
from clients import Clients
from security import SecuritySchemes

schema = os.getenv('API_SCHEMA')
clients_list = os.getenv('API_CLIENTS')

security_schemes = SecuritySchemes(schema)
clients = Clients(clients_list)

print(clients.clients)
print(security_schemes.schemes)
print(os.getenv("EXHAUSTIVE"))

time.sleep(4)
print("pretend we're running")
time.sleep(4)
sys.exit("Will this report error back to Go?")
