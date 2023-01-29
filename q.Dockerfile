FROM zhangshuainbb/ubuntu:20.04

WORKDIR /root

COPY ./bin/ca .
COPY ./configs ./configs
RUN chmod +x ca

CMD ["./ca", "tls", "-c", "configs/config.toml"]