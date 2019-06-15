FROM golang:1.12 as builder
RUN mkdir -p /go/src/github.com/kaznowski-brothers/blog-backend/grpc-service
ADD . /go/src/github.com/kaznowski-brothers/blog-backend/grpc-service/
WORKDIR /go/src/github.com/kaznowski-brothers/blog-backend/grpc-service
RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build -v -o grpc-service .

FROM alpine
COPY --from=builder /go/src/github.com/kaznowski-brothers/blog-backend/grpc-service/grpc-service /app/
EXPOSE 10000
WORKDIR /app
CMD ./grpc-service --port 10000
