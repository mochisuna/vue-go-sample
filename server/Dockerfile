# --- local ---
# go
FROM golang:1.12-alpine as build-go
WORKDIR /go/server
COPY . .
RUN apk add --no-cache git make  && \
  go get github.com/oxequa/realize && \
  make build

# --- production ---
# go
FROM alpine as production-go
WORKDIR /app
COPY --from=build-go /go/server/bin/server .
COPY --from=build-go /go/server/list.json .
RUN addgroup go \
  && adduser -D -G go go \
  && chown -R go:go /app/server
EXPOSE 8080
CMD ["./server"]
