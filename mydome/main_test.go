package main

import (
	"fmt"
	"testing"
)

/**
------test的要素------
1、文件名必须以test结尾，否则执行go test 命令时该文件将不会被执行
2、文件中的方法必须以Test开头，否则该方法将不会被执行
3、方法的参数必须传入 testing.T|B的指针

  性能测试
4、方法的参数必须传testing.B
5、方法名必须以Bench开头
6、


*/

func TestMyprint(t *testing.T) {
	re := Myprint()
	fmt.Println("re:", re)
	if re != 5 {
		t.Errorf("Myprint error")
	}
}
