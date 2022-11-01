ARG GO_VERSION="1.19"
ARG UBUNTU_VERSION="22.04"

FROM docker.io/golang:${GO_VERSION} AS golang
FROM docker.io/ubuntu:${UBUNTU_VERSION} AS ubuntu

# ---

FROM golang AS build-base

RUN apt-get update -y && apt-get install -y upx

# ---

FROM build-base AS build

WORKDIR /opt/invoice-api

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -ldflags="-s -w" -o bin/ ./cmd/...
RUN upx bin/*

# ---

FROM ubuntu AS runtime-base

ENV DEBIAN_FRONTEND="noninteractive"

RUN apt-get update -y && apt-get install -y wget
RUN wget -q https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
RUN apt-get install -y ./google-chrome-stable_current_amd64.deb
RUN rm google-chrome-stable_current_amd64.deb

# ---

FROM runtime-base

COPY --from=build /opt/invoice-api/bin/web /usr/local/bin/invoice-api-web

ENV PORT=8000
EXPOSE ${PORT}

CMD [ "invoice-api-web" ]
