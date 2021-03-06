FROM golang:1.16-alpine AS builder

ENV CGO_ENABLED=0

# Add source code
RUN apk add --update --no-cache \
        make \
        ca-certificates \
        git

WORKDIR /src

COPY go.mod go.sum tools.go Makefile ./
RUN make install-tools

COPY . .
RUN make gendoc
RUN go build -o ./out/gxample ./cmd/gxample

# Multi-Stage production build
FROM alpine

VOLUME ["/usr/app/config.yml", "/usr/app/static","/usr/app/docs"]

# Retrieve the binary from the previous stage
ENV TZ=Asia/Seoul
RUN apk add --update --no-cache tzdata
WORKDIR /usr/app
COPY --from=builder /src/docs/swagger.json ./docs/api-docs
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
COPY --from=builder /src/out/gxample ./

# Set the binary as the entrypoint of the container
ENTRYPOINT ["/usr/app/gxample"]
CMD ["serve"]
