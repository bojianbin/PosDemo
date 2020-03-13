# PosDemo

server:简单的http json格式验证服务。
client:简单的http json发送脚本。

server端编译后的可执行文件在【release】处获取,PosDemo是Linux amd64版本，PosDemo.exe是windows amd64版本。

服务端运行方法：(注：监听8023端口，url是/v1/pos/operate/backflow，短连接)
        ./PosDemo

client处的python脚本是很久之前做dg200的POS功能时用的一个自测工具，此处的client和server是配套自洽的，可提供简单的查阅与验证。
