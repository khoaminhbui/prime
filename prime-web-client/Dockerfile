FROM node:10.22.0

WORKDIR /app

COPY package.json .

RUN yarn install

COPY . .

CMD ["yarn", "start"]

EXPOSE 3000