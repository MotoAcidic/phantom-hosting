FROM golang:latest

WORKDIR /go/src/phantom-hosting
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build

EXPOSE 8000

CMD ["./phantom-hosting"]