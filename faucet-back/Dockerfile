
FROM mysql:latest

# 设置root密码
ENV MYSQL_ROOT_PASSWORD=root1234
COPY init.sql /docker-entrypoint-initdb.d/init.sql

#复制文件
COPY ./config.toml /app/
COPY ./faucet-app /app/

#声明服务端口
EXPOSE 8090

# 启动MySQL服务并运行Web项目启动命令
#ENTRYPOINT ["/app/faucet-app", "/app/config.toml"]
#CMD sleep 10 && /app/faucet-app /app/config.toml