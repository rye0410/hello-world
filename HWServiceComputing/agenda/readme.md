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

![1](C:\Users\86159\Desktop\new\1.png)

![2](C:\Users\86159\Desktop\new\2.png)

&emsp;自此我们可以先写一个简单的小程序来验证环境是否搭建完成：

&emsp;使用cobra init并导入packet进行初始化，再添加命令register。在第三方生成的rRun回调函数中填写回显用户名的语句并进行调用：

![示例init](C:\Users\86159\Desktop\new\示例init.png)

![实例add-register](C:\Users\86159\Desktop\new\实例add-register.png)

&emsp;可以得到如下结果：

![示例-运行](C:\Users\86159\Desktop\new\示例-运行.png)

&emsp;可以看出成功显示了命令参数传递的user即TestUser。环境检验成功，可以进行后续实验。



