FROM golang:1.24.5 AS builder

WORKDIR /app
COPY ./ ./
RUN go mod download
COPY amf/cmd/main.go amf/cmd/main.go
RUN go build -ldflags "-linkmode 'external' -extldflags '-static'" -o /app/server amf/cmd/main.go

FROM scratch
WORKDIR /
COPY --from=builder /app/server /amf/server
EXPOSE 9010
ENTRYPOINT ["/amf/server"]