#!/usr/bin/env python3
import gevent.monkey
import gevent
import urllib
import urllib.request
import sys
from gevent import socket
import json
import time
gevent.monkey.patch_socket()

ali_1 = {
    "shopId":"3344",
    "cashier":"bo",
    "type":1,
    "posCode":"1001",
    "timeStamp":"2015-9-4 00:01:02"
}

ali_2 = {
    "shopId":"3344",
    "cashier":"薄建彬",
    "type":2,
    "posCode":"1001",
    "timeStamp":"2015-9-4 00:01:02",
    "content":{
        "index":4,
        "item_id":"6",
        "title":"大头鱼",
        "sku_code":"00010101",
        "unit_price":"19.80",
        "quantity":1,
        "member_level":"VIP金卡",
        "member_no":"123456",
        "member_title":"薄建彬",
        "total_quantity":1
    }
}
ali_2_1 = {
    "shopId":"3344",
    "cashier":"薄建彬",
    "type":2,
    "posCode":"1001",
    "timeStamp":"2015-9-4 00:01:02",
    "content":{
        "index":4,
        "item_id":"6",
        "title":"大头鱼",
        "sku_code":"00010101",
        "unit_price":"19.80",
        "quantity":0.354676,
        "member_level":"VIP金卡",
        "member_no":"123456",
        "member_title":"薄建彬",
        "total_quantity":1
    }
}
ali_3 = {
    "shopId":"3344",
    "cashier":"薄建彬",
    "type":3,
    "posCode":"1001",
    "timeStamp":"2015-9-4 00:01:02",
    "content":{
        "amount": "10.08",
		"pay_amount": ["20", "0.08"],
		"time": "2019-12-02 10:09:36",
		"discount": "0.00",
		"change": "10",
		"list_no": "742071713505115462",
		"pay_type": ["现金结算", "抹零"]
    }
}
ali_3_1 = {
    "shopId":"3344",
    "cashier":"薄建彬",
    "type":3,
    "posCode":"1001",
    "timeStamp":"2015-9-4 00:01:02",
    "content":{
        "amount": "100",
		"pay_amount": ["100"],
		"time": "2019-12-02 10:14:31",
		"discount": "20",
		"change": "20",
		"list_no": "742084577120628267",
		"pay_type": ["支付宝结算"]
    }
}
ali_4 = {
    "shopId":"3344",
    "cashier":"薄建彬",
    "type":4,
    "posCode":"1001",
    "timeStamp":"2015-9-4 00:01:02",
    "content":{
        "event_id":"3",
        "action":"退货",
        "authorizer":"12345"
    }
}
ali_4_1 = {
    "shopId":"3344",
    "cashier":"薄建彬",
    "type":4,
    "posCode":"1001",
    "timeStamp":"2015-9-4 00:01:02",
    "content":{
        "event_id":"3",
        "action":"打开钱箱",
        "authorizer":"12345"
    }
}
ali_4_2 = {
    "shopId":"3344",
    "cashier":"薄建彬",
    "type":4,
    "posCode":"1001",
    "timeStamp":"2015-9-4 00:01:02",
    "content":{
        "event_id":"3",
        "action":"清空购物车",
        "authorizer":"12345"
    }
}
ali_4_3 = {
    "shopId":"3344",
    "cashier":"薄建彬",
    "type":4,
    "posCode":"1001",
    "timeStamp":"2015-9-4 00:01:02",
    "content":{
        "event_id":"3",
        "action":"取消交易",
        "authorizer":"12345"
    }
}
ali_4_4 = {
    "shopId":"3344",
    "cashier":"薄建彬",
    "type":4,
    "posCode":"1001",
    "timeStamp":"2015-9-4 00:01:02",
    "content":{
        "event_id":"3",
        "action":"终止交易",
        "authorizer":"12345"
    }
}
ali_4_5 = {
    "shopId":"3344",
    "cashier":"薄建彬",
    "type":4,
    "posCode":"1001",
    "timeStamp":"2015-9-4 00:01:02",
    "content":{
        "event_id":"3",
        "action":"删减商品",
        "authorizer":"12345"
    }
}
ali_4_6 = {
    "shopId":"3344",
    "cashier":"薄建彬",
    "type":4,
    "posCode":"1001",
    "timeStamp":"2015-9-4 00:01:02",
    "content":{
        "event_id":"3",
        "action":"删除商品",
        "authorizer":"12345"
    }
}

