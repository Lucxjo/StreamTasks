FROM golang:1.17-alpine as base
RUN apk --update --no-cache add bash build-base ca-certificates git

WORKDIR /app
COPY . .

ENV GOOS=linux
ENV CGO_ENABLED=1

RUN go build --ldflags="-s -w" -o ./stasc-svr


FROM alpine
WORKDIR /app
COPY --from=base /app/stasc-svr ./stasc-svr
COPY --from=base /app/public ./public
EXPOSE 8080

CMD ["./stasc-svr"]

LABEL org.opencontainers.image.source https://github.com/Lucxjo/StreamTasks
