FROM node
WORKDIR /app/
COPY package.json /app/package.json
# RUN npm install --force -g yarn
# RUN yarn install
RUN npm install
COPY src /app/src
COPY public /app/public
EXPOSE 3000
ENTRYPOINT [ "npm", "start" ]