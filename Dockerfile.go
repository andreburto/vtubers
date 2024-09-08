FROM golang:1.22 AS build

WORKDIR /app

COPY ./src /app

RUN go get . && \
    go build -o vtubers . && \
    chmod +x vtubers

FROM busybox:latest

EXPOSE 8080

WORKDIR /app

COPY --from=build /app/vtubers /app/vtubers

CMD ["/app/vtubers"]
