package main

import "TreeHole/dao"

func main() {
	// 连接
	dao.Client()
	// 初始化表
	dao.AutoTables()

}
