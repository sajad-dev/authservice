FROM golang:1.24

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./xds
# RUN go test ./...

EXPOSE 8082

ENTRYPOINT ["./xds"]

