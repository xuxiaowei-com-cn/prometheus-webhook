FROM alpine:3.17.3

# 维护人员
MAINTAINER 徐晓伟 xuxiaowei@xuxiaowei.com.cn

WORKDIR /usr/local/sbin

# 添加 普罗米修斯Webhook
ADD prometheus_webhook .

WORKDIR /root

# 测试版本号
RUN prometheus_webhook --version

# 端口
EXPOSE 8080

CMD prometheus_webhook
