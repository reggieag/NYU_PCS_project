import time
import os
import sys
from clients import Clients

print(os.environ)

time.sleep(4)
print("pretend we're running")
time.sleep(4)
sys.exit("Will this report error back to Go?")
