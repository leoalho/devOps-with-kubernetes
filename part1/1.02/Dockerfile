FROM golang:1.15 as builder
ARG CGO_ENABLED=0
WORKDIR /app
COPY . .
RUN go build

FROM scratch
COPY --from=builder /app /app
ENTRYPOINT ["/app/server"]
