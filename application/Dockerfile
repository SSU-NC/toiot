FROM golang:1.13 as builder

WORKDIR $GOPATH/src/github.com/KumKeeHyun/toiot/application
COPY . .

RUN go build -o main .


FROM alpine:latest
WORKDIR /bin/
COPY --from=builder /go/src/github.com/KumKeeHyun/toiot/application/main .

ENTRYPOINT [ "./main" ]
