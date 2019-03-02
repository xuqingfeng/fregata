FROM alpine:latest

RUN apk add --update-cache --no-cache ca-certificates && \
    update-ca-certificates

COPY out/fregata-linux-amd64 /usr/local/bin/fregata
COPY out/fregatad-linux-amd64 /usr/local/bin/fregatad
COPY etc/fregatad.conf /etc/fregata/fregatad.conf

EXPOSE 2017

ENTRYPOINT ["/usr/local/bin/fregatad"]
CMD ["-config", "/etc/fregata/fregatad.conf"]