FROM node:20-alpine AS web

WORKDIR /buildApp

COPY ./web .
RUN npm i
RUN npm run build

FROM golang:1.21 AS build
WORKDIR /build
COPY . .
RUN apt-get update && apt-get install libopus0 libopus-dev -y
RUN CGO_ENABLED=1 GOOS=linux go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o ./bin/main cmd/main.go
RUN CGO_ENABLED=1 GOOS=linux go build -o ./bin/server cmd/server.go

FROM alpine:3 AS final
WORKDIR /app
COPY --from=web /buildApp/build ./static
COPY --from=build /build/bin/main .
COPY --from=build /build/bin/server .
RUN apk add libc6-compat opus ffmpeg
COPY ./docker.sh .
RUN chmod +x ./docker.sh
ENV KLAMMERAEFFCHEN_BOTTOKEN=token
ENV KLAMMERAEFFCHEN_BOTCLIENTID=id
ENV KLAMMERAEFFCHEN_CLIENTSECRET=secret
ENV KLAMMERAEFFCHEN_SERVERPORT=3000
ENV KLAMMERAEFFCHEN_OAUTHREDIRECT=redirect
ENV KLAMMERAEFFCHEN_DASHBOARDURI=dashboard

EXPOSE 3000
CMD ["./docker.sh"]