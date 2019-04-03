# centos安装docker
## 1.yum配置阿里云的源
```
cd /etc/yum.repos.d/

#下载阿里云yum源
wget http://mirrors.aliyun.com/repo/Centos-7.repo
mv CentOS-Base.repo CentOS-Base.repo.bak
mv Centos-7.repo CentOS-Base.repo
```
## 2.重置yum源
``` 
yum clean all
yum makecache
```

## 3.开始安装docker
``` 
yum list docker-ce
yum -y install docker-ce
docker -v
systemctl start docker
docker info
```