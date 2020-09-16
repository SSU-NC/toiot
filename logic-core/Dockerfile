FROM golang:1.13

WORKDIR $GOPATH/src/github.com/KumKeeHyun/toiot/logic-core
COPY . $GOPATH/src/github.com/KumKeeHyun/toiot/logic-core

RUN go build -o main .

ENTRYPOINT [ "./main" ]
