package service

import (
	"net/http"
	"os"

	"github.com/mux"
	"github.com/negroni"
	"github.com/render"
)

// 服务配置和生成
func NewServer() *negroni.Negroni {
	//支持本服务中html和json响应
	formatter := render.New(render.Options{
		//描述选项类型，在此支持html文件扩展
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})
	//构建negroni中间件库，其支持http.Handle
	n := negroni.Classic()
	//初始化用于执行http请求的后续路由列表
	//使用mx简化http的工作
	mx := mux.NewRouter()

	initRoutes(mx, formatter)
	//http.Handler和negroni.Handler间转换
	n.UseHandler(mx)
	return n
}

//
func initRoutes(mx *mux.Router, formatter *render.Render) {
	//获取环境变量中WEBROOT解析字符串
	webRoot := os.Getenv("WEBROOT")
	if webRoot == "" {
		//解析环境变量失败
		//若成功得到当前文件路径则设置为webROot，否则提示错误
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
		}
	}
	//**************************************************************************
	//访问静态文件系统绑定的本地端口，使用不同后缀名进入相应目录
	//相应目录中注册了处理器函数_XXHandler，对相应请求做出回应
	//定位URL到 webRoot + "/assets/" 为虚拟根目录的文件系统。
	mx.HandleFunc("/unknown", UnknownHandler(formatter))
	mx.HandleFunc("/table", InfoHandler(formatter))
	mx.HandleFunc("/js", JsonHandler(formatter)).Methods("GET")
	//http.Dir为强制类型转化而非函数，将字符串格式对应当文件系统
	//PathPrefix添加前缀路径接受文件服务的Handler，处理http请求进行路由
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))
	//**************************************************************************
}
