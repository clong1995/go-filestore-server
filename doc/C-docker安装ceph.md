# docker安装ceph
## 0.准备工作
``` 
# 创建准备目录,挂载volume
cd 
mkdir -p docker/ceph
mkdir -p docker/var/lib/ceph
chown -R 64045:64045 docker/var/lib/ceph/osd
chown -R 64045:64045 docker/osd

# 创建ceph-network
docker network create --driver bridge --subnet 172.20.0.0/16 ceph-network
docker network ls
docker network inspect ceph-network
```
## 1.拉取ceph/mon镜像到本地
```
docker pull ceph/daemon
```

## 2.运行ceph
``` 

docker run -itd --net=host --name=mon -v /Users/xx/docker/ceph:/etc/ceph -v  /Users/xx/docker/var/lib/ceph/:/var/lib/ceph   -e MON_IP=192.168.8.106 -e CEPH_PUBLIC_NETWORK=192.168.0.0/16 ceph/daemon mon

```