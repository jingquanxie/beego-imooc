package main

import (
	_ "imooc/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// bee generate scaffold user -fields="id:int64,name:string,gender:int,age:int" -driver=mysql -conn="root:@tcp(127.0.0.1:3306)/imooc"
func main() {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/imooc")
	beego.Run()
}
