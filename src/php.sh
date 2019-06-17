PHP=`pwd`
docker run --name php \
	--rm \
	-it \
	-v $PHP/code:/code \
	php \
	php /code/index.php
	时空节点反抗螺丝钉咖啡碱
