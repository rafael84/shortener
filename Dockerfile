FROM alpine:3.5
ADD target/shortener-server /bin
CMD /bin/shortener-server
