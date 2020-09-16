FROM golang:1.13

WORKDIR $GOPATH/src/github.com/KumKeeHyun/toiot/application
COPY . $GOPATH/src/github.com/KumKeeHyun/toiot/application

RUN go build -o main .

ENTRYPOINT [ "./main" ]