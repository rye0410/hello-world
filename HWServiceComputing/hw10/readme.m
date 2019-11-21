# 使用RESTful设计博客API

<div align=right style="font-weight:bold;font-size:15px">马欢 16340165</div>

## 协议

指定API与用户间的通信协议，在此为在个人博客平台的API设计，因此使用HTTPs协议。

## 资源与URI表达

指定API部署地址，在此设计为部署于个人博客下的专用*URL：* *https://api.myblog.com* 之下。其应该独立于主*URL：* *https://myblog.com* 以分离服务和资源，避免 “低内聚、高耦合” 的不良系统设计方案。

对于资源组织，围绕 *MyBlog* 实体和所需业务进行以下建模：

**MyBlog 业务与实体（部分）：**

> - MYBlog
>
>   > - users  // 用户
>   >
>   >   > - key //密码
>   >   >
>   >   > - followers // 被关注者集
>   >   > - following // 关注者集
>   >   > - events // 用户登陆后操作集
>   >
>   > - search // 文章搜索集
>   >
>   > - events //登陆前操作集
>   >
>   > - articles
>   >
>   >   > - imgs // 图像内容
>   >   > - contents // 文本内容
>   >
>   > .................................................

以一个article实体为例，查看其属性和关联资源：

*article:*

```
  {
    "title": "服务计算作业",
    "id": 1,
	"author": "16340165马欢"
	"url": "https://api.myblog.com/articles/1",
    "img_url": "https://api.myblog.com/articles/imgs/{img_list}",
    "content_url": "https://api.myblog.com/articles/contents/{contents_list}",
    "author_url": "https://api.myblog.com/users/16340165马欢"
  }
```

又如users集关联着users：

*users：*

```
  {
   	......
    "followers_url": "https://api.myblog.com/users/16340165马欢/followers/users/张三",
    "following_url": "https://api.myblog.com/users/16340165马欢/following/users/李四",
    ......
  }
```



## MyBlog API 版本控制

随着博客服务业务需求变化，需要添加新的资源集合以支持业务升级。其升级过程中可能改变资源集合间关系和资源中数据结构。因此需要引入MyBlog API版本控制服务来指定MyBlog API公开的功能和资源，使得客户端应用程序可以提交定向到指定版本的功能或资源的请求。

对于MyBlog API的版本控制服务，使用标头版本控制方法：

在服务中实现指示资源的版本的自定义标头 ***VMyBlog-Header***。需要客户端应用程序将相应标头添加到所有请求（当省略版本标头时，处理客户端请求的代码可以使用默认值：

假定默认版本为 *myblog-api-version=1.0* ，定义其客户端请求方式如下：

```
Get https://api.myblog.com/article/1 HTTP/1.1
VMyBlog-Header: myblog-api-version=1.0
```

则有服务端回应如下：

```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8

{
	"id":1,
	"title":"服务计算作业",
	"address":"./content/1.html"
}
```

如上所示，在默认版本 *myblog-api-version=1.0* 情况下服务端处理对编号为1博文请求时，返回结果为以 文章编号、文章标题、文章 *Html* 地址组成的 *Json*。 

当使用*myblog-api-version=1.3* 为版本进行相同请求时：

```
Get https://api.myblog.com/article/1 HTTP/1.1
VMyBlog-Header: myblog-api-version=1.3
```

可以得到如下回应：

```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8

{
	"id":1,
	"title":"服务计算作业",
	"address":"./content/1.html", 
	"publish-time":"2019-11-12-15:16:17"
}
```

可以看到对于同一请求做了*myblog-api-version=1.3* 的版本指定后，返回的 *Json* 中相较于 *myblog-api-version=1.0*  版本多了发布时间。

而当请求省略版本标头时：

```
Get https://myblog.com/articles/1 HTTP/1.1
```

会得到默认为 *myblog-api-version=1.0* 指定版本下同样的处理结果。



## 数据表现

MyBlog API 使用 Json 格式数据作为对资源的表达形式，使得其符合标准易于扩展和交互。在定义资源间关系时间，如同之前所给的资源users所表现出来的一样，使用 LInks 链接来表达：

使用link去表达资源之间的联系：

```
  {
   	......
    "followers_url": "https://api.myblog.com/users/16340165马欢/followers/users/张三",
    "following_url": "https://api.myblog.com/users/16340165马欢/following/users/李四",
    ......
  }
```

而 self link 为资源提供了上下文环境，使得客户端可以更加容易获得当前资源位置。因此可以及将以上资源重新设计为：

```
  {
   	......
    "followers_url": "https://api.myblog.com/users/16340165马欢/followers/users/张三",
    "following_url": "https://api.myblog.com/users/16340165马欢/following/users/李四",
    "url": "https://api.myblog.com/users/16340165马欢"，
    ......
  }

```

而对于 MyBlog 中的集合数据资源，比如文章集，将其总体进行self和kind的属性定义，与文章个体的定义形成区分，使得其在检索过滤等操作时均具有更强便利性：

```
{
	"self": "https://api.myblog.com/articles",
	"kind": "平台所有文章",
	"contents": [
	 	 {
	 		"self": "https://api.myblog.com/articles/1/contents/1",
	 		"kind": "技术类文章",
		 	"title": "使用RESTful设计博客API",
		 	"id": "1"
		 },
		 {
			 "self": "https://api.myblog.com/articles/3/contents/66",
			 "kind": "文学类文章",
			 "title": "麦田里的守望者",
			 "id": "66"
		 }
		 ......
	 ]
	 ......
}

```



## 交互与请求

### Token 验证 -- key

使用 POST 操作将 token 存在 http 头部，保证对于user的无状态身份验证请求：

