# Docker 镜像制作

## docusaurus 制作

1. 创建 docusaurus 源文件
> npx create-docusaurus@latest docusaurus classic

2. 创建 Dockerfile 文件

> touch Dockerfile

3. 构建
> docker build -t docs:v1 .

4. 使用指定**端口**启动容器
> docker run -t -d -p 80:3000 --name d1 docs:v1

5. 使用 root 登陆 docker 容器
   
> docker exec -it -u root cf6e71b62c24 bash

6. 检查容器网络

获取容器网络 ID
>  docker container inspect containerId 

查看容器网络
>  docker network inspect containerId 

查看容器日志
> ocker logs -f containerId

绑定本地磁盘到docker容器

```bash
docker run -t -d -p 3004:3000 \
--name docsv3 \
--mount type=bind,source=/Users/ziv/github.com/favorite/workspace/docusaurus,target=/home/workspace/docusaurus \
> docs:v2
```

```bash
docker run -t -d -p 3004:3000 \
> --name mySite \
> --mount type=bind,source=/Users/ziv/github.com/favorite/workspace/docusaurus,target=/home/workspace/docusaurus \
> -w /home/workspace/docusaurus \
> node:16 \
> sh -c "yarn install && yarn start"
```

## Q&A

1. docker容器启动不了的问题
****
https://www.jianshu.com/p/f1e7a1630c64

1. docker 数据存储和共享

https://docs.docker.com/storage/volumes/

3. Docker 磁盘挂载 
https://docs.docker.com/get-started/06_bind_mounts/