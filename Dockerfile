
FROM golang:1.18 AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o runtime cmd/main.go

FROM linuxserver/ffmpeg

WORKDIR /app

COPY --from=build /app/runtime .

ENTRYPOINT [ "./runtime" ]