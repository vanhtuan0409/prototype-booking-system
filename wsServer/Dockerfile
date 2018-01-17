FROM node:alpine

# Create app directory
WORKDIR /var/booking

# Install app dependencies
# A wildcard is used to ensure both package.json AND package-lock.json are copied
# where available (npm@5+)
COPY package*.json ./
COPY yarn.lock ./
RUN yarn install

# Bundle app source and build dist
COPY . .
RUN yarn run build


EXPOSE 4000
CMD ["yarn", "start"]
