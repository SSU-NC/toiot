FROM golang:1.13

WORKDIR $GOPATH/src/github.com/seheee/PDK/application
COPY . $GOPATH/src/github.com/seheee/PDK/application

RUN go build -o main .

ENTRYPOINT [ "./main" ]