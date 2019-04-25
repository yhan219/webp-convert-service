# webp转化服务

----------

基于`docker` 用 `golang` 实现的实时webp转化服务

## Docker

### 构建
 以下两种服务任选其一
#### 从 docker hub pull
> docker pull yhan219/webp-convert-service

#### 通过dockerfile构建
> docker build -t yhan219/webp-convert-service:1.0 .

### 运行
> docker run -d --name webp-service -p 80:80 yhan219/webp-convert-service:1.0


## 使用
默认监听80端口和`/`路径,其他路径返回`404`状态码

参数如下:

- **url**: 需要转换的图片地址,如果带参数,需要先urlEncode
- 其他参数: 添加 [转化参数](https://developers.google.cn/speed/webp/docs/cwebp),其中bool被解释为标志参数,例如`nostrong`

## 示例
需要转化的图片地址:
> https://www.baidu.com/img/bd_logo1.png

webp转化地址:
> http://localhost:80/?url=https://www.baidu.com/img/bd_logo1.png

带参数的转化:
> http://localhost:80/?url=https://www.baidu.com/img/bd_logo1.png&q=1&nostrong=true&z=5

以上链接会生成如下命令行:
> cwebp -q 1 -nostrong -z 5 -o - -- -
