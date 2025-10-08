FROM golang:1.24

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o ./build/authservice ./cmd/authentication
# RUN go test ./...

EXPOSE 3001

ENTRYPOINT ["./build/authservice"]
