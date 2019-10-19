# logShow

## 简介
logShow是一个实时查看日志文件的软件
### 特点
基于sse,比websocket更轻量

部署基于docker,更简单便捷

支持多个文件查看

### 用法

以ubuntu为例：

使用前需要现在电脑中安装docker

```sh
docker pull ccr.ccs.tencentyun.com/kongzhou/logshow:v1
```

```sh
docker run \
	-t \
	--name logshow \
	-p 8083:8083 \
	--env UID=yourusername \
	--env PASSWORD=yourpassword \
	--env SESSIONKRY=yoursessionkey \
	-v $(pwd)/yourlogdir:/home/logfiles \
	ccr.ccs.tencentyun.com/kongzhou/logshow:v1
```

最后只需要用浏览器打开 localhost:8083 即可