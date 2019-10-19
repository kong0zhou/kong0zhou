docker build -t ccr.ccs.tencentyun.com/kongzhou/logshow:v1 .
#删除所有空白镜像
docker rmi $(docker images -f "dangling=true" -q)