FROM golang:1.10-alpine3.8 as builder

WORKDIR /go/src/app

COPY go-hpa.go .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w"

FROM scratch

WORKDIR /go/src/app

COPY --from=builder /go/src/app/app .

CMD ["./app"]