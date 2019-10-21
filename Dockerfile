FROM golang:latest

ENV GO111MODULE=on

WORKDIR /app
ADD . /app

RUN go get github.com/pilu/fresh

EXPOSE 9090

CMD ["fresh"]
