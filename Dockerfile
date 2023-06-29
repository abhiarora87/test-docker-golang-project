FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go build -o server .

EXPOSE 3030 

CMD ["./server"]