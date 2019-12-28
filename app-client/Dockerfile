FROM node:alpine

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY src /app/src
COPY public /app/public

COPY package.json /app/package.json
RUN npm install -g --silent
RUN npm install react-scripts -g --silent

CMD ["npm", "start"]