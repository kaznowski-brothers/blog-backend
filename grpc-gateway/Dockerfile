FROM golang:1.12
RUN mkdir -p /go/src/github.com/kaznowski-brothers/blog-backend/grpc-gateway
ADD . /go/src/github.com/kaznowski-brothers/blog-backend/grpc-gateway/
WORKDIR /go/src/github.com/kaznowski-brothers/blog-backend/grpc-gateway
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 8080
CMD /go/bin/grpc-gateway --port :8080 --endpoint grpc-service:10000
