FROM ubuntu:kinetic-20220830

RUN apt-get update \
  && apt-get install -y python3-pip python3-dev \
  && cd /usr/local/bin \
  && ln -s /usr/bin/python3 python \
  && pip3 install --upgrade pip

RUN mkdir -p /opt/app

COPY . /opt/app

WORKDIR /opt/app

ARG test_word

CMD [ "python", "./modified_palindrome.py"]