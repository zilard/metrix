FROM ubuntu:16.04

ADD build/bin/metrix /bin

EXPOSE 8080

CMD ["metrix"]
