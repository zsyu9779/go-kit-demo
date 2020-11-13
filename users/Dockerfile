FROM golang

RUN mkdir -p /go/src/go-kit-demo

ADD . /go/src/go-kit-demo

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run go-kit-demo/users/cmd  -watch go-kit-demo/users
