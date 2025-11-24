# go
# --- local ---
FROM golang:1.25-alpine as build-go
WORKDIR /go/server
COPY . .
RUN apk add --no-cache git make && \
  go install github.com/air-verse/air@latest && \
  make build

# --- production ---
FROM alpine as production-go
WORKDIR /app
COPY --from=build-go /go/server/bin/server .
COPY --from=build-go /go/server/list.json .
RUN addgroup go \
  && adduser -D -G go go \
  && chown -R go:go /app/server
EXPOSE 8080
CMD ["./server"]
