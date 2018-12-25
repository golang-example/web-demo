#!/usr/bin/python
# coding: utf-8

import json
import urllib2
import random

print "run user user validate..."

port = 8080
host = 'http://127.0.0.1:' + str(port)
uri = '/webdemo/user/v1/add'

def gen_param():
    return {
        "userName" : "abcd" + bytes(random.randint(0,100)),
        "pwd" : "123456"
    }

def request_and_parse(_values):
    print _values
    request = urllib2.Request(host + uri, json.dumps(_values))
    request.add_header('X-Client', '1.0.1;Andirod;4.1;zh-CN;Africa/Lagos;4G')
    request.add_header('X-Web-Demo-RequestId', '123456789')
    return json.loads(urllib2.urlopen(request).read())

def check_miss_param(_values):
    _jsonResponse = request_and_parse(_values)
    # 验证
    if _jsonResponse['code'] != 3:
        print "Should hint 'Param Error'."
        print _values
        print _jsonResponse
        exit(-1)

# 验证缺少参数
values = gen_param()
del values['userName']
check_miss_param(values)

values = gen_param()
del values['pwd']
check_miss_param(values)

# 发送完整的请求
values = gen_param()
jsonResponse = request_and_parse(values)
print jsonResponse

if jsonResponse['code'] != 0:
    print "user add validate error."
    exit(-1)
