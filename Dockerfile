# Set the base image to the latest version of Golang
FROM golang:latest

# Set the working directory inside the container to /go/src/app
WORKDIR /go/src/app

# Copy all the files from the current directory to the /go/src/app directory inside the container
COPY . .

# Build the Go application inside the container and create an executable named "server"
RUN go build -o server .

# Expose port 3030 of the container for incoming network connections
EXPOSE 3030

# Specify the command to run when the container starts, executing the "server" executable
CMD ["./server"]