#!/bin/bash
# 用户自定义函数
# 计算输入的两个参数之和
function getSum(){
	SUM=$[ $n1 + $n2 ]
	echo "get sum = $SUM"
}
read -p "请输入第一个数：" n1
read -p "请输入第二个数：" n2
# 调用函数
getSum $n1 $n2 
