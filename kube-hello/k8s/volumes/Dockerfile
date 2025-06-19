FROM node:18

WORKDIR /app
COPY k8s/volumes/query-logger.js .
COPY yarn.lock .
COPY package.json .
RUN yarn install
CMD ["node", "query-logger.js"]
EXPOSE 8000