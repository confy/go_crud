FROM golang:1.17-alpine

RUN apk update && apk add --no-cache gcc musl-dev

WORKDIR $GOPATH/src/github.com/confy/go_crud
ENV GO111MODULE=on

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main

EXPOSE 8080

CMD [ "./main" ]