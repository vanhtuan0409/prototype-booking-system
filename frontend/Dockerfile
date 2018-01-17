FROM node:alpine

ARG REACT_APP_API_URL
ARG REACT_APP_SOCKET_URL

# Create app directory
WORKDIR /var/booking

# Install app dependencies
# A wildcard is used to ensure both package.json AND package-lock.json are copied
# where available (npm@5+)
COPY package*.json ./
COPY yarn.lock ./
RUN yarn install
RUN yarn global add serve

# Bundle app source and build dist
COPY . .
RUN yarn run build

EXPOSE 3000
CMD serve -s -p 3000 build
