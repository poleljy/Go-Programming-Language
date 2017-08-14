## 简单的HTTP POST
通过HTTP向服务器发送POST请求提交数据，都是通过form表单提交的，代码如下：
```
<form method="post" action="http://w.sohu.com" >
    <input type="text" name="txt1">
    <input type="text" name="txt2">
</form>
```
提交时会向服务器端发出这样的数据（已经去除部分不相关的头信息），数据如下：
```
POST / HTTP/1.1

Content-Type:application/x-www-form-urlencoded
Accept-Encoding: gzip, deflate
Host: w.sohu.com
Content-Length: 21
Connection: Keep-Alive
Cache-Control: no-cache

txt1=hello&txt2=world
```
对于普通的HTML Form POST请求，它会在头信息里使用Content-Length注明内容长度。头信息每行一条，空行之后便是Body，即“内容”（entity）。
Content-Type是application/x-www-form-urlencoded，这意味着消息内容会经过URL编码，就像在GET请 求时URL里的QueryString那样

## POST上传文件
最早的HTTP POST是不支持文件上传的,`Content-Type`的类型扩充了`multipart/form-data`用以支持向服务器发送二进制数据,

要使表单能够上传文件，首先第一步就是要添加form的enctype属性，enctype属性有如下三种情况:
```
application/x-www-form-urlencoded   表示在发送前编码所有字符（默认）
multipart/form-data   				不对字符编码。在使用包含文件上传控件的表单时，必须使用该值。
text/plain    						空格转换为 "+" 加号，但不对特殊字符编码。
```

## Multipart/form-data
http协议规定了以ASCII码传输，建立在tcp、ip协议之上的应用层规范，规范内容把http请求分为3个部分：状态行，请求头，请求体。

1、`multipart/form-data` 的基础方法是post，也就是说是由post方法来组合实现的

2、`multipart/form-data` 与post方法的不同之处：请求头，请求体。

3、`multipart/form-data` 的请求头必须包含一个特殊的头信息： `Content-Type`，且其值也必须规定为 `multipart/form-data` ，同时还需要规定一个内容分割符用于分割请求体中的多个post的内容，如文件内容和文本内容自然需要分割开来，不然接收方就无法正常解析和还原这个文件了。具体的头信息如下：
```
Content-Type: multipart/form-data; boundary=${bound}  
```
其中${bound} 是一个占位符，代表我们规定的分割符，可以自己任意规定，但为了避免和正常文本重复了，尽量要使用复杂一点的内容。如：--------------------56423498738365

4、`multipart/form-data` 的请求体也是一个字符串，不过和post的请求体不同的是它的构造方式，post是简单的name=value值连接，而multipart/form-data则是添加了分隔符等内容的构造体。具体格式如下:
```
--${bound}
Content-Disposition: form-data; name="Filename"

HTTP.pdf
--${bound}
Content-Disposition: form-data; name="file000"; filename="HTTP协议详解.pdf"
Content-Type: application/octet-stream

%PDF-1.5
file content
%%EOF

--${bound}
Content-Disposition: form-data; name="Upload"

Submit Query
--${bound}--
```
其中${bound}为之前头信息中的分割符，如果头信息中规定为123，那么这里也要为123,；可以很容易看出，这个请求体是多个相同的部分组成的：每一个部分都是以--加分隔符开始的，然后是该部分内容的描述信息，然后一个回车，然后是描述信息的具体内容；如果传送的内容是一个文件的话，那么还会包含文件名信息，以及文件内容的类型。上面的第二个小部分其实是一个文件体的结构，最后会以--分割符--结尾，表示请求体结束。 


## 实例分析：
```
<form method="post"action="http://w.sohu.com/t2/upload.do" enctype=”multipart/form-data”>
    <input type="text" name="desc">
    <input type="file" name="pic">
</form> 
```

浏览器将会发送以下数据：
```
POST /t2/upload.do HTTP/1.1
User-Agent: SOHUWapRebot
Accept-Language: zh-cn,zh;q=0.5
Accept-Charset: GBK,utf-8;q=0.7,*;q=0.7
Connection: keep-alive
Content-Length: 60408
Content-Type:multipart/form-data; boundary=ZnGpDtePMx0KrHh_G0X99Yef9r8JZsRJSXC
Host: w.sohu.com

--ZnGpDtePMx0KrHh_G0X99Yef9r8JZsRJSXC
Content-Disposition: form-data;name="desc"
Content-Type: text/plain; charset=UTF-8
Content-Transfer-Encoding: 8bit

[......][......][......][......]...........................
--ZnGpDtePMx0KrHh_G0X99Yef9r8JZsRJSXC
Content-Disposition: form-data;name="pic"; filename="photo.jpg"
Content-Type: application/octet-stream
Content-Transfer-Encoding: binary

[图片二进制数据]
--ZnGpDtePMx0KrHh_G0X99Yef9r8JZsRJSXC--
```