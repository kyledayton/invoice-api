ARG GO_VERSION="1.23"
ARG BUSYBOX_VERSION="1.36"

FROM docker.io/golang:${GO_VERSION} AS golang
FROM docker.io/busybox:${BUSYBOX_VERSION} AS busybox

# ---

FROM golang AS build

WORKDIR /opt/invoice-api

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -ldflags="-s -w" -o bin/ ./cmd/...

# ---

FROM busybox

COPY --from=build /opt/invoice-api/bin/web /usr/local/bin/invoice-api-web

ENV PORT=8000
EXPOSE ${PORT}

CMD [ "invoice-api-web" ]
