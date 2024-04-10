Faucet 是一个水龙头项目。主要功能是检查任务完成情况，并自动发放代币。后端部分用go实现，前端部分用react实现。

### 构建
    1. 构建前端项目生成 build 文件夹。
        nmp run build
    2. 构建后端项目生成目标文件 faucet-app。
        go build
    3. 将 build 文件夹复制到 faucet-app 同级目录，并重命名为 public。 

### 运行
    - 启动 faucet-app 需要指定配置文件路径
        faucet_app ./config.toml

    - 没有指定配置文件会从 FAUCET_CONF 环境变量中寻找，如果没有找到会在当前目录生成一份配置文件模板 faucet_config.toml。

### 主页路径
host/faucet/