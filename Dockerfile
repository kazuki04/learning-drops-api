FROM golang:1.17.2-alpine3.14
COPY ./ /app/go
WORKDIR /app/go
RUN go build -o main . 
CMD ["/app/go/main"]