```
{
    type: "POST",
    url: "https://api.myblog.com/users/16340165马欢/key",
    headers: {'Authorization': token}
}

```

这种验证方式由于服务器和用户信息间没有关联，客户端存储的token无状态且不存储session信息，使得负载均衡服务器可以将用户请求传递到任何一台服务器上，具有较强的安全性和可扩展性。

### 阅读 / 浏览 -- articles

一篇文章由2个关联资源组成：文本内容集合和图库（**注意 ：在此并不关联作为 *user* 的作者，因为在建模过程中作者已经作为属性以字符串形式存在于资源中，因此若访问者想要获取作者更多信息，如访问作者主页，则再即时通过在字符串显示处以逻辑计算形式通过字符串在 api.myblog.com/users/?name 下进行过滤检索得到**）。

则需要使用 Get 操作 在 html头 请求2个url资源，将其组建为一个具有图文的文章。

其中 article/{id}/... 中的id是将文章以类别差异进行划分，如article/1/ 下存有技术类文本和技术标签图片，而article/3/ 下存有文学类内容等等。

### 业务组合 -- event

在 MyBlog API 下的复杂操作可以通过 *event* 资源对于原子操作（POST、GET、PUT、PATCH和DELETE）的序列集合进行封装而作为接口使用。即 *event* 资源提供了基于HTTP请求和简单逻辑的高层抽象接口，如 Following 喜欢的作者后平台弹出来他的高频访问 Following。

这意味着若是合法基于HTTP请求和支持语言的规范进行编码，可以类似于Open API一样对 MyBlog API 进行开源扩展。

### 数据分页 --  articles

在 MyBlog API 中，涉及到的最庞大数据集合当属平台文章集合，在进行诸如过滤检索操作时，为了保证平台的安全性，需要进行分批返回的业务，使用数据分页方法实现：

以获取所有文学类文章（ {article_list} 中编号为3 ）为例：

```
GET https://api.myblog.com/articles/3/content?limit=10,offset=0 
Return
{
	"self": "https://api.myblog.com/articles/3/content?limit=10,offset=0",
	"kind": "文学类文章",
	"pageOf": "https://api.myblog.com/articles/3",
	"next": "https://api.myblog.com/articles/3/content?limit=10,offset=10",
	"contents": [...]
}

```

其中 *limit=10, offset=0*  表示从0位置开始每次分页获取10篇 *pageof ：https://api.myblog.com/articles/3* 下的文章，可以看到下次获取时 *(next)* ，其偏移量 *offset* 已经变为10。

### 过滤 / 检索 -- articles, users

如果记录数量很多，服务器不可能都将它们返回给用户。API应该提供参数，过滤返回结果。

当资源记录过多时，除了在之前讲到的数据分页服务对于返回过程的稳定性维护外，还需要考虑到用户需求。MyBlog API 提供参数对于返回进行过滤：

- ?limit=num：指定返回记录的数量
- ?offset=num：指定返回记录的开始位置。
- ?page=2&per_page=num：指定第几页，以及每页的记录数。
- ?sortby=name&order=asc：指定返回结果按照哪个属性排序，以及排序顺序。
- ?type_id=1：指定筛选条件

其设计允许存在冗余，即允许API路径和URL参数偶尔有重复。比如以下两种请求返回内容相同：

```
GET  https://api.myblog.com//articles/1/contents
GET  https://api.myblog.com//articles?article_id=1

```

并且依据需求支持只返回需要的文章信息，例如在对博客平台文章进行文章标题的信息提取时，假设只需要标题和作者作为关键词，则只需要返回这两项数据而不需要对于庞大的 content 和 img 进行返回：

```text
GET https://api.myblog.com/articles?fields=title,author 

```

返回的resource中，只包含name,color,location三种信息。

### 错误处理

MyBlog API 的一系列操作请求需要得到服务器返回的状态码判断请求在服务端的处理情况。若发生错误时，依据HTTP协议规范，会返回 4xx 的状态码，此时需要向用户返回出错信息。以 Error 为键名，将出错信息作为键值进行返回：

- 400 INVALID REQUEST - [POST/PUT/PATCH]：用户发出的请求有错误，服务器没有进行新建或修改数据的操作，该操作是幂等的。
- 401 Unauthorized - [*]：表示用户没有权限（令牌、用户名、密码错误）。
- 403 Forbidden - [*] 表示用户得到授权（与401错误相对），但是访问是被禁止的。
- 404 NOT FOUND - [*]：用户发出的请求针对的是不存在的记录，服务器没有进行操作，该操作是幂等的。
- 406 Not Acceptable - [GET]：用户请求的格式不可得（比如用户请求JSON格式，但是只有XML格式）。
- 410 Gone -[GET]：用户请求的资源被永久删除，且不会再得到的。
- 422 Unprocesable entity - [POST/PUT/PATCH] 当创建一个对象时，发生一个验证错误。

### HATEOAS支持

设计的 MyBlog API 支持Hypermedia，即返回结果中提供链接，连向其他API方法，使得用户不查文档，也知道下一步应该做什么。

比如，当用户向 https://api.myblog.com/ 根目录发出请求，会得到这样一个文档。

```
{
    "link": 
    	{  
            "rel": "collection https://www.myblog.com/articles",  
            "href": "https://api.example.com/articles",  
            "title": "List of article",  
            "type": "application/vnd.yourformat+json"
        }
}

```

文档中有一个 link 属性给出下一步可调用的API，其中 rel 表示API与当前网址的关系，href表示API的路径，title表示 API 的标题，type表示返回类型。其作用便是为客户提供了 url 的层级组织情况。

