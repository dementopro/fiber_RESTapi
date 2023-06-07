FROM golang:1.20.4 AS GO

WORKDIR /opt/toc

ADD . /opt/toc

RUN CGO_ENABLED=0 go build -o toc /opt/toc/cmd/main.go

FROM alpine:3.16.2

WORKDIR /

COPY --from=GO /opt/toc/toc .
COPY config/config.yaml ./config/

CMD ["./toc"]

EXPOSE 7000