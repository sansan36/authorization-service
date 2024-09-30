ARG GOLANG_TAG=1.22.6-alpine
FROM golang:${GOLANG_TAG}

WORKDIR /app

COPY go.mod go.sum .env ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./app

EXPOSE 8081
EXPOSE 9091

CMD ["./main"]
