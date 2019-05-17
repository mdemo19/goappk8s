#!/bin/sh
# mode=$1
# platform=$2

imageVersion="0.0.1"
echo $imageVersion
sleep 1s

function taskBuildGo(){
  echo "[\033[36m build go linux: --start \033[0m]"
  CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -v -o ./build/app-server
  echo "[\033[36m build go linux: --end \033[0m]"
}

function taskBuildDockerImage(){
  echo "[\033[36m build docker image: --start \033[0m]" 
  docker image build -t liam1803/mb-gotest:$imageVersion .
  echo "[\033[36m build docker image: --end \033[0m]"
}

function taskBuildDockerPush(){
  echo "[\033[36m build docker push: --start \033[0m]"   
  docker push liam1803/mb-gotest:$imageVersion
  echo "[\033[36m build docker push: --end \033[0m]"
}  

#taskBuildDockerPush
taskBuildDockerImage
