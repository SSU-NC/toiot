FROM golang:1.13

WORKDIR $GOPATH/src/github.com/KumKeeHyun/toiot/health-check
COPY . $GOPATH/src/github.com/KumKeeHyun/toiot/health-check

RUN go build -o main .

ENTRYPOINT [ "./main" ]