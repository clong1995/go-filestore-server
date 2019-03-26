## 3-1 MYSQL基础知识
### 为什么选Mysql
- Sql数据库与非SQL数据库
```
SQLServer Oracle MYSQL      MongoDB HBase Redis 
```
- Mysql的特点与优劣点
```
小
稳定
社区活跃
多平台
满足一般场景
缺高级功能 
```
- Mysql的适应场景
```
存储关系型数据 很多。。
```
### 服务架构变迁
- 用户-> 上传server -> 用户表|文件表|本地存储

### Mysql安装配置
- 安装模式
```
单点模式
主从模式
    Master DB->写日志->Bin Log
    Slave IO线程读取Bin log 
    Slave IO线程->写日志->Relay Log
    Slave RelayLog -> 读日志 -> SQL线程 -> 回放
多主模式
```
## 3-2 MYSQL主从数据同步演示
```
sudo docker ps 
sudo netstat -antup | grep docker
# 3306 3307

# Slave
mysql -uroot -h127.0.0.1 -P3307 -p

# Master
mysql -uroot -h127.0.0.1 -p
show master status;

# Slave
change master to master_host='192.168.2.244',master_user='reader','master_password='reader',master_log_file='binlog.000002',master_log_pos=0;
start slave;
show slave status\G;
# 查看Slave_IO_Running + Slave_SQL_Running = Yes

# Master
create database test1 default character set utf8;
show databases;

# Slave
show databases;

# Master 
create table tbl_test(`user` varchar(64) not null, `age` int(11) not null) default charset utf8;
show tables;

# Slave
show tables;

# Master 
insert into tbl_test(user, age) values('xiaoming',18);

# Slave 
select * from tbl_test;
```
## 3-3 文件表的设计及创建
```sql
create database fileserver default character set utf8;

use fileserver;
# 创建文件表
CREATE TABLE `tbl_file` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `file_hash` char(40) NOT NULL DEFAULT '' COMMENT '文件hash',
  `file_name` varchar(256) NOT NULL DEFAULT '' COMMENT '文件名',
  `file_size` bigint(20) DEFAULT '0' COMMENT '文件大小',
  `file_addr` varchar(1024) NOT NULL DEFAULT '' COMMENT '文件存储位置',
  `create_at` datetime default NOW() COMMENT '创建日期',
  `update_at` datetime default NOW() on update current_timestamp() COMMENT '更新日期',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '状态(可用/禁用/已删除等状态)',
  `ext1` int(11) DEFAULT '0' COMMENT '备用字段1',
  `ext2` text COMMENT '备用字段2',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_file_hash` (`file_sha1`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

show create tbl_file;
```

### Mysql 分库分表
- 水平分表
```
假设分为256张文件表
按文件hash值后两位来切分
则以: tbl_${file_hash}[:-2]的规则到对应表进行存取
```
### Golang操作Mysql
- 访问Mysql
```
使用Go标准接口
增删改查操作 
```
## 3-4 编码实战: "云存储"系统之持久化元数据到文件表
- github.com/go-sql-driver/mysql
- OnFileUploadFinished
## 3-5 编码实战: "云存储"系统之从文件表中获取元数据
- GetFileMeta

## 3-6 Ubuntu中通过Docker安装配置Mysql主从节点

## 3-7 本章小结
### 使用Mysql小结
- 使用sql.DB来管理数据库连接对象
- 通过sql.Open来创建协程安全的sql.DB对象
- 优先使用Prepared Statement

### 本章小结
- Mysql特点与应用场景
- 主从架构与文件表设计逻辑
- Golang与Mysql的亲密接触