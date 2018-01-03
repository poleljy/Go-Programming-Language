#### 写在前面
基于`HTTP`协议的`Web API`是时下最为流行的一种分布式服务提供方式。之前在审查写的 `Restful API` 的时候有些东西没理解清楚，后来查阅一些资料的时候看到 `Thomas Hunter II `在17年6月份发表介绍`HTTP API`设计的四篇文章（`Requests`,`Responses`,`Bodies`,`API Standards`）和该作者写的一本小册子`Consumer Centric API Design`，内容比较清晰，所以打算翻译下保留下来。第一次翻译这么长的文章，文字功底也比较拙劣...orz

点击阅读原文查看原英文，可能需要翻墙...orz

#### HTTP请求简介
`HTTP` 协议是位于传输层 `TCP` 协议之上的应用层协议。客户端通过 `request/response` 的形式向服务端发起请求和接受应答。下面是一个完整的 `HTTP` 请求:
```
POST /v1/animal HTTP/1.1        <-- Request Line
Host: api.example.org           <-- Request Headers
Accept: application/json
Content-Type: application/json
Content-Length: 24
                                <-- Two Newlines
{                               <-- Body
  "name": "Gir",
  "animal_type": "12"
}
```
第一行称为请求行，包含三个部分：请求方法（`Methods`）,统一资源标识符(`Endpoints`)和 `HTTP` 版本。接下来文中会详细讲解 `HTTP Methods` 和 `Endpoints`
请求行下面是头部信息（`Headers`），以 `key: value`的形式展示，头部信息分行显示。头部 `key` 最好以首字母大写，使用连字符的形式表示。服务端在任何情况下都应该总是接收头部信息。发送相同重复数据的情况下头部信息也会重复出现。
一个 `Http`请求只包含请求行和头部信息也是合法的。在 `GET`，`DELETE`，`HEAD`和`OPTIONS`情况下的请求总是这样。同样可以提供一个请求体（`Body`），最后一行头部信息下空两行是请求体。请求体大部分使用在`POST`，`PUT`和`PATCH`请求时。理论上来说，`GET`也是可以提供一个请求体的，在 `Elasticsearch`（译者注：一个基于`RESTful` web接口的全文搜索引擎）被大量使用。

#### Endpoints(Paths)
设计`RESTful` HTTP API的时候一个比较重要的是要能够抽象描述服务的功能，通过这样方式，所有的资源实体都能够以 `CRUD`（`Create`, `Read`, `Update`, `Delete`）来操作。在请求路径中不应该出现具体操作的动词。
使用最多的方式是使用不同的集合来表示不同相关资源。比如，服务端包含很多不同的公司和雇员信息，应该有一个集合叫 `companies` 和一个 `employees` 集合。我们可以使用特定的标识从集合中获取单个公司或雇员的信息。
使用这种模式表示资源和子集资源，我们可以得到更多的组合形式。比如，一个雇员能被一家公司雇佣。下面的示例能表示资源和集合之间的关系：
```
/companies
/companies/{company_id}
/employees
/employees/{employee_id}
/companies/{company_id}/employees
/companies/{company_id}/employee/{employee_id}
```
如果需要访问公司集合可以使用 `/companies`。如果需要访问某个具体公司信息可以使用 `/companies/{company_id}`。如果需要访问某个公司下个的某个职员信息，可以使用 `/companies/{company_id}/employee/{employee_id}`。接下来的部分会讲解不同 `Method` 来操作资源。

#### GET Method
`GET` 方法类似于 `CRUD` 中的 `Read`。`GET` 请求被认为是**安全**的，也就是说不会改变服务端状态。`GET` 请求是**幂等**的，也就是说多次相同操作不会有副作用。`GET` 请求不应该有请求体（body）。
当使用 `GET` 作用于集合时表示获取集合里面的资源列表。我们可以获取所有资源或者使用约束条件获取部分资源。下面有个例子表示去获取雇员列表：
```
GET /companies/twitter/employees?maxStartDate=2016-06-23&perPage=10&offset=20
```
上面的请求表示“获取在Twitter工作至少一年以上的员工信息，每页显示10个，偏移量20”。 `Endpoint`（`URL`路径）符合某些特定的 `RESTful` 标准，即使查询部分（`?key=val&key2=val2`）更像 `ad-hoc`。如果没有指定筛选条件，服务端会返回所有符合的资源。有的时候服务端会返回大量数据记录，所以服务端一般会有一个默认的 `limit` 来限制返回的数据量。
有的时候会需要获取指定资源的特定信息。这种情况下，客户端需要指定白名单或者黑名单来确定要返回的信息。示例如下：
```
GET /employees/tlhunter?fields=name,age,twitter
```
这个请求表示“查询表示为tlhunter员工的姓名，年龄和Twitter号”。如果没有指定需要返回的白名单或黑名单，服务端会返回资源的所有信息。

