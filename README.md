# 验证码识别部署文档

我们的验证码识别分为两个项目

1. [tensorflow 验证码识别服务] (https://gitee.com/_admin/yzm)
2. [验证码识别网站服务](https://github.com/zm-dev/yzm-client)

服务间使用 [grpc](https://github.com/grpc/grpc) 通信。同时两个项目已经 `push` 到阿里云 docker 仓库，可以使用 docker 轻易部署。

## 验证码识别网站部署方法
1. 下载 [docker-compose.yml](https://raw.githubusercontent.com/zm-dev/yzm-client/master/docker-compose.yml) 文件。
2. 进入 docker-compose.yml 所在的目录
3. 执行 `docker-compose up -d`
完整命令
```
mkdir yzm
cd yzm
wget https://raw.githubusercontent.com/zm-dev/yzm-client/master/docker-compose.yml
docker-compose up -d
```
等待镜像下载和启动完毕,直接打开浏览器访问 http://localhost:8102 即可访问到验证码识别的网站。
<img src="https://github.com/zm-dev/yzm-client/blob/master/screenshots/1.png?1" />
该网站的操作方法麻烦您浏览我们提供的视频。

## 生成 mappings.txt 的方法:
1. **(不推荐)** 上传 `zip` 文件到网站进行识别
首先将待识别的验证码打包成 `.zip` 的文件，然后在我们提供的验证码识别网站上面直接上传`.zip` 文件，并点击识别。等待识别完成后点击下载 `mappings.txt` 即可
<img src="https://github.com/zm-dev/yzm-client/blob/master/screenshots/5.png" />

上图为 *上传 .zip 文件* 界面

<img src="https://github.com/zm-dev/yzm-client/blob/master/screenshots/6.png" />

上图为 *点击下载 mappings.txt 文件* 界面

2. **(推荐方法)** 上面的方法识别速度较慢，因为通过 [grpc](https://github.com/grpc/grpc) 与底层的[tensorflow 验证码识别服务](https://gitee.com/_admin/yzm)通信毕竟有延时。我们可以直接使用[tensorflow 验证码识别服务](https://gitee.com/_admin/yzm)中提供的命令行工具来生成 `mappings.txt`

### 首先准备以下文件(假设以下文件存放在 `/root/yzm` 下！！！)
<img src="https://github.com/zm-dev/yzm-client/blob/master/screenshots/tree_1_4.png">


获取帮助命令:
```
docker run registry.cn-hangzhou.aliyuncs.com/zm-dev/yzm:latest /usr/bin/env python /app/run.py -h
```

批量识别命令:
```
docker run -v /root/yzm/:/test/:rw registry.cn-hangzhou.aliyuncs.com/zm-dev/yzm:latest /usr/bin/env python /app/run.py -o /test/data-1/mappings.txt /test/data-1/
```
命令执行完毕后会得到以下输出:
```
您没有使用-c选项指定验证码分类，已经自动判断分类为：1
识别完成，正在排序并写入 mappings 文件中...
成功生成文件：/test/data-1/mappings.txt
```
打开 `/root/yzm/data-1/mappings.txt` 即可看到第一类验证码的识别结果，其他类别验证码的识别方法一样。


单个识别命令:
```
docker run -v /root/yzm/:/test/:rw registry.cn-hangzhou.aliyuncs.com/zm-dev/yzm:latest /usr/bin/env python /app/run.py -o /test/data-2/mappings.txt /test/data-2/0000.jpg
```
命令执行完毕后会得到以下输出:
```
您没有使用-c选项指定验证码分类，已经自动判断分类为：2
识别的结果为：0000,RNFYE
```

## 第 5 类验证码识别方法
### 首先准备以下文件(假设以下文件存放在 `/root/data-5` 下！！！)
<img src="https://github.com/zm-dev/yzm-client/blob/master/screenshots/tree_5.png">

执行命令:
```
docker run -v /root/data-5/:/test/:rw registry.cn-hangzhou.aliyuncs.com/zm-dev/yzm:latest /usr/bin/env python /app/run.py -c 5 -o /test/mappings.txt /test/
```
命令执行完毕后会得到以下输出:
```
识别完成，正在排序并写入 mappings 文件中...
成功生成文件：/test/mappings.txt
```
打开 `/root/data-5/mappings.txt` 即可看到第 5 类验证码的识别结果。


