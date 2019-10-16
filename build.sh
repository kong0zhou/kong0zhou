docker build -t logshow/kz:v1 .
#删除所有空白镜像
docker rmi $(docker images -f "dangling=true" -q)