#### POST Method
`POST` 方法类似于 `CRUD` 中的 `Create`。`POST` 请求被认为是**不安全**的，也就是说会改变服务端状态。`POST` 请求是**非幂等**的，也就是说多次相同操作会创建不同的记录。`POST` 请求应该有请求体（body）。
`POST` 请求服务端会在集合下新建一个资源信息。下面的例子表示服务端会添加一条新雇员信息：
```
POST /employees
```
请求体里面应该要包含新添加的职员信息。
如果 `POST` 请求指向的是集合的子集会被用于创建一种新的关系。比如，如果我们需要一家公司雇佣一名职员，可以使用下面请求：
```
POST /companies/twitter/employees
```
请求体里面应该要包含已经存在雇员的信息，一般是标识或者雇员的主键。这样的话这个职员就会被指定公司雇佣。
如果服务端是可扩展性的数据集合，比如数据库，服务根目录（集合前一级）执行一个 `POST` 请求会创建一个新的集合。 

#### DELETE Method
`DELETE` 方法很明显类似于 `CRUD` 中的 `Delete`。`DELETE` 请求被认为是**不安全**的，也就是说会改变服务端状态。然而`DELETE` 请求是**幂等**的，也就是说多次相同操作不会有副作用（删除一次指定雇员，再删除也是同样效果）。`DELETE` 请求不应该有请求体（body）。
下面的例子通过移除和公司的关系，可以从一个公司解雇或终止一个职员：
```
DELETE /companies/twitter/employees/jkup
```
当然如果我们想从系统中完全移除职员信息，可以如下：
```
DELETE /employees/jkup
同样，如果服务端是个数据库，`DELETE` 一个集合会整个清除集合记录。
```

#### PUT/PATCH Methods
`PUT/PATCH` 类似于 `CRUD` 中的 `Update`。两个都被认为是**不安全**的。它们一般被认为是**幂等**的，因为执行多次有相同的结果。在只是单纯替换给定资源属性的情况下是正确的。然而，如果服务端可以改变某一属性，比如不断增加的属性，它们就是**非幂等**的。`PUT` 和 `PATCH` 都应该有一个请求体。
`PUT` 和 `PATCH` 的区别是什么？然而，这取决于你使用哪种 `API`。一种解释是 `PUT` 请求是完整更新，缺失的属性会被赋值为空。E.g.雇员信息里面有姓名和职位属性，执行只更新姓名的 `PUT` 请求会置空职位信息。 `PATCH` 请求部分更新记录。只有出现在请求体里面的字段被更新，没有出现的字段保持原来的值不变。一些 `API` 只提供 `PUT` 或 `PATCH` 来执行更新。
如果只想更新职员的姓名，可以执行以下的请求，请求体包含姓名信息：
```
PATCH /employees/tlhunter
```
如果想提升某个公司的某个职工职位，可以执行下面的请求，请求体里面包含新职位信息：
```
PATCH /companies/twitter/employees/kng
```

#### HEAD/OPTIONS Methods
`HEAD` 和 `OPTIONS` 请求有点不同。它们和 `CRUD` 的 `Read` 很类似，不过不返回资源信息。它们是**安全**和**幂等**的，因为它们不改变状态和不会有副作用。它们都不应该有请求体。

`HEAD` 请求用于返回资源的头部信息。这些头部信息包含一些有用的元数据信息，例如资源的失效期。如果获取资源消耗很大（比如比较大的文档传输很慢或者是一次昂贵的数据库操作），需要使用 `HEAD` 请求来代替 `GET` 请求。

`OPTIONS` 请求在区分浏览器访问的服务主机地址和浏览器访问的web应用主机时会很有用。现代浏览器实现了一个特性`CORS (Cross Origin Resource Sharing)`。web应用使用浏览器去访问远程主机上的文件时，在发起正式请求前会发起一次 `OPTIONS` 请求。服务端会在头部返回该资源的访问权限信息，比如跨域是否能访问该资源。如果这个权限检查失败了，那么浏览器就不会发起正常的请求了。

#### Request Headers
有许多不同种类的请求头部。特别是浏览器会发送大量的头部信息以及一直有新的头部字段被标准化。下面是一些常见的请求头和它们的含义，每一个都应该被 `HTTP API` 所支持：
* Accept: 客户端接收内容格式列表。比较常见的`HTTP API`内容格式是 `application/json`。
* Accept-Language: 客户端希望返回内容的编码格式。比如：`en-US`, `en`, `de-DE`。
* Content-Length: 请求体的字节长度。如果请求体的大小已知，该值会被设置。
* Content-Type: 请求体内容编码格式。只有在有请求体时才会发送。
* Host: HTTP请求主机地址，应该总是被发送出去。应用会忽略，在虚拟主机和请求路由时会很有用。
* User-Agent: 含有客户端标识信息的头部字段。微服务下识别客户端和哪个服务通信很有用。
这些是标准的头部，然而同样还有很多`ad-hoc`。一般经验来看，在创建一个自己的头部字段前应该去看一下是否已经有头部字段能满足你的需求。如果没有一般会添加一个前缀 `X`，比如：`X-API-Version` 用来表示API版本。