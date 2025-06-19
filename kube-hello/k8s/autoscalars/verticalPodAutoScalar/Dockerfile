FROM node:18

WORKDIR /app
COPY k8s/autoscalars/horizontalPodAutoScalar/server.js .
COPY yarn.lock .
COPY package.json .
RUN yarn install
CMD ["node", "server.js"]
EXPOSE 3000