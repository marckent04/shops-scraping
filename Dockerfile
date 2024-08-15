FROM node:22-alpine AS front-build

WORKDIR app

COPY front .

RUN yarn && yarn build


FROM golang:1.22-alpine

WORKDIR app

COPY . .

COPY docker.env .env

RUN go install

RUN apk add chromium g++

COPY --from=front-build app/dist front/dist

CMD go run main.go

EXPOSE 4000


