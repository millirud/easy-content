FROM node:18.16.0-alpine3.17 AS base

WORKDIR /home/app

COPY package*.json ./

RUN npm install && npm cache clean --force

COPY . .

ENV HOST=0.0.0.0 \
    PORT=3000


FROM base AS dev
CMD ["npm", "run", "dev"]

FROM base AS prod
CMD ["npm", "run", "prod"]

