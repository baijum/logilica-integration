FROM golang:1.21 as builder

RUN apt-get update -y
RUN apt-get install -y golang

WORKDIR /workspace

COPY / /workspace/
RUN go build

FROM debian:12
WORKDIR /
COPY --from=builder /workspace/logilica-integration .
EXPOSE 8080
ENTRYPOINT ["/logilica-integration"]
