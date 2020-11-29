FROM node:14.15.1

WORKDIR /usr/src/app/client

COPY package.json yarn.lock ./

RUN yarn install

COPY . .

RUN chown -R node:node /usr/src/app/client && \
    chmod -R 744 /usr/src/app/client && \
    chown -R node:node /usr/src/app/client/node_modules && \
    chmod -R 777 /usr/src/app/client/node_modules

# コンテナを起動した時のプロセスのOwnerを指定  <---- changed
USER node

CMD [ "yarn", "start" ]