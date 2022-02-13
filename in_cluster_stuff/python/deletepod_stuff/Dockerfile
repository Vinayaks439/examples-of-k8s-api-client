FROM python:3.7.3

RUN pip3 install -U pip \
    && pip3 install kubernetes
  
WORKDIR /usr/local/bin

COPY deletefailedsuceededpod.py /usr/local/bin/.

CMD ["python3","deletefailedsuceededpod.py"]
