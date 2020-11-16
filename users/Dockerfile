FROM golang

RUN mkdir -p /go/src/go-kit-demo
RUN git config --global https.proxy http://127.0.0.1:1080
RUN git config --global https.proxy https://127.0.0.1:1080
ADD . /go/src/go-kit-demo


RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run go-kit-demo/users/cmd  -watch go-kit-demo/users
