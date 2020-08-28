# base image
FROM node:12.18.3 as builder

# set working directory
WORKDIR /app

# `/app/node_modules/.bin`을 $PATH 에 추가
ENV PATH /app/node_modules/.bin:$PATH

# app dependencies, install 및 caching
COPY package.json ./
COPY package-lock.json ./
RUN npm install --silent 
RUN npm install react-scripts@3.4.3 -g --silent

# add app
COPY . ./

# start app
CMD ["npm", "start"]