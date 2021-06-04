FROM golang:1.16-alpine AS builder

ENV CGO_ENABLED=0

# Add source code
RUN apk add --update --no-cache \
        make \
        ca-certificates \
        git

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o out/gxample /src/cmd/gxample/main.go

# Multi-Stage production build
FROM alpine

VOLUME ["/usr/app/config.yml", "/usr/app/static"]

# Retrieve the binary from the previous stage
WORKDIR /usr/app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
COPY --from=builder /src/out/gxample ./
# COPY --from=builder /src/config.yml ./config.yml

# Set the binary as the entrypoint of the container
ENTRYPOINT ["/usr/app/gxample"]
CMD ["serve"]
