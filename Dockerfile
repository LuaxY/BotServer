FROM alpine:3.4

RUN apk -U add ca-certificates

EXPOSE 6555
EXPOSE 5557
EXPOSE 80

ADD botserver /bin/botserver
COPY certs /bin/certs
COPY static /bin/static

CMD ["botserver", "/bin"]