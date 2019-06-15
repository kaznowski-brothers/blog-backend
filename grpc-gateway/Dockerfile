FROM golang:1.12 as builder
RUN mkdir -p /go/src/github.com/kaznowski-brothers/blog-backend/grpc-gateway
ADD . /go/src/github.com/kaznowski-brothers/blog-backend/grpc-gateway/
WORKDIR /go/src/github.com/kaznowski-brothers/blog-backend/grpc-gateway
RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build -v -o grpc-gateway .

FROM alpine
COPY --from=builder /go/src/github.com/kaznowski-brothers/blog-backend/grpc-gateway/grpc-gateway /app/
EXPOSE 8080
WORKDIR /app
CMD ./grpc-gateway --port :8080 --endpoint grpc-service:10000
