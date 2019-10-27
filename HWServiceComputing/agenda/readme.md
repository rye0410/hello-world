# **CLI 命令行实用程序开发实战 - Agenda**

## 实验内容

&emsp;命令行实用程序并不是都像 cat、more、grep 是简单命令。go 项目管理程序，类似 java 项目管理 maven、Nodejs 项目管理程序npm、git 命令行客户端、 docker 与 kubernetes 容器管理工具等等都是采用了较复杂的命令行。即一个实用程序同时支持多个子命令，每个子命令有各自独立的参数，命令之间可能存在共享的代码或逻辑，同时随着产品的发展，这些命令可能发生功能变化、添加新命令等。因此，符合 OCP 原则 的设计是至关重要的编程需求。

## 实验目的

* 熟悉 go 命令行工具管理项目
* 综合使用 go 的函数、数据结构与接口，编写一个简单命令行应用 agenda
* 使用面向对象的思想设计程序，使得程序具有良好的结构命令，并能方便修改、扩展新的命令,不会影响其他命令的代码
* 项目部署在 Github 上，合适多人协作，特别是代码归并
* 支持日志（原则上不使用debug调试程序）

## 实验过程

### 前期准备

&emsp;首先进行环境搭建，本次实验为了方便快捷，我先在本地中部署环境完成了实验，再将它拷贝到goonline中。本次实验为了实现 POSIX/GNU-风格参数处理，–flags，包括命令完成等支持，选用第三方包cobra。在安装cobra包过程中获取失败的部分再本地的$GOPATH/src/golang.org/x目录下直接使用git clone下载资源并安装：

![1](https://github.com/rye0410/hello-world/blob/master/HWServiceComputing/agenda/pic/1.png)

![2](https://github.com/rye0410/hello-world/blob/master/HWServiceComputing/agenda/pic/2.png)

&emsp;自此我们可以先写一个简单的小程序来验证环境是否搭建完成：

&emsp;使用cobra init并导入packet进行初始化，再添加命令register。在第三方生成的rRun回调函数中填写回显用户名的语句并进行调用：

![示例init](https://github.com/rye0410/hello-world/blob/master/HWServiceComputing/agenda/pic/示例init.png)

![实例add-register](https://github.com/rye0410/hello-world/blob/master/HWServiceComputing/agenda/pic/实例add-register.png)

&emsp;可以得到如下结果：

![示例-运行](https://github.com/rye0410/hello-world/blob/master/HWServiceComputing/agenda/pic/示例-运行.png)

&emsp;可以看出成功显示了命令参数传递的user即TestUser。环境检验成功，可以进行后续实验。

### 实验主体

&emsp;agenda开发要求实现至少两个命令，在此挑选：

* **用户注册**

  * 注册新用户时，用户需设置一个唯一的用户名和一个密码。另外，还需登记邮箱及电话信息。

  * 如果注册时提供的用户名已由其他用户使用，应反馈一个适当的出错信息；成功注册后，亦应反馈一个成功注册的信息。

    ---

    调用register命令后显示如下：

    ![注册](https://github.com/rye0410/hello-world/blob/master/HWServiceComputing/agenda/pic/注册.png)首先使用 -h 查看register命令的帮助信息如上，可以看出调用使用

    &emsp;可执行程序 -u 用户名 -p 密码 -e 邮箱 -t 电话

    的形式进行，并随后输入了命令创建 haha 用户，我们可以在位于. \\entity\\data 中的User.txt中查询到用户信息确实被创建：

    ![注册后](https://github.com/rye0410/hello-world/blob/master/HWServiceComputing/agenda/pic/注册后.png)

    此外，功能还对注册时的密码、邮箱和电话格式，以及是否存在重名用户进行了检验，并在注册失败时显示相关错误：

    ![注册失败](https://github.com/rye0410/hello-world/blob/master/HWServiceComputing/agenda/pic/注册失败.png)

* **用户登录**

  * 用户使用用户名和密码登录 Agenda 系统。

  * 用户名和密码同时正确则登录成功并反馈一个成功登录的信息。否则，登录失败并反馈一个失败登录的信息。

    ---

    仍然先用 -h 查询命令使用信息，并根据命令格式进行用户登陆操作。在进行登陆活动时，若登陆信息错误，根据相应错误提示信息：密码错误或不存在该用户。而在登陆成功后将用户信息加入. \\entity\\data 中的Host.txt中进行标记，后续功能可以据其判断当前登陆用户：

    ![登陆](https://github.com/rye0410/hello-world/blob/master/HWServiceComputing/agenda/pic/登陆.png)

* **用户登出**

  * 已登录的用户登出系统后，只能使用用户注册和用户登录功能。

    ---

    较为简单，若用户已登陆则将登陆信息清理，并进行提示：

    ![退出](https://github.com/rye0410/hello-world/blob/master/HWServiceComputing/agenda/pic/退出.png)

* **帮助显示**

  * 显示每个功能的详细信息。

    ---

    除了对于每一个具体命令能显示该命令的详细信息之外，还可以展示所有存在功能信息：

    ![-h帮助](https://github.com/rye0410/hello-world/blob/master/HWServiceComputing/agenda/pic/-h帮助.png)

### 组件

&emsp;agenda约定了entiy和cmd之间的接口服务，在entity中：

* JsonFormat：

  提供Json格式和agenda执行过程中UserList间的转化，使得创建的用户信息可以被保存在本地并在需要时被随时取用。其信息分别以包含所有字段的形式和只包含用户名字段的形式在User.txt和Host.txt文件中存取，帮助程序实现相应功能。

* Storage：

  包含一个User类型的切片，如果说JsonFormat提供了形式的转化，而在此便是在中转操作指定的形式存入内存中，并按需提取。其集中体现在RegisterUser等函数（这些函数通过service被cmd中同功能函数调用）在程序执行过程中的User类型切片和本地User.txt文件间载入载出。

* User：

  提供了对于user的各个信息的检测，如电话号码要11位、邮件的正则表达式、密码和用户名合法位数等等，保障了合法用户的建立，在cmd中注册、登陆等需要对于这些字段进行验证的功能而言都需用到。

---

&emsp;agenda还提供了log服务，用于对于用户操作过程的记录：

![日志](https://github.com/rye0410/hello-world/blob/master/HWServiceComputing/agenda/pic/日志.png)

&emsp;可以分别就上面的命令和下面的日志对照查询操作。

