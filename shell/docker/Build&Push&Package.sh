#!/bin/bash

#build
ProjectName=web-demo

if [ $1 = test ]
then
    RegistryAddr=test.docker.registry:5000
elif [ $1 = pro ]
then
    RegistryAddr=pro.docker.registry:5000
    docker login $RegistryAddr -u username -p password
else
    echo "must input the project enviroment, limit in \"test\" or \"pro\""
    exit
fi

SOURCE_PATH=$GOPATH"/src/${ProjectName}"
INSTALL_PACK_PATH=$(pwd)

cd $SOURCE_PATH

CGO_ENABLED=0 $GOPATH/bin/godep go build -o ./shell/docker/$ProjectName

if [ -f ./shell/docker/$ProjectName ]; then
    echo "Build Success."
else
    echo "Build Failed."
    exit
fi

cp -r templates ./shell/docker/

cd $INSTALL_PACK_PATH

branch=`git symbolic-ref --short -q HEAD`
tag=`git rev-parse --short HEAD`

Image_Name=${RegistryAddr}/$(echo $ProjectName | tr '[A-Z]' '[a-z]')/${branch}:v${tag}

#create build-deploy file
#macbook
sed -i '' -e '2a \
ProjectName='${ProjectName} docker-entrypoint.sh

#linux
#sed -i "2a ProjectName=${ProjectName}" docker-entrypoint.sh

#build docker image
docker build -t $Image_Name .

#push
#docker push $Image_Name

#delete
#macbook
sed -i '' '3d' docker-entrypoint.sh
#linux
#sed -i '3d' docker-entrypoint.sh

rm -rf ${ProjectName}

rm -rf templates
