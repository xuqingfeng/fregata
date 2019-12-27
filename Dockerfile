FROM alpine:latest

RUN apk add --update-cache --no-cache ca-certificates && \
    update-ca-certificates

COPY out/fregata-linux-amd64 /usr/local/bin/fregata
COPY etc/fregata.conf /etc/fregata/fregata.conf

EXPOSE 2017

ENTRYPOINT ["/usr/local/bin/fregata"]
CMD ["-config", "/etc/fregata/fregata.conf"]