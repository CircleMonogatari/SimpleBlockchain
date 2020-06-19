FROM golang:latest
LABEL Fisher="circlemono@outlook.com"

ENV GOPROXY https://goproxy.cn

EXPOSE 8080

RUN git clone https://github.com/CircleMonogatari/SimpleBlockchain.git \
    && cd SimpleBlockchain  \
    && go build \
    && nohub ./SimpleBlockchain -mode 1
