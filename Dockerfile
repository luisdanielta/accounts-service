FROM golang:1.20

WORKDIR /usr/src/app

COPY . .

EXPOSE 9001

RUN go mod download