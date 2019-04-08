# docker安装rabbitMq
## 0.准备工作
```
cd 
mkdir -p docker/rabbitmq 
```
## 1.拉取镜像
``` 
docker pull rabbitmq
```

## 2.安装rabbitMq
``` 
docker run -d --name myrabbitmq -p 5672:5672 -p 15672:15672 -v /Users/xx/docker/rabbitmq:/var/lib/rabbitmq --hostname myRabbit -e RABBITMQ_DEFAULT_VHOST=my_vhost  -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=admin rabbitmq

```
## 