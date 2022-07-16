FROM golang:1.18.3

WORKDIR /go/src/github.com/hamid/housingAnywhere
COPY . .

RUN go test tests/dns_test.go

RUN go build -o /docker-dns

EXPOSE 3030

CMD [ "/docker-dns" ]