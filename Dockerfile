FROM golang:1.21-bookworm AS builder
ENV WORKDIR /app
WORKDIR $WORKDIR

COPY . $WORKDIR

RUN go mod download
RUN go install github.com/cespare/reflex
RUN go build -o /wiki ${WORKDIR}/cmd/nullwiki/wiki.go 

FROM debian:bookworm-slim AS release
WORKDIR /

COPY --from=builder /wiki /wiki

EXPOSE 8080

CMD ["/wiki"]
