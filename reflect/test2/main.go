package main

func main() {

	// 定义了两个匿名函数
	test1 := func(v1 int, v2 int) {
		// t.Log(v1, v2)
	}
	test2 := func(v1 int, v2 int, s string) {
		// t.Log(v1, v2, s)
	}

	// 定义一个适配器函数用作统一处理接口，其大致结构如下：
	bridge := func(call interface{}, args ...interface{}) {

	}

	// 实现调用test1对应的函数
	bridge(test1, 1, 2)
	// 实现调用test2对应的函数
	bridge(test2, 1, 2, "test2")

}
