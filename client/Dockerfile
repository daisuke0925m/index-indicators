# TODO alpine
FROM node:14.15.1

WORKDIR /usr/src/app/client

COPY package.json yarn.lock ./

RUN yarn install

COPY --chown=node:node . .

USER node

CMD [ "yarn", "start" ]
