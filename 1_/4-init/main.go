package main

import (
	"goStudy/1_/4-init/lib1"
	_ "goStudy/1_/4-init/lib2"
)

func main() {
	//调用包的时候回先执行包中的init方法，之后才会执行main方法
	lib1.Test()

}
