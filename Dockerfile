FROM golang:1.17.2-bullseye as develop

WORKDIR /app

RUN apt-get update
RUN apt-get install --yes nodejs npm
RUN go get github.com/go-delve/delve/cmd/dlv@v1.7.2
RUN npm install --global nodemon@2.0.14

COPY go.* ./
RUN go mod download -x
