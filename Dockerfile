
FROM golang:1.20 AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o runtime cmd/main.go

FROM alpine:3.17

RUN apk add --no-cache libc6-compat

WORKDIR /app


COPY --from=build /app/runtime .

RUN adduser --system runner
RUN chown runner runtime

USER runner

ENTRYPOINT [ "./runtime" ]