ali_5 = {
    "shopId":"3344",
    "cashier":"薄建彬",
    "type":5,
    "posCode":"1001",
    "timeStamp":"2015-9-4 00:01:02",
    "content":{
        "order_id":"101010",
        "memeber_no":"232323",
        "member_level":"VIP"
    }
}
ali_6 = {
    "shopId":"3344",
    "cashier":"薄建彬",
    "type":6,
    "posCode":"1001",
    "timeStamp":"2015-9-4 00:01:02",
    "content":{
        "duty_code":"101010",
        "order_id":"232323"
    }
}


ali_7 = {
    "shopId":"3344",
    "cashier":"薄建彬",
    "type":7,
    "posCode":"1001",
    "timeStamp":"2015-9-4 00:01:02",
    "content":{
        "duty_code":"101010",
        "order_id":"232323"
    }
}
ali_8 = {
    "shopId":"3344",
    "cashier":"薄建彬",
    "type":8,
    "posCode":"1001",
    "timeStamp":"2015-9-4 00:01:02",
    "content":{
        "sku_code":"101010",
        "unit_price":"23",
        "title":"大头鱼",
        "quantity":1
    }
}

raw = ("POST /v1/pos/operate/backflow HTTP/1.1\r\n"
"Accept-Encoding: identity\r\n"
"Content-Type: application/x-www-form-urlencoded\r\n"
"Content-Length: 99\r\n"
"Host: 192.168.7.253:8024\r\n"
"User-Agent: Python-urllib/3.7\r\n"
"Connection: close\r\n\r\n"
"{\"shopId\": \"3344\", \"cashier\": \"bo\", \"type\": 1, \"posCode\": \"1001\", \"timeStamp\": \"2015-9-4 00:01:02\"}")
raw2 = ("POST /v1/pos/operate/backflow HTTP/1.1\r\n"
"Accept-Encoding: identity\r\n"
"Content-Type: application/x-www-form-urlencoded\r\n"
"Content-Length: 99\r\n"
"Host: 192.168.7.253:8024\r\n"
"User-Agent: Python-urllib/3.7\r\n"
"Connection: close\r\n\r\n"
"{\"shopId\": \"3344\", \"cashier\": \"bo\", \"type\": 1, \"posCode\": \"1001\", \"timeStamp\": \"2015-9-4 00:01:02\"}")
def tcp_long():
    data = []
    while True:
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        s.connect(('192.168.7.253', 8024))
        #print(raw)
        #s.send(raw.encode("utf-8"))
        #data.append(s)
        time.sleep(200)
        s.close()
        

def tcp_short():
    Ddata = [ali_1,ali_2,ali_2_1,ali_3,ali_3_1,ali_4,ali_4_1,ali_4_2,ali_4_3,ali_4_4,ali_4_5,ali_4_6,ali_5,ali_6,ali_7,ali_8]
    i = 0
    while True:
        req = urllib.request.Request("http://127.0.0.1:8023/v1/pos/operate/backflow")
        _data = json.dumps(Ddata[i])
        print(i,_data)
        with urllib.request.urlopen(req,data=_data.encode("utf-8")) as f:
            print("Status:",f.status)
            print("Data:",f.read().decode("utf-8"))
            i = i + 1
            if i >= len(Ddata):
                i = 0
            time.sleep(5) 

if __name__ == '__main__':
    tcp_short()