FROM golang:1.23

WORKDIR /app

RUN apt update && \
  apt install -y librdkafka-dev

CMD ["tail", "-f", "/dev/null"]