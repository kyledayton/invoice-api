ARG GO_VERSION="1.19"
ARG UBUNTU_VERSION="22.04"

FROM docker.io/golang:${GO_VERSION} AS golang
FROM docker.io/ubuntu:${UBUNTU_VERSION} AS ubuntu

# ---

FROM golang AS build-base

# ---

FROM build-base AS build

WORKDIR /opt/invoice-api

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -ldflags="-s -w" -o bin/ ./cmd/...

# ---

FROM ubuntu AS runtime-base

RUN apt-get update -y && apt-get install -y software-properties-common
RUN add-apt-repository ppa:xtradeb/apps -y && apt-get update -y && apt-get install -y chromium

# ---

FROM runtime-base

COPY --from=build /opt/invoice-api/bin/web /usr/local/bin/invoice-api-web

ENV PORT=8000
ENV CHROME_EXECUTABLE=chromium
EXPOSE ${PORT}

CMD [ "invoice-api-web" ]
