FROM golang:1.24.13-alpine3.22 AS build

COPY . /app
WORKDIR /app
RUN go build -o gopen-api

FROM alpine:3.22 AS run
COPY --from=build /app/gopen-api /app/gopen-api
WORKDIR /app
CMD ["./gopen-api"]
