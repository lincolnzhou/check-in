FROM golang:1.10.1-alpine3.7 as golang
MAINTAINER LincolnZhou "875199116@qq.com"

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk update && \
    apk add  bash  && \ 
    rm -rf /var/cache/apk/*   /tmp/*

WORKDIR /go/src/github.com/lincolnzhou/check-in

ADD . .
RUN ./control build

CMD ["./control", "rundocker"]
