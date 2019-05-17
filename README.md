# goappk8s

# 资源
来源：
https://www.qikqiak.com/post/write-kubernets-golang-service-step-by-step/
项目：
https://github.com/mdemo19/goappk8s



# 运行测试：
```
cd goappk8s
// $ go install github.com/mdemo19/goappk8s && ./bin/goappk8s
go build && ./goappk8s

```

# 更新流程：
## 1.改go代码
## 2.本地测试 $ go run *.go
## 3.生成build文件；
## 4.生成docker image 
## 5.上传 docker image;
## 6.k8s yaml 修改
## 7.k8s update;
## 8.测试；


# 更新脚本
```
sudo chmod  777 ./build.sh
./build.sh
```

# 目录结构
rs : 程序运行时依赖等资源文明总目录，存放配置，图片等
build.sh 打包sh;
Dockerfile  docker
doc : 其它资源文件；


