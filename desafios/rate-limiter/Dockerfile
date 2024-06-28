FROM golang:1.22.4-alpine3.20 AS builder


WORKDIR /app

COPY . . 


RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o bin/api cmd/main.go

FROM scratch
COPY --from=builder /app/bin/api .
COPY --from=builder /app/.env .
# COPY /app/.env ./cmd


CMD ["./api"]
