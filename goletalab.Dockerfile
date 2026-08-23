FROM golang:1.26.5
WORKDIR /app
ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=sum.golang.google.cn
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN go build ./...
