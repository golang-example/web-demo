#!/usr/bin/python
# coding: utf-8

import urllib2
import json

url = "http://192.168.1.52:2379/v2/keys/config/web-demo/"

def read_json(file_name):
    with open(file_name) as json_file:
        data = json.load(json_file)
        return data

def sent_request(k, v):
    request = urllib2.Request(url+k, 'value='+bytes(v))
    request.get_method = lambda:'PUT'
    request = urllib2.urlopen(request)
    print request.read()

def do_action(file_name):
    json_dict = read_json(file_name)
    map(lambda k: sent_request(k, json_dict[k]), json_dict.keys())

do_action("json/" + 'prd.json')
