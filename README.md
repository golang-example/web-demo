# go web框架 iris demo

## iris 显示html没有问题
直接返回html内容
访问 http://127.0.0.1:8080/webdemo/html/v1/index

## 使用md文件markdown显示接口文档
访问md文件 http://127.0.0.1:8080/webdemo/html/v1/passenger/login.md
页面显示的md内容,缺少了样式 不好看,不行 整合swagger试试

## 使用swagger写接口文件
https://ieevee.com/tech/2018/04/19/go-swag.html
按照文档 安装好swag命令 需要翻墙下载  配置好终端代理翻墙 如用蓝灯翻墙
接口文件编写格式 参看xinyun-user项目
### 注意 swagger只支持gin框架  它的包为gin-swagger 
想访问swagger必须启动 gin进程 监听一个端口
我在iris项目中,不能同时启动, 就加了个参数判断

在项目根目录里执行swag init，生成docs/docs.go
```shell
swag init
godep go build
只能同时启动一个, 所以启动不同端口
./xinyun-user -swagger=true 启动gin的swagger文档
./xinyun-user  启动iris项目, 默认不启动gin的swagger文档
```

main代码如下:
```shell
if *swagger {
		//启动访问swagger文档
		port := ":8081"
		Log.Info("listen on :%s", port)
		//gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
		r := gin.Default()
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		//监听端口
		r.Run(port)
	} else {
		Log.Info("listen on :%s", Cfg.ListenPort)
		//监听端口
		web.RunIris(Cfg.ListenPort)
	}
```

访问swagger
http://127.0.0.1:8081/swagger/index.html
访问接口
http://127.0.0.1:8080/xinyun-user/passenger/v1/login

## 端口不同 跨域问题
使用swagger调用接口发现有问题  就是跨域的问题
F12页面调试工具 Console显示跨域报错, 还发现所有的请求都是用的Options方式去请求的

### 解决方案一: 使用nginx代理  成功 正常
```shell
vim /usr/local/etc/nginx/nginx.conf 修改配置
upstream xinyun-user_upstream {
    server 127.0.0.1:8080;
}
 
upstream swag_upstream {
     server 127.0.0.1:8081;
}

location /xinyun-user {
  proxy_pass http://xinyun-user_upstream;
  proxy_set_header  Host                $host;
  proxy_set_header  X-Real-IP           $proxy_protocol_addr;
  proxy_set_header  X-Forwarded-For     $proxy_add_x_forwarded_for    ;
  proxy_set_header  X-Forwarded-Port    $server_port;
}

location /swagger {
  proxy_pass http://swag_upstream;
  proxy_set_header  Host                $host;
  proxy_set_header  X-Real-IP           $proxy_protocol_addr;
  proxy_set_header  X-Forwarded-For     $proxy_add_x_forwarded_for    ;
  proxy_set_header  X-Forwarded-Port    $server_port;
}
sudo nginx 启动
```

访问 
http://127.0.0.1/swagger/index.html
测试接口,发的是post请求,也能正常响应数据

### 解决方案二: 让所有接口支持Options方式  有问题
在web下的iris.go文件里方法 switch判断外加上 mw.Options(info.Uri, info.Handler...)
让所有接口都支持Options方式请求,如下:
```shell
func ProcessMapping(mw iris.MuxAPI, registerInfo *[]regInfo) {
	for _, info := range *registerInfo {
		switch info.Method {
		case "Get":
			mw.Get(info.Uri, info.Handler...)
		case "Post":
			mw.Post(info.Uri, info.Handler...)
		case "Put":
			mw.Put(info.Uri, info.Handler...)
		case "Delete":
			mw.Delete(info.Uri, info.Handler...)
		case "Options":
			mw.Options(info.Uri, info.Handler...)
		}
		fmt.Println("--------")
		mw.Options(info.Uri, info.Handler...)
	}
}
```
还需要设置允许跨域访问
在web下的AccessLogMiddleware.go里加入
//允许跨域访问
ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")

访问swagger
http://127.0.0.1:8081/swagger/index.html
测试接口,还是发的Options请求, 参数传不了, 由于服务端支持Options请求, 会有响应数据 参数错误

使用Postman工具 发送Options请求,带参数, 会有正确的响应数据 是成功的
但swagger页面上不知道怎么传参数

如果接口不支持options方式, 访问时直接报404, iris处理了,根本没进项目里.所以根本还是要iris允许跨域

我猜: 这里只是让所有的接口都支持options方式的请求,没有根本上允许跨域访问, 
按道理是要在初始化iris的时候设置允许跨域,估计要在iris框架源码里加.
java里,像spring boot项目里,初始化配置config时,允许跨域
```java
@Configuration
public class WebMvcConfig implements WebMvcConfigurer {
    @Override
    public void addCorsMappings(CorsRegistry registry) {
        //允许跨域
        registry.addMapping("/**")
                .allowCredentials(true)
                .allowedHeaders("*")
                .allowedOrigins("*")
                .allowedMethods("*");
    }
}
```

#### 总结: nginx是最好的解决方案, 或者使用gin框架,现在也流行gin框架
