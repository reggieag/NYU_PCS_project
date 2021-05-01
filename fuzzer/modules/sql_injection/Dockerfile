FROM python:3

WORKDIR /fuzzer

COPY . /fuzzer/

RUN pip install -r /fuzzer/requirements.txt

ENTRYPOINT [ "python", "/fuzzer/run.py"]
