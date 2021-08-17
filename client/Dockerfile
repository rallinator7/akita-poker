FROM node:lts-alpine3.14

RUN addgroup -S akita-poker-client && adduser -S akita-poker-client -G akita-poker-client

COPY . /home/akita-poker-client

WORKDIR /home/akita-poker-client

USER akita-poker-client

ENTRYPOINT ["node", "dist/index.js"]