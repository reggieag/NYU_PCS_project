FROM python:3

WORKDIR /fuzzer

COPY . /fuzzer/

RUN python -m pip install -e /fuzzer/api

ENTRYPOINT [ "python", "/fuzzer/run.py"]
