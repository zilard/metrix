FROM golang:latest

ADD build/bin/metrix /bin

EXPOSE 8080

CMD ["metrix"]
