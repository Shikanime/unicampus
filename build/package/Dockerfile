FROM golang:1.11.4 AS builder
WORKDIR /opt/build
COPY ./cmd/admission /opt/build
RUN make build-admission

FROM golang:1.11.4-alpine
WORKDIR /opt/app
COPY --from=builder /opt/built/build-admission .
CMD ["/opt/app/build-admission"]
