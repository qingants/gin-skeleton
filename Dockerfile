# docker build -t ibroomcorn/gin-skeleton .
# docker run --rm --name upload ibroomcorn/gin-skeleton


FROM alpine:latest

MAINTAINER ibroomcorn@gmail.com

RUN apk add tzdata \
&& ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&& echo "Asia/Shanghai" > /etc/timezone

#ENV TZ=Asia/Shanghai

WORKDIR /srv/app/

COPY ./gin-skeleton /srv/app/

#COPY ../../conf/gin-skeleton.ini /srv/app/conf/gin-skeleton.ini

#EXPOSE 8091

ENTRYPOINT ["./gin-skeleton",  "-f",  "./conf/conf.ini"]