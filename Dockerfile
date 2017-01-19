From golang

RUN go get -u github.com/gorilla/mux
ADD . /go/src/github.com/iheanyi/switch-bot
RUN go install github.com/iheanyi/switch-bot
CMD ["/go/bin/switch-bot"]
EXPOSE 8080
