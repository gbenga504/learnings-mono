FROM --platform=linux/amd64 node:18

WORKDIR /app
COPY log-server.js .
COPY node_modules ./node_modules
COPY ./package.json .
COPY yarn.lock .
RUN yarn install
CMD ["node", "log-server.js"]
EXPOSE 8000