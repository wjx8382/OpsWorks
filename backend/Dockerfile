FROM golang:alpine
WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct && \
    go env
RUN go mod download
RUN go mod verify
RUN go build -o main .
CMD ["nohup", "./main", "--bind-address=0.0.0.0", ">/dev/null", "2>&1", "&"]