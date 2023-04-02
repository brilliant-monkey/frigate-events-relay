
FROM golang:1.20 AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o runtime cmd/main.go

FROM alpine:3.17

WORKDIR /app

COPY --from=build /app/runtime .

RUN adduser --system runner
RUN chown runner runtime

USER runner

ENTRYPOINT [ "./runtime" ]