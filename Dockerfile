FROM golang:1.23-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go test -v ./utest .

RUN go build -o main .

FROM alpine:3.17

WORKDIR /app

COPY --from=builder /app .

EXPOSE 8080

CMD ["./main"]