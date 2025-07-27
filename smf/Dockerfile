FROM golang:1.24.5 AS builder

WORKDIR /app
COPY ./ ./
RUN go mod download
COPY smf/cmd/main.go smf/cmd/main.go
RUN go build -ldflags "-linkmode 'external' -extldflags '-static'" -o /app/server smf/cmd/main.go

FROM scratch
WORKDIR /
COPY --from=builder /app/server /smf/server
EXPOSE 40
ENTRYPOINT ["/smf/server"]