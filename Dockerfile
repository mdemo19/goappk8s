FROM alpine
# RUN apk add -U tzdata
# RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime
COPY ./build/app-server /usr/local/bin/app-server
COPY ./rs /usr/local/bin/rs
EXPOSE 8080
CMD [ "app-server" ]