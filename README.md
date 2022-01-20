# beego-imooc
学习beego第一课

# GO 语言框架实战相关命令（2）

## hello wold

### 编码

```
package main  // 声明 main 包，表明当前是一个可执行程序

import "fmt"  // 导入内置 fmt 包

func main(){  // main函数，是程序执行的入口
    fmt.Println("Hello World!")  // 在终端打印 Hello World!
}

```

### 编译

```
// 常见命令
go build  //示将源代码编译成可执行文件。
go build hello // 便宜指定目录
go build -o heiheihei.exe // 输出指定格式的文件

go install //表示安装的意思，它先编译源代码得到可执行文件，然后将可执行文件移动到GOPATH的bin目录下。因为我们的环境变量中配置了GOPATH下的bin目录，所以我们就可以在任意地方直接执行可执行文件了。
go get  // 下载代码所需要的包、依赖库等

```

### 运行

双加运行可执行文件即可

## beego 实战
文档地址: https://www.topgoer.com/beego%E6%A1%86%E6%9E%B6/


### 下载&配置
```

// 进入go path 的 src目录
go get -u github.com/astaxie/beego
go get -u github.com/astaxie/bee // bee 工具
go get -u github.com/beego/bee // bee 工具

bee version

```

### 新建项目
新建项目的方式是使用bee工具带的脚手架新建，命令如下：

```
bee new imooc
```

### 启动

新建后直接启动
```
bee run

```
这个过程中可能会提示安装一些东西，或配置一些命令 ，比如：

```
F:\Workplace\go\src\imooc>bee run
______
| ___ \
| |_/ /  ___   ___
| ___ \ / _ \ / _ \
| |_/ /|  __/|  __/
\____/  \___| \___| v1.12.0
2021/12/03 15:51:55 INFO     ▶ 0001 Using 'imooc' as 'appname'
2021/12/03 15:51:55 INFO     ▶ 0002 Initializing watcher...
go: github.com/astaxie/beego@v1.12.1: missing go.sum entry; to add it:
        go mod download github.com/astaxie/beego
2021/12/03 15:51:57 ERROR    ▶ 0003 Failed to build the application: go: github.com/astaxie/beego@v1.12.1: missing go.sum entry; to add it:
        go mod download github.com/astaxie/beego
        
```

至于 go mod 命令的具体作用 参照 go help 吧，暂时没看懂


```
// 还有可能让你安装这个
go get github.com/shiena/ansicolor

go get github.com/astaxie/beego 
... 


```

### 启动成功
```
F:\Workplace\go\src\imooc>bee run
______
| ___ \
| |_/ /  ___   ___
| ___ \ / _ \ / _ \
| |_/ /|  __/|  __/
\____/  \___| \___| v1.12.0
2021/12/03 15:54:08 INFO     ▶ 0001 Using 'imooc' as 'appname'
2021/12/03 15:54:08 INFO     ▶ 0002 Initializing watcher...
imooc/controllers
imooc/routers
imooc
2021/12/03 15:54:22 SUCCESS  ▶ 0003 Built Successfully!
2021/12/03 15:54:22 INFO     ▶ 0004 Restarting 'imooc.exe'...
2021/12/03 15:54:22 SUCCESS  ▶ 0005 './imooc.exe' is running...
2021/12/03 15:54:24.640 [I] [asm_amd64.s:1371]  http server Running on http://:8080
2021/12/03 15:54:47.037 [D] [server.go:2887]  |            ::1| 200 |     4.0393ms|   match| GET      /     r:/
2021/12/03 15:54:47.117 [D] [server.go:2887]  |            ::1| 200 |      997.7µs|   match| GET      /static/js/reload.min.js

```

## 编码

### beego 框架脚手架 的初始代码接口

```
F:\Workplace\go\src\imooc>tree
F:.
├─conf
├─controllers
│  ├─default.go
├─models
├─routers
│  ├─router.go
├─static
│  ├─css
│  ├─img
│  └─js
├─tests
├─views
└─main.go


```

### 使用bee 工具自动生成代码

```

bee generate scaffold user -fields="id:int64,name:string,gender:int,age:int" -driver=mysql -conn="root:@tcp(127.0.0.1:3306)/imooc"

会根据表结构自动生成controller，models,

```


最后run一下

## 参考&鸣谢

https://www.imooc.com/learn/1053
