FROM golang:1.6

RUN go get github.com/codegangsta/gin
RUN go get github.com/mattes/migrate
ADD . /go/src/github.com/mefellows/onegeek-ftg-incident-management
WORKDIR /go/src/github.com/mefellows/onegeek-ftg-incident-management
RUN go get
