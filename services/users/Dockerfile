FROM golang:1.17-buster as builder

WORKDIR /onefence
COPY ../.. /onefence

RUN CGO_ENABLED=0 go build -o /bin/app .

FROM golang:1.17-buster

WORKDIR /onefence

COPY --from=builder /bin/app /onefence/app

EXPOSE 8080

ENTRYPOINT ["/onefence/app"]