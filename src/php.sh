PHP=`pwd`
docker run --name php \
	--rm \
	-it \
	-v $PHP/code:/code \
	php \
	php /code/index.php
