FROM golang:1.22-alpine as build
COPY . /app

WORKDIR /app
RUN go build -o truthordare .

FROM alpine:3
COPY --from=build /app/truthordare /app/truthordare
COPY ./views /app/views
WORKDIR /app
CMD ["/app/truthordare"]
EXPOSE 1323
