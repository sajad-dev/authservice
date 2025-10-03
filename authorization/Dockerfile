FROM golang:1.24

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./build/authservice ./cmd/authorize
# RUN go test ./...

EXPOSE 3000

ENTRYPOINT ["./build/authservice"]
