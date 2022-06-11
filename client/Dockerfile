FROM node:16.14.2-alpine as build-stage

RUN mkdir /usr/app

COPY . /usr/app

WORKDIR /usr/app

RUN npm install

ENV PATH usr/src/app/node_modules/.bin:$PATH
ENV REACT_APP_BACKEND http://host.docker.internal:3030

RUN npm run build

# ADD NGINX

FROM nginx:alpine

WORKDIR /usr/share/nginx/html

RUN rm -rf ./* 

COPY --from=build-stage /usr/app/build .

ENTRYPOINT ["nginx" , "-g" , "daemon off ;"]