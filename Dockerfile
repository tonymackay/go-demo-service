FROM golang:1.15.5-buster AS builder
ARG VERSION=dev
WORKDIR /build
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -o demoservice \
  -ldflags="-w -X=main.version=${VERSION}" main.go

FROM scratch
COPY --from=builder /build/demoservice /demoservice
ENTRYPOINT [ "/demoservice" ]
