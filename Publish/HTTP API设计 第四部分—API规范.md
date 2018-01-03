#### 写在前面
基于`HTTP`协议的`Web API`是时下最为流行的一种分布式服务提供方式。之前在审查写的 `Restful API` 的时候有些东西没理解清楚，后来查阅一些资料的时候看到 `Thomas Hunter II `在17年6月份发表介绍`HTTP API`设计的四篇文章（`Requests`,`Responses`,`Bodies`,`API Standards`）和该作者写的一本小册子`Consumer Centric API Design`，内容比较清晰，所以打算翻译下保留下来。第一次翻译这么长的文章，文字功底也比较拙劣...orz

点击阅读原文查看原英文，可能需要翻墙...orz

#### Simple Envelope
第一种方式基本上是主观想象的。返回给客户端的是简单的应答信息或者请求的资源。然而大部分情况下需要提供元数据信息。应答消息的头部能提供部分的元数据信息，但是还是不能完全包含客户端需要指定的所有信息。
让我们考虑一些错误的情况。当然我们可以提供`4XX`或`5XX`这样的返回值来提示请求失败，但是如何获取更多的信息。不论采取何种方式，我们都应该保证客户端总是能知道一个请求是否失败和从哪里去获取错误信息。
如果我们只是简单的返回资源或资源集合，我们会做出一些不够明智的决定，例如给资源添加一个错误信息属性。相反，我们应该返回一个实际数据的父类对象。我们定义这种标准化的`json`对象叫`Envelope`，封装了重要的数据。
下面是个失败请求的简单封装：
``` json
{
  "error": "database_connection_failed",
  "error_human": "Unable to establish database connection",
  "data": null
}
```
可以看到有两个错误字段。第一个是`error`，是机器可读的错误编码。很多服务选择使用一个数值来表示，但是为什么选择一个不可读的数字而不用字符串来表示？我们同样需要一个人类可读的字符串。理论上这个可读字符串应该翻译到指定`Accept-Language`的语言，然后展示给最终用户。代码中去比较人类可读的错误字符串会令人疯狂，特别是它们会发生变化，所以需要使用两个不同的属性字段。
如果是成功的返回，可以添加如下的附加属性：
``` json
{
  "error": null,
  "error_human": null,
  "data": [{"id": "11"}, {"id": "12"}],
  "offset": 10,
  "per_page": 10
}
```
这种情况下我们仍然需要一个`error`属性字段，没有错误就赋值为`null`。`data`字段包含了客户端请求的内容，上面这种情况下客户端请求的是个资源集合。最后，我们还有两个附加元数据属性，`offset`和`per_page`。在这里，客户端请求第二页的数据，每页10条记录。

#### JSON API
(JSON API)[http://jsonapi.org/]是移除数据中冗余部分并返回数据给客户端的一种标准。比如考虑网站上通过`API`提交大量博客的场景。每个博客有唯一的内容，比如标题，标识和文本。但是每个博客都会有冗余信息，比如作者信息。
通常这些情况下我们在每次提交请求的时候都会收到冗余的作者信息。如果我们的博客大部分都是同一个作者，就会浪费很多资源。`JSON API`允许我们定义不同类型的资源之间的关系来消除冗余。看下嘛的例子：
``` JSON
{
  "data": [
    {
      "type": "articles",
      "id": "1",
      "attributes": {
        "title": "Article Title",
        "body": "Content"
      },
      "relationships": {
        "author": {
          "data": {
            "id": "42",
            "type": "people"
          }
        }
      }
    }
  ],
  "included": [
    {
      "type": "people",
      "id": "42",
      "attributes": {
        "name": "John",
        "age": 80
      }
    }
  ]
}
```

#### GraphQL
(GraphQL)[http://graphql.org/]是facebook开发的`API`标准。它包含了自定义请求数据格式。通常应答是以`json`的形式返回。请求结构里面要求客户端指明需要返回的属性字段，所以白名单属性是内置的。这点满足移动客户端只需要返回重要的字段，使用较少字节数的需求。
`GraphQL`另外一个特性就是从不同集合里面返回的应答消息里的请求字段是相互关联的。这点在设计从不同服务请求数据接口是非常吸引人。`GraphQL`能够通过一次请求就能获取需要的所有信息，节省了客户端从不同的集合通过多次请求获取数据。
通过带body的`POST`的形式请求单一资源信息。`GraphQL`不是`RESTful HTTP`的实现，可以和`HTTP`完全独立出来。下面是`GraphQL`的请求示例：
```
{
  user(id: "tlhunter") {
    id
    name
    photo {
      id
      url
    }
    friends {
      id
      name
    }
  }
}
```
下面是相对应的回复消息：
```
{
  "data": {
    "user": {
      "name": "Thomas Hunter II",
      "id": "tlhunter",
      "photo": { "id": "12", "url": "http://im.io/12.jpg" },
      "friends": [
        { "name": "Rupert Styx", "id": "rupertstyx" }
      ]
    }
  }
}
```