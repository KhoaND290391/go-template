FROM  golang:1.13-alpine as builder
RUN mkdir -p /app && rm -rf /app/*
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
FROM alpine:latest
EXPOSE 80
WORKDIR /root/
COPY --from=builder /app/bin/* .
RUN ls
ENTRYPOINT ["./main"]