FROM alpine:3.4

RUN apk -U add ca-certificates

EXPOSE 6555
EXPOSE 5557
EXPOSE 80

ADD botserver /bin/botserver
ADD certs/ca.crt /bin/certs/ca.crt
ADD certs/client.crt /bin/certs/client.crt
ADD certs/client.key /bin/certs/client.key
ADD certs/passphrase.txt /bin/certs/passphrase.txt

CMD ["botserver", "/bin"]