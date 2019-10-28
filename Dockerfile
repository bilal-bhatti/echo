FROM scratch

ADD "https://curl.haxx.se/ca/cacert.pem" "/etc/ssl/certs/ca-certificates.crt"
ADD "./dist/http-echo-linux-amd64" "/http-echo"
ENTRYPOINT ["/http-echo"]
