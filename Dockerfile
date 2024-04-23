# stage de build
FROM golang:1.22 AS build

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/server ./cmd/server/main.go

RUN ["chmod", "+x", "/app/server"]

# stage imagem final
FROM scratch

COPY --from=build /app/server ./

EXPOSE 8000

ENTRYPOINT [ "./server" ]
