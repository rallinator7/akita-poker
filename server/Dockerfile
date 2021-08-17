FROM alpine:latest

RUN addgroup -S akita-poker-server && adduser -S akita-poker-server -G akita-poker-server

COPY ./akita-poker-server /home/akita-poker-server

WORKDIR /home/akita-poker-server

USER akita-poker-server

ENTRYPOINT ["./akita-poker-server"]