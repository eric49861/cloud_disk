# cloud_disk
基于go-zero实现一个简易的云盘系统

## 说明
- 和该项目相关的文档都在文档说明文件夹下
    - 需求文档.md
    - 数据库设计.md
    - 接口设计.md
    - 工程相关.md
    - 业务逻辑
- 给项目添加配置
    在运行项目之前需要添加配置文件
  - 网关api目录下，需要创建etc目录，并在其中添加clouddisk-api.yaml文件
    ```yaml
    Name: clouddisk-api
    Host: 0.0.0.0
    Port: 8888

    UserModule:
    Etcd:
    Hosts:
    - ETC serverIP:2379
    Key: user.rpc
    ```
  - 在rpc目录的各个服务目录下创建etc目录，并在其中添加响应的配置文件
    
    比如在user目录下创建etc/user.yaml
    ```yaml
    Name: user.rpc
    ListenOn: 0.0.0.0:8080

    MysqlConfig:
    dsn: mysql连接配置

    RedisConfig:
    addr: RedisServerIP:6379
    password: RedisPass
    db: 0

    Etcd:
    Hosts:
    - ETCDServerIP:2379
    Key: user.rpc
    ```
 - 在define目录下需要添加自己的邮箱授权码