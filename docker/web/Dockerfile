FROM node:23.1-bullseye

WORKDIR /

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 3000
