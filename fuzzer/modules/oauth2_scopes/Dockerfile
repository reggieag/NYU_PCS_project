FROM microsoft/restler-fuzzer:7.3.0

WORKDIR /fuzzer

COPY . /fuzzer/

RUN pip install -r /fuzzer/requirements.txt

ENTRYPOINT [ "/usr/bin/python3", "/fuzzer/run.py"]
