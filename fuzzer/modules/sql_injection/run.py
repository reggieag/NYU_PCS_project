import time
import os
import sys

from strategies.union_strategy import union_attack_test

# TODO: Pull and construct from environment variables
HOST_ADDRESS = 'http://127.0.0.1:8080'

# TODO: Dynamically generate in Restler
GRAMMAR_FILE = os.path.join(os.path.dirname(__file__), 'Compile/grammar.json')


print(os.environ)

union_attack_test(GRAMMAR_FILE, HOST_ADDRESS)

time.sleep(4)
print("pretend we're running")
time.sleep(4)
sys.exit("Will this report error back to Go?")
