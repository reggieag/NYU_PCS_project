import time
import os
import sys

from strategies.string_interpolation import string_interpolation_test
from strategies.lab6_string_interpolation import lab6_string_interpolation_test

# TODO: Pull and construct from environment variables
HOST_ADDRESS = 'http://127.0.0.1:8080'

# TODO: Dynamically generate in Restler
GRAMMAR_FILE = os.path.join(os.path.dirname(__file__), 'Compile/grammar.json')


print(os.environ)

string_interpolation_test(GRAMMAR_FILE, HOST_ADDRESS)
lab6_string_interpolation_test(GRAMMAR_FILE, HOST_ADDRESS)

time.sleep(4)
print("pretend we're running")
time.sleep(4)
sys.exit("Will this report error back to Go?")
