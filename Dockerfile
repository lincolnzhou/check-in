FROM golang:1.10.1-alpine3.7 as golang

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk update && \
    apk add  bash  && \ 
    rm -rf /var/cache/apk/*   /tmp/*

WORKDIR /go/src/github.com/lincolnzhou/check-in

ADD ./backend ./backend
ADD control .
RUN ./control build

FROM node:9.11.1-alpine as node
ADD ./frontend/ /data/check-in 
WORKDIR /data/check-in
RUN npm install && npm run build

FROM alpine:3.7
MAINTAINER LincolnZhou "875199116@qq.com"
ENV TZ='Asia/Shanghai' 
RUN TERM=linux && export TERM
USER root 
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk update && \
    apk add bash && \ 
    echo "Asia/Shanghai" > /etc/timezone && \
    rm -rf /var/cache/apk/*   /tmp/*

WORKDIR /data/check-in
ADD control /data/check-in/control
COPY --from=golang /go/src/github.com/lincolnzhou/check-in/backend/backend /data/check-in/backend/backend
COPY --from=golang /go/src/github.com/lincolnzhou/check-in/backend/config.toml /data/check-in/backend/config.toml
COPY --from=node /data/check-in/static /data/check-in/backend/static

CMD ["./control", "rundocker"]
