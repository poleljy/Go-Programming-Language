#### 写在前面
基于`HTTP`协议的`Web API`是时下最为流行的一种分布式服务提供方式。之前在审查写的 `Restful API` 的时候有些东西没理解清楚，后来查阅一些资料的时候看到 `Thomas Hunter II `在17年6月份发表介绍`HTTP API`设计的四篇文章（`Requests`,`Responses`,`Bodies`,`API Standards`）和该作者写的一本小册子`Consumer Centric API Design`，内容比较清晰，所以打算翻译下保留下来。第一次翻译这么长的文章，文字功底也比较拙劣...orz

点击阅读原文查看原英文，可能需要翻墙...orz

#### HTTP应答简介
服务端收到一个客户端请求之后总会返回一个回复(排除一些特别严重情况的发生)。应答消息格式和请求消息格式类似，不过还是有不少的区别。下面是一个完整的`HTTP`消息回复：
```
HTTP/1.1 200 OK                       <-- Status Line
Date: Wed, 14 Jun 2017 23:23:01 GMT   <-- Response Headers
Content-Type: application/json
Access-Control-Max-Age: 1728000
Cache-Control: no-cache
                                      <-- Two Newlines
{                                     <-- Body
  "id": "12",
  "created": "2017-06-14T23:22:59Z",
  "modified": null,
  "name": "Gir",
  "animal_type": "12"
}
```
第一行是状态行，包含两个信息。第一个是协议信息，这里是`HTTP 1.1`，接下来是个空格。第二个信息是请求状态，由一个机器可读的状态编码，一个空格和一个人可读的解析词，接下来是新的一行。解析词和状态编码是一一对应的，不过在大多数情况下都可以忽略这个过时字符串。
和`HTTP`请求一样，应答消息头里面也包含一系列的`key: value`键值对，下面是新的一行。理论上来说头部信息可以是任何形式，不过严谨的服务端会返回首字母大写，使用连字符的词组。头部字段可以重复出现如果有多个值，比如单个应答里面有多个`Set_Cookie`头部字段允许设置多个`cookies`。
状态码用不同区间的数字区分，不同区间表示不同类型的状态。文中接下来描述了不同的数字区间和区间里面的数值含义。
头部最下面是可选的消息体，消息体前面需要有两个空行。几乎所有的应答都会包含一个消息体，即使理论上是可选的。

#### 1XX Status Codes – Informational
`1XX`区间(100-199)里面的数字表示报告状态码。我们整个职业生涯永远不需要处理这个区间里面的状态码。
* `101 Switching Protocols`: used for websockets

#### 2XX Status Codes – Successful
`2XX`区间(200-299)里面的数字表示成功状态码。反馈给客户端的信息就是操作成功了，无论是创建一个资源或者是简单的获取一个资源。理论上来说它们是最常被使用的状态码。
* 200 OK: 可以代表所有操作都成功。
* 201 Created: 资源被创建。
* 202 Accepted: 创建或更新资源成功，不过操作是异步的。
* 204 No Content: 请求成功，不过没有请求体。

#### 3XX Status Codes – Redirection
`3XX`区间(300-399)里面的数字表示客户端需要重定向。不包含消息体，会有一个`Location`的消息头，包含了一个客户端需要重定向的`URL`。
* 301 Moved Permanently: 资源被移动到另外一个新地址，需要访问新地址。
* 302 Found: 资源在新地址，不过每次还是需要检查下旧地址。

#### 4XX Status Codes – Client Error
`4XX`区间(400-499)里面的数字表示客户端错误。当一个`Unsafe`的请求返回该区间的状态码，那么这个请求一般还没有改变服务端状态。
* 400 Invalid Request: 一般的客户端错误
* 401 Unauthorized: 客户端需要提供权限认证请求头
* 403 Forbidden: 客户端禁止访问该资源
* 404 Not Found: 资源不存在
* 405 Method Not Allowed: 请求路径存在，只是不支持该请求类型
* 406 Not Acceptable: 服务端无法对请求头里面的`Accept`类型产生应答

#### 5XX Status Codes – Server Error
`5XX`区间(500-599)里面的数字表示服务端错误。当返回这类错误很可能是服务端发生错误，也许请求根本没被接收。如果一个`Unsafe`请求返回该状态，就无法知道服务端的状态了。这类的错误应当不计成本的被避免出现。
* 500 Internal Server Error:一般的服务端错误
* 501 Not Implemented: 服务端不支持该类型的请求路径
* 503 Service Unavailable: 服务端临时无法访问，可能是数据库断开了连接
* 521 Web Server Is Down: 中间服务器无法连接远程服务

#### 应答头部
和请求类似，应答消息也有大量的头部字段可以使用。头部字段描述了应答的元数据信息。通常你可直接使用标准的头部字段，不过有时候你需要自建字段，以`X-`开头，例如`X-Request-ID`。下面列出了常见的字段：
* Cache-Control: 缓存设置，比如`no-cache`如果资源不应该被缓存
* Content-Language: 消息体语言，比如`en-US`
* Content-Length: 消息体长度，如果一开始就知道
* Content-Type: 消息体类型，比如`application/json`
* Date: 服务器时间
* Expires: 消息体过期时间
* Server: 很少用到，被用来标识服务器
