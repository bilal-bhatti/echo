FROM scratch

ADD "https://curl.haxx.se/ca/cacert.pem" "/etc/ssl/certs/ca-certificates.crt"
ADD "./dist/echo-linux-amd64" "/echo"
ENTRYPOINT ["/echo"]
