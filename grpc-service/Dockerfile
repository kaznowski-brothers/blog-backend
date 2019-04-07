FROM golang:1.12
RUN mkdir -p /go/src/github.com/kaznowski-brothers/blog-backend/grpc-service
ADD . /go/src/github.com/kaznowski-brothers/blog-backend/grpc-service/
WORKDIR /go/src/github.com/kaznowski-brothers/blog-backend/grpc-service
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 10000
CMD /go/bin/grpc-service --port 10000