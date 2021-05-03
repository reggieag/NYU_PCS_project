FROM fuzzer-base:latest

WORKDIR /fuzzer

COPY . /fuzzer/

RUN cp -rf /fuzzer_base /fuzzer/api

RUN python -m pip install -e /fuzzer/api
RUN pip install -r /fuzzer/requirements.txt

ENTRYPOINT [ "python", "/fuzzer/run.py"]
