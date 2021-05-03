FROM python:3

WORKDIR /fuzzer

COPY . /fuzzer/

RUN python -m pip install -e /fuzzer/api
RUN pip install -r /fuzzer/requirements.txt

ENTRYPOINT [ "python", "/fuzzer/run.py"]
