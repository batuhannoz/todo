FROM golang:1.20-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0

COPY . .

RUN go build main.go

FROM alpine

COPY --from=builder /app /app

CMD ["./app/main"]