#### 写在前面
基于`HTTP`协议的`Web API`是时下最为流行的一种分布式服务提供方式。之前在审查写的 `Restful API` 的时候有些东西没理解清楚，后来查阅一些资料的时候看到 `Thomas Hunter II `在17年6月份发表介绍`HTTP API`设计的四篇文章（`Requests`,`Responses`,`Bodies`,`API Standards`）和该作者写的一本小册子`Consumer Centric API Design`，内容比较清晰，所以打算翻译下保留下来。第一次翻译这么长的文章，文字功底也比较拙劣...orz

点击阅读原文查看原英文，可能需要翻墙...orz

#### JSON (JavaScript Object Notation)
`JSON`是当前主流`HTTP API`首选的数据序列化格式。`JSON`的格式十分简单，几分钟就可以成为专家。格式十分简单，大多数语言下都能比较容易序列化和反序列化。本质上来说`json`是以`key/value`键值对的形式存在的，数值类型包括字符串，数字，布尔，空值或者数组和对象等多种类型。
相比较于其他需要复杂序列化逻辑和啰嗦语法查询数据的数据格式，如`xml`，`json`十分容易使用。下面是一个`json`文档示例：
``` json
{
  "strings": "value",
  "numbers": 12,
  "moreNumbers": -0.1,
  "booleans": true,
  "nullable": null,
  "array": [1, "abc", false],
  "object": { "very": "deep" }
}
```

#### Attribute Name Casing
在使用`json`或者任意一种格式的时候，总是需要被考虑或者也是常常被忽略的一个问题就是属性字段名命名。下面列出三种选项：
* snake_case: 使用最多的字节不过最易读
* PascalCase: 需要按最多`shift`次数
* camelCase: 按下较少`shift`
每一种命名方式都不会比其他更好或更差。不同的平台也有自己的偏好，比如微软`.NET`中`PascalCase`比较常见，`snake_case`和`camelCase`在开源项目中广泛使用。
无论你选择支持何种格式，在所有的`Requests`和`Responses`中都保持一致很重要。不要混用这些风格，这样会困扰开发者去记忆何种场景下使用何种风格。即使你的服务只是封装了各种其他混用了风格的服务，你的服务也应该选择一种风格而且转化过来。项目中一般会让内部变量，属性名和`API`风格保持一致，虽然这样做也不会让使用者带来任何好处。

#### Booleans
布尔字段命名时应该使用积极向上的名词。在正面和负面的单词之间切换是一个令人困惑的事情，需要常常查看文档。下面是一些例子：
* 使用`enable`替换`disable`
* 使用`public`替换`private`
在命名一个布尔字段的时候，有的时候会过度强调是布尔类型的。例如，没必要命名为`is_cool`或`cool_flag`，只需要命名为`cool`就够了。如果使用标志字段，那么整个`API`都应该保持一致。

#### Timestamps
序列化时间字段有多种不同的标准格式。然而其中最重要的一种格式就是`ISO 8601`。下面是几个示例：
* "2017-06-15T04:23:46+00:00": Using a numeric offset from UTC
* "2017-06-15T04:23:46Z": Using “Zulu” UTC time
* "2017-06-15T04:23:46.987Z": Variable precision for milliseconds
这种格式使用字符串来表示。一个重要的特性就是，假设所有时间属于同一个时区，当按照字母表排序的时候是按照时间先后顺序排序的。易于人类可读，格式简洁。
另外一个选择就是`UNIX`时间戳，不可读且容易混淆。一个带毫秒精度的示例如`1493268311123`和一个不带精度的`1493268311`是不同的。除非在文档开头标注精度，否则很难知道是否带有额外的时间精度。

#### Identifiers
标识符一般都用字符串来表示。即使用的是自增的整数来表示标识符，也应该转换为字符串。下面是一个来自`chess`的tweet，他们的API使用整形作为标识字段类型导致IOS上的应用宕机：
>“32-bit iOS devices are experiencing issues due to limitations interpreting game IDs over 2,147,483,647. Fix should be out in 48 hours :)”
— @chesscom
理论上来说有两种类别的标识符可供选择。每种都有优缺点。在特定的情况下选择特定的标识类型。
* Incremental (e.g. Integer, Base62): Efficient to store
* Random (e.g. UUID): Difficult to guess values or count collection size
随机的IDs有一个额外的特性很吸引人:在一个高可用的分布式解决方案中，如果你不想有一个中心节点记录IDs而且不想做自增来产生IDs.如果IDs大小的有足够熵的话，产生冲突的可能会非常小。

#### Versioning
在重构会维护`API`的时候，常常会因为改变导致客户端出现不兼容的情况。如果这些变更是特别重要的，而且希望旧的客户端在旧版本的`API`下依然能工作，需要同时支持一个新的版本。这样客户端可以在合适的时间切换到新的版本。
有三种流行的方式来标识`API`的版本：
* URL Segment (LinkedIn, Google+, Twitter): https://api.example.org/v1/*
* Accept Header (GitHub): Accept: application/json+v1
* Custom Header (Joyent CloudAPI): X-Api-Version: 1 
新版本的`API`只需要保证旧版本下的`API`能正常运行就行了。例如，添加一个资源或集合新的可选属性不应该有版本冲突。改变数据类型，添加必填参数，移除一个参数或集合，改变默认的筛选条件(例如默认的所有记录到每页100条记录)，这些改变都会导致向后不兼容。
当决定废弃一个旧版本的`API`的时候，应该给使用者足够的时间和提醒。提前把更改通知使用方并且提供可能的迁移方案。如果完整废弃一个就版本`API`应当提供一个截止时间。