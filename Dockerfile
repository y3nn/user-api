# --- Stage 1: Builder Stage ---
FROM golang:1.24.3 AS builder

# Set working directory in the container
WORKDIR /app 

# Copy go.mod and go.sum files and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download 

# Copy all other source code and directories
COPY . .
# Build the executable. CGO_ENABLED=0 for static build, -o app sets the output filename
RUN CGO_ENABLED=0 go build -o app ./cmd

# --- Stage 2: Final Image ---
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/app /app/app

# Expose port (informational only, does not open it)
EXPOSE 8080
# Command to run when the container starts
CMD [ "/app/app" ]
