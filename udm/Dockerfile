FROM golang:1.24.5 AS builder

WORKDIR /app
COPY ./ ./
RUN go mod download
COPY udm/cmd/main.go udm/cmd/main.go
RUN go build -ldflags "-linkmode 'external' -extldflags '-static'" -o /app/server udm/cmd/main.go

FROM scratch
WORKDIR /
COPY --from=builder /app/server /udm/server
EXPOSE 8000
ENTRYPOINT ["/udm/server"]