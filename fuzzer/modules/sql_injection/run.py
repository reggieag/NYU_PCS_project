import os

from strategies.union_strategy import union_attack_test

# TODO: Pull and construct from environment variables
HOST_ADDRESS = 'http://127.0.0.1:8080'

# TODO: Dynamically generate in Restler
GRAMMAR_FILE = os.path.join(os.path.dirname(__file__), 'Compile/grammar.json')


union_attack_test(HOST_ADDRESS, GRAMMAR_FILE)
