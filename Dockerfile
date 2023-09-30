# Use an official Golang runtime as a parent image
FROM golang:1.20

# Set the working directory to /db
WORKDIR /db

# Copy the current directory contents into the container at /db
COPY . /db

# Build the Go db
# RUN go mod init we-crawler
RUN go mod tidy
RUN go build -o connect .

# Run the Go db
CMD ["/db/connect"]
