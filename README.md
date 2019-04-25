# webp convert service

----------

[中文文档](README_CN.md)

Based on `docker` with `golang` that convert images to webp

## Docker

### build
 There are two ways to construct it, either way.
#### pull from docker hub
> docker pull yhan219/webp-convert-service:lastest

#### build with dockerfile
> docker build -t yhan219/webp-convert-service:1.0 .

### run
> docker run -d --name webp-service -p 80:80 yhan219/webp-conevert-service:1.0


## Usage
The service listens on port 80 for GET requests on the root path (/). Any other path returns a 404 not found status.

### Params
- **url**: The URL of the image to convert.
- another param: key-value request params that are passed on to the appropriate [cwebp](https://developers.google.cn/speed/webp/docs/cwebp) binary. Boolean values are interpreted as flag arguments (e.g.: -nostrong).

## Sample
image url is:
> https://www.baidu.com/img/bd_logo1.png

then convert url is:
> http://localhost:80/?url=https://www.baidu.com/img/bd_logo1.png

add convert param,eg:
> http://localhost:80/?url=https://www.baidu.com/img/bd_logo1.png&q=1&nostrong=true&z=5

will have the effect of the following command-line being executed on the server:
> cwebp -q 1 -nostrong -z 5 -o - -- -






