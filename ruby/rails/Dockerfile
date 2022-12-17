FROM ruby:3.0.3-alpine3.15

WORKDIR /app

RUN apk add --no-cache linux-headers libxml2-dev make gcc libc-dev nodejs tzdata
RUN apk add --virtual build-packages --no-cache build-base curl-dev

