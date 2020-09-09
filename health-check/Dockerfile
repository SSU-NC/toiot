FROM golang:1.13

WORKDIR $GOPATH/src/github.com/KumKeeHyun/PDK/health-check
COPY . $GOPATH/src/github.com/KumKeeHyun/PDK/health-check

RUN go build -o main .

ENTRYPOINT [ "./main" ]