FROM node:20-alpine AS web

WORKDIR /buildApp

COPY ./web .
RUN npm i
RUN npm run build

FROM golang:1.21-bullseye AS build
WORKDIR /build
COPY . .
RUN apt-get update && apt-get install libopus-dev -y
ENV CGO_ENABLED=1
RUN go mod download
RUN go build -o ./bin/main cmd/main.go
RUN go build -o ./bin/server cmd/server.go

FROM alpine:3 AS final
WORKDIR /app
COPY --from=web /buildApp/build ./static
COPY --from=build /build/bin .
ENV KLAMMERAEFFCHEN_BOTTOKEN=token
ENV KLAMMERAEFFCHEN_BOTCLIENTID=id
ENV KLAMMERAEFFCHEN_CLIENTSECRET=secret
ENV KLAMMERAEFFCHEN_SERVERPORT=3000
ENV KLAMMERAEFFCHEN_OAUTHREDIRECT=redirect
ENV KLAMMERAEFFCHEN_DASHBOARDURI=dashboard

EXPOSE 3000
ENTRYPOINT ["sh", "./docker.sh"]