FROM golang:latest
LABEL Fisher="circlemono@outlook.com"

# 设置go mod proxy 国内代理
# 设置golang path
ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
    GO111MODULE=on \
    CGO_ENABLED=1
WORKDIR /admin
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
    && go list \
    && go build -o app main.go

EXPOSE 8081
ENTRYPOINT /admin/app