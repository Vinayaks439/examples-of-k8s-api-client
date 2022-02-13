FROM python3.7.3

RUN pip3 install -U pip \
    && pip3 install -r requirements.txt
  
WORKDIR /usr/local/bin

COPY deletefailedsuceededpod.py /usr/local/bin/.

CMD ["python3","deletefailedsuceededpod.py"]
