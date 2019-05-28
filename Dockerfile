FROM golang:latest

WORKDIR $GOPATH/src/github.com/kbrimm/tic-tac-toe
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go get github.com/stretchr/testify/assert

CMD ["tic-tac-toe"]
