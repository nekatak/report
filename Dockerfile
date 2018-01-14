FROM ubuntu:latest

RUN apt-get update && apt-get install golang -y

RUN apt-get install software-properties-common -y && apt-get install git -y

RUN add-apt-repository ppa:masterminds/glide -y && apt-get update && apt-get install glide

RUN mkdir /go && mkdir /go/src && mkdir /go/bin
WORKDIR /go/src/reporter
COPY . .

ENV GOPATH=/go

RUN glide up

CMD ["go", "run", "reporter.go"]