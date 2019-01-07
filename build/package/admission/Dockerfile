FROM golang:1.11.4 AS builder
WORKDIR /opt/build
COPY ./cmd/admission /opt/build
RUN go build -o main .

FROM golang:1.11.4-alpine
WORKDIR /opt/app
COPY --from=builder /opt/built .
CMD ["/opt/app/main"]
