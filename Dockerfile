FROM golang:alpine

RUN apk update && apk add --no-cache git

RUN mkdir /application

WORKDIR /application

COPY . .

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/application/binary"]