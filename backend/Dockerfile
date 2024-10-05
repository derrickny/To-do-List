# Dockerfile for backend (backend/Dockerfile)

# Use the official Go image as a base, version 1.19
FROM golang:1.19

# Set the working directory in the container to /backend
WORKDIR /backend

# Copy the current directory (i.e., the directory containing the Dockerfile)
# into the container at the working directory
COPY . .

# Run the command `go mod tidy` to tidy up the go.mod file and remove any
# unnecessary dependencies
RUN go mod tidy

# Run the command `go build -o main .` to build the Go program in the current
# directory and output the resulting executable to a file named "main"
RUN go build -o main .

# Verify that the main binary exists
RUN ls -l main

# List the contents of the /backend directory for debugging
RUN ls -l /backend

# Inform Docker that the container listens on port 8080
EXPOSE 8080

# Set the default command to run when the container starts to "./main", which
# runs the executable produced by the previous step
CMD ["./main"]