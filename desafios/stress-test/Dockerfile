FROM golang:1.22.4-alpine3.20 AS builder


WORKDIR /app

COPY . . 

RUN go mod download
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o bin/cli ./main.go

FROM scratch
COPY --from=builder /app/bin/cli .

ENTRYPOINT ["./cli"]
