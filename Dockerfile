from golang:latest AS builder

workdir /app

copy go.mod ./

# run go.mod download

copy . . 

RUN CGO_ENABLED=0 GOOS=linux go build -o s3 .

from alpine:latest

EXPOSE 8686

COPY --from=builder /app/s3 /usr/local/bin/s3
ENTRYPOINT ["/usr/local/bin/s3"]
