FROM golang:1.24.5 AS builder

WORKDIR /app
COPY ./ ./
RUN go mod download
COPY upf/cmd/main.go upf/cmd/main.go
RUN go build -ldflags "-linkmode 'external' -extldflags '-static'" -o /app/server upf/cmd/main.go

FROM scratch
WORKDIR /
COPY --from=builder /app/server /upf/server
EXPOSE 8805/udp
ENTRYPOINT ["/upf/server"]