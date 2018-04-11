FROM golang:1.9

WORKDIR /go/src/github.com/ryomak/kurisuuu
COPY . .
CMD ["go","run","./cmd/main.go"]

