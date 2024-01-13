# First stage: build the Go binary
FROM golang:1.21 as builder

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy the Go modules and sum files
COPY . .

# Build the Go app
RUN go build -o app .

# Second stage: build the final image
FROM golang:1.21

# Set the working directory
WORKDIR /usr/src/app

# Copy the binary from the builder stage
COPY --from=builder /usr/src/app/app .

# Copy the index.html into the container
COPY index.html .

# Run the binary
CMD ["./app"]
