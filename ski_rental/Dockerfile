FROM golang:alpine

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["ski_rental"]
