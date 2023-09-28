FROM golang:1.20

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .
WORKDIR cmd/app
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-monitor
WORKDIR /app

EXPOSE 8080

# Run
CMD ["/go-monitor"]