FROM golang:1.13

WORKDIR $GOPATH/src/github.com/KumKeeHyun/PDK/logic-core
COPY . $GOPATH/src/github.com/KumKeeHyun/PDK/logic-core

RUN go build -o main .

ENTRYPOINT [ "./main" ]