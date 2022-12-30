FROM golang:latest

RUN mkdir /client

WORKDIR /client

COPY  ./ ./

RUN go env -w GO111MODULE=on

RUN go mod download

RUN #go build -o main .

RUN go build cmd/main.go

CMD ["./main"]