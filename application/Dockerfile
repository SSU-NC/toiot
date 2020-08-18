FROM golang:1.13

WORKDIR $GOPATH/src/github.com/KumKeeHyun/PDK/application
COPY . $GOPATH/src/github.com/KumKeeHyun/PDK/application

RUN go build -o main .

ENTRYPOINT [ "./main" ]