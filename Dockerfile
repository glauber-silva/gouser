from golang:1.21-alpine

RUN apk add --no-cache mysql-client

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o user-service ./cmd/user-service
RUN chmod +x user-service

COPY scripts/wait-for-mysql.sh /app/wait-for-mysql.sh
RUN chmod +x /app/wait-for-mysql.sh
EXPOSE 8080

CMD ["/app/wait-for-mysql.sh", "mysql", "./user-service"]

