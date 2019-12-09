# 简单 web 服务与客户端开发实战小结

<div align=right style="font-weight:bold;font-size:15px">17软工 马欢 16340165</div>

**本次项目仓库地址： https://github.com/SYSU-SimpleBlog**

#### 分工

在本次作业中，我们小组共6人实现了简单的web服务端和客户端，其中分工2人负责前端，4人负责api设计和后端设计。我在实验中参加后端小组，主要编写了登陆和评论相关api，测试保证了登陆api的正确性，协助另外的同学修改和测试评论部分。文章相关api的编写和测试由另外两名同学进行，这两名同学也完成了摘取真实网页博客内容的工作，而我辅助完成对评论部分的相关填充。

## 实验感想

一改以往将实验感想写在最后的安排，这次我希望将感想写在最前面，以此表示对小组中每一个成员的感谢。由于临近期末，大家的时间都安排比较紧，因此在完成其余期末相关作业同时小组成员间还要兼顾其余成员的时间便显得非常不易。在本次作业中，由于api出来后所有人才能完成接下来的任务，其中某些又是有着依赖关系的（如创建评论需要用户的登陆信息即token等），而前后端耦合后也需要按具体效果进行项目的微调。因此即使做到了前后端分离开发，每个人的工作和其他人还是有着不可分割的联系。为了不影响其他人的进度，大家都在尽快完成自己的工作，还会放下手头工作帮助遇到问题的同学解决难题，甚至会通宵赶进度、生着病去完医院回来赶紧接手工作、在群里帮助组员工作到了晚饭时间才赶去吃顿中午饭。而我在工作中也认识了两个不熟悉的新朋友，虽然这次作业任务繁重，我在应用所学技能的同时还很好地感受到了团队协作的力量，确实受益匪浅。在此也想对每一个小组成员表示感谢！

## API设计

在本次实验中我们采用 REST v3 风格，设计了 6个API 服务，并使用 swagger-editor 来进行对 API文档的编写，其使用yaml语法，例如对 signin 而言：

```yaml
  /user/signin:
    get:
      tags:
      - "user"
      summary: "sign in"
      description: "Check user with username and password"
      operationId: "SignIn"
      produces:
      - "application/json"
      responses:
        "200":
          description: "Successful Operation"
          schema:
            $ref: "#/definitions/inline_response_200"
        "404":
          description: "Not Found"
          schema:
            $ref: "#/definitions/inline_response_404"
```

生成如下文档部分：


<img src="C:\Users\86159\Desktop\fuwu\apisign.png" alt="apisign" style="zoom:67%;" />

而完整文档可通过以下链接查看：

**https://sysu-simpleblog.github.io/Proj-doc/**

## 实现过程

无论是哪个api，都涉及到了与数据库之间的交互，对于初始平台需要在 BoltDB 中生成初始数据，例如对于 user：

```go
func CreateUser() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("User"))
		if b == nil {
			//create table "xx" if not exits
			b, err = tx.CreateBucket([]byte("User"))
			if err != nil {
				log.Fatal(err)
			}
		}

		//insert rows
		for i := 0; i < 10; i++ {
			err := b.Put([]byte("user"+strconv.Itoa(i)), []byte("pass"+strconv.Itoa(i)))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("user"+strconv.Itoa(i), "pass"+strconv.Itoa(i))
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
```

对于数据库和表的存在信息进行了检测，随后在初始时增添了 user0 ~ user9 的账户，其密码对应 pass0 ~ pass9。随后便可利用这些信息进行登陆：

由于本次平台测试部署在本地，在使用 Post 方法在对静态资源进行请求时会得到 405 的服务拒绝 ，在此改用 Get 方法模拟登陆过程，来完成前端所发送拼接字符串的解析和校验过程，在此使用的字符串格式借鉴对于 uri 的过滤方式，来在 request 中进行格式化的字符串切分：

```go
u, err := url.Parse(r.URL.String())
fatal(err)
m, _ := url.ParseQuery(u.RawQuery)
fmt.Println(m)
var user User
user.Username = m["username"][0]
user.Password = m["password"][0]
```

随后在 BoltDB 中进行输入用户信息和存储信息的匹配过程：

```go
err = db.View(func(tx *bolt.Tx) error {
    	b := tx.Bucket([]byte("User"))
	if b != nil {
    		v := b.Get([]byte(user.Username))
		if ByteSliceEqual(v, []byte(user.Password)) {
    			return nil
		} else {
    			return errors.New("Username and Password do not match")
		}
	} else {
    		return errors.New("Username and Password do not match")
	}
})
```

在匹配成功后需要创建并返回 token 以供后续api校验方法的使用：

```go
token := jwt.New(jwt.SigningMethodHS256)
claims := make(jwt.MapClaims)
claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
claims["iat"] = time.Now().Unix()
token.Claims = claims
```

其登陆过程进行后端测试结果如下：

<img src="C:\Users\86159\Desktop\fuwu\curlSignin.png" alt="curlSignin" style="zoom:80%;" />

以 createComment 为例，在对其他 api 进行方法操作时，首先对于 request 请求的token进行解析：

```go
token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
	func(token *jwt.Token) (interface{}, error) {
    		fmt.Println(token)
		return []byte(comment.Author), nil
	})
```

再通过以下方法做到服务端对于客户端保存当前登陆用户的 token 的校验：

```go
if token.Valid {
    	err = db.Update(func(tx *bolt.Tx) error {
    		b, err := tx.CreateBucketIfNotExists([]byte("Comment"))
		if err != nil {
    			return err
		}
		id, _ := b.NextSequence()
		encoded, err := json.Marshal(comment)
		var str string
		str = strconv.Itoa(Id) + "_" + strconv.Itoa(int(id))
		return b.Put([]byte(str), encoded)
	})
	if err != nil {
    		response := InlineResponse404{err.Error()}
		JsonResponse(response, w, http.StatusBadRequest)
		return
	}
	JsonResponse(comment, w, http.StatusOK)
} else {
    	response := ErrorResponse{"Token is not valid"}
	JsonResponse(response, w, http.StatusUnauthorized)
}
```

通过这种方式，实际上完成了 JWT 方案，在用户验证登陆成功后服务器返回 token 并被保存在客户端，而用户后续的请求操作则通过客户端中取出 token 并返还服务器进行验证，以对用户操作进行身份校验。

我们编写了测试函数再后端进行各个 api 的功能检测，其验证结果的部分显示如下：

![3](C:\Users\86159\Desktop\fuwu\backtest.png)

## 实验结果

而整个实验前后端耦合后的最终结果显示如下：

登陆界面： (/sigin)

![signin](C:\Users\86159\Desktop\fuwu\signin.png)

登陆后进入个人博客列表： (/articleList)

![ariticleList](C:\Users\86159\Desktop\fuwu\ariticleList.png)

删除文章，由右下角可以看出其实现了分页功能：![delete](C:\Users\86159\Desktop\fuwu\delete.png)

删除结果：可以发现上一个图中的最后一条博客已无

![afterdelete](C:\Users\86159\Desktop\fuwu\afterdelete.png)

点进文章查看博文内容：

![viewarticle](C:\Users\86159\Desktop\fuwu\viewarticle.png)

查看博文最底部评论部分：

![comments](C:\Users\86159\Desktop\fuwu\comments.png)

发表一条新的评论：（发表内容如上图可看出）

![afteradd](C:\Users\86159\Desktop\fuwu\afteradd.png)

可以看出评论正常添加，且评论部分也实现了分页功能。
