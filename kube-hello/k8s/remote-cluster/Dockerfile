FROM --platform=linux/amd64 node:18

WORKDIR /app
COPY k8s/remote-cluster/example.js .
COPY yarn.lock .
COPY package.json .
RUN yarn install
CMD ["node", "example.js"]
EXPOSE 8000