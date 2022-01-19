> *机密 CONFIDENTIAL, COPYRIGHT © AFMOBI GROUP* 

## 查询用户是否存在

-----


**1. 协议**


| Protocol |  Method  | Request Content-Type             | Response Content-Type               |
| -------- | -------- | -------------------------------- | ----------------------------------- |
| HTTP     | POST     | application/json; charset=utf-8  | application/json; charset=utf-8     |

-----

**2. Request-URL**


*URL:*
> https://host:port/api/login/v1/isexist


*Request Parameters：*
> 通用参数:

   [URL 请求中的数字签名参数](http://git.helloaf.com:3000/TUDC/server-doc/src/master/common/%E6%95%B0%E5%AD%97%E7%AD%BE%E5%90%8D%E6%8C%87%E5%BC%95.md)

*Header:*
>  [HTTP 通用 Request Header](http://git.helloaf.com:3000/TUDC/server-doc/src/master/common/HTTP%E9%80%9A%E7%94%A8Header.md)



-----

**3. Request-Data**


示例数据：
```
{
    "accountType":3,
    "cc":"86",
    "account":"18512345678",
    "imei":"xxxxx",
    "machineModel":"Tecno C8",
    "channel":"xxxxxx"
}
```

字段描述：

| 字段名           |  类型(最大长度)   | 是否必填 | 说明                                                              |
| --------------- | --------------- | ------- | ---------------------------------------------------------------  |
| accountType     | int             | 是      | 账户类型（enum: 1：username,2:email,3: phone）                     |
| cc              | string(50)      | 否      | 国家码，国家码不以加号和零开头 当accounttype=3时必填                    |
| account         | string(50)      | 是      | 用户账号                                                           |
| imei            | String(50)     | 是      | imei                                                               |
| machineModel    | string(50)      | 是      | 手机型号(eg:"Infinix X553","Tecno C5","Huawei P9")                 |
| channel         | String(200)     | 是      | 渠道                                                              |




-----


**4. Response-Data Success**

*Header:*  
   [HTTP 通用 Response Header](http://git.helloaf.com:3000/TUDC/server-doc/src/master/common/HTTP%E9%80%9A%E7%94%A8Header.md)

-----



示例数据：
```
{
  "code": 0,
  "clientKey": "afdafafaf",
  "regAppid": "xxxxxxx",
  "regZone": "xxxxx"
}
```
字段描述：

| 字段名             |  类型(最大长度)   | 是否必填 | 说明                                            |
| ----------------  | ---------------   | ------- | -------------------------------------------- |
| code              | int               | 是      | 返回码，0 用户存在                              |
| clientKey         | string    | 是        | pwd 加密盐值                                   |
| regAppid          | string    | 是        | 该用户注册来源appid                             |
| regZone           | string    | 是        | 该用户注册地区                                  |



-----

**5. Response-Data Error**

示例数据：
```
{
    "code": 110002,
    "msg" : "xxxxxx"
}
```
字段描述：

| 字段名           |  类型(最大长度)   | 是否必填 | 说明                                            |
| --------------- | ---------------  | ------- | ----------------------------------------------   |
| code            | int(5)           | 是      | 错误码:<br/>110002: 账号不存在<br/>[通用error code 定义](http://git.helloaf.com:3000/TUDC/server-doc/src/master/common/common_error_code%20%E5%AE%9A%E4%B9%89.md)                |
| msg             | string(Any)      | 是      | 错误信息                                          |

