FROM golang
WORKDIR /app/
COPY . .
RUN go mod download
EXPOSE 8090
ENTRYPOINT [ "go", "run", "server.go"]