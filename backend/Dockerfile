FROM golang:alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    pghost=postgres

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main

WORKDIR /dist

RUN cp /build/main .

# FROM scratch

# COPY --from=builder /dist/main /

CMD ["/dist/main"]

# CMD ["sh"]