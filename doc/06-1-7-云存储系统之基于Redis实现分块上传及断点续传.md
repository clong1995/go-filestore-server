## 6-1 分块上传与断点续传原理
### 分块上传与断点续传
- 两个概念
``` 
- 分块上传: 文件切成多块，独立传输，上传完成后合并
- 断点续传: 传输暂停或异常中断后，可基于原来进度重传

```
- 几点说明
```
- 小文件不建议分块上传
- 可以并行上传分块，并且可以无序传输
- 分块上传能极大提高传输效率
- 减少传输失败后重试的流量及时间
```
- 具体流程
``` 
1. 初始化上传        Initiate Multipart Upload
2. 上传分块（并行)    Upload Part -> Upload Abort
3. 通知上传完成       Complete Multipart Upload  Upload Query
```
### 服务架构变迁
- 用户->分块上传
- 分块上传<->本地存储
- 分块上传->Redis缓存|Hash计算
- 分块上传->用户文件表
- 分块上传->唯一文件表

## 6-2 编码实战: Go实现Redis连接池（存储分块信息）
### 分块上传通用接口
``` 
- InitiateMultipartUploadHandler    初始化分块信息
- UploadPartHandler                 上传分块
- CompleteUploadPartHandler         通知分块上传完成
- CancelUploadPartHandler           取消上传分块
- MultipartUploadStatusHandler      查看分块上传的整体状态
```

### 接口:上传初始化
- 判断是否已经上传过
- 生成唯一上传ID
- 缓存分块初始化信息

### redis操作
``` 
redis-cli
auth testupload

keys *
quit

```
### redis连接池
- newRedisPool

## 6-3 编码实战: 实现初始化分块上传接口
### InitialMultipartUploadHandler
- 1.解析用户请求参数
- 2.获得redis的一个连接
- 3.生成分块上传的初始化信息
- 4.将初始化信息写入到redis缓存
- 5.将响应初始化数据返回到客户端


## 6-4 编码实战: 实现分块上传接口
### UploadPartHandler
- 1.解析用户请求参数
- 2.获得redis的一个连接
- 3.获得文件句柄，用于存储分块内容
- 4.更新redis缓存状态
- 5.返回处理结果到客户端

## 6-5 编码实战: 实现分块合并接口
### CompleteUploadHandler
- 1.解析用户请求参数
- 2.获得redis的一个连接
- 3.通过uploadid查询redis并判断是否所有分块上传完成
- 4.TODO: 合并分块
- 5.更新唯一文件表及用户文件表
- 6.响应处理结果

## 6-6 分块上传场景测试+小结
- test_mpupload
- go run main.go
- go run test_mpupload.go
- 手动合并验证
```
cat `ls | sort -n` > /tmp/a
shalsum /tmp/a
```
### 上传取消
- 删除已存在的分块文件
- 删除redis缓存状态
- 更新mysql文件status

### 上传状态查询
- 检查分块上传状态是否有效
- 获取分块初始化信息
- 获取已上传的分块信息

### 本章小结
- 1.分块上传与断点续传的概念
- 2.分块上传流程的讲解
- 3.几个重要接口的逻辑实现与演示

## 6-7 文件断点下载原理