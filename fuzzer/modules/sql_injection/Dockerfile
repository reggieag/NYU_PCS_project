FROM microsoft/restler-fuzzer:7.3.0

WORKDIR /fuzzer

COPY run.py /fuzzer/run.py

ENTRYPOINT [ "/usr/bin/python3", "/fuzzer/run.py"]
