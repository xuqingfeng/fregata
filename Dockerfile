FROM alpine:3.5

RUN apk --update --no-cache add ca-certificates

COPY out/fregata-linux-amd64 /usr/local/bin/fregata
COPY out/fregatad-linux-amd64 /usr/local/bin/fregatad
COPY etc/fregata.conf /etc/fregata/fregata.conf

EXPOSE 2017

ENTRYPOINT ["/usr/local/bin/fregatad"]
CMD ["-config", "/etc/fregata/fregata.conf"]