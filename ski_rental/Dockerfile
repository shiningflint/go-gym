FROM golang:alpine

WORKDIR /go/src/github.com/shiningflint/go-gym/ski_rental

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["ski_rental"]
