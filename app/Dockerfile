# Create build stage based on buster image
FROM golang:1.16-buster AS builder

WORKDIR /app

COPY go.* .

RUN go mod download

COPY ./ ./

RUN go build -o /hello_go_http

RUN chmod +x /hello_go_http

EXPOSE 8080

ENTRYPOINT ["/hello_go_http"]