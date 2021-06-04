# syntax = docker/dockerfile:1.2

FROM golang:1.15-alpine AS build

WORKDIR /src

ENV TZ=Asia/Seoul
RUN apk add --update --no-cache tzdata

COPY go.* ./
RUN --mount=type=cache,target=/root/.cache/go-build \
	go mod download \
	&& go get -u github.com/cosmtrek/air

COPY . .

ENTRYPOINT ["air"]

