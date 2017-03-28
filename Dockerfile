FROM alpine:3.5

MAINTAINER https://github.com/xuqingfeng

COPY out/fregated /usr/local/bin/fregated
COPY out/fregate /usr/local/bin/fregate
COPY etc/fregate.conf /etc/fregate.conf

EXPOSE 2017

ENTRYPOINT ["/usr/local/bin/fregated"]
CMD ["-config", "/etc/fregate.conf"]

