FROM golang:1.22 AS build

WORKDIR /app

COPY ./test /app

RUN go get . && \
    go build -o vtubers . && \
    chmod +x vtubers

FROM busybox:latest

COPY --from=build /app/vtubers /app/vtubers

ENTRYPOINT ["/app/vtubers"]
