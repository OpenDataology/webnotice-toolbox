FROM golang:alpine
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o webnotice_toolbox ./webnotice-scanner/main.go
CMD ["./webnotice_toolbox"]