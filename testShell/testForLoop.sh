#!/bin/bash
# 打印命令行输出的参数

# 方式一：使用 $*
for i in "$*"
do
	echo "the num is $i"
done

# 方式二：使用 $@
for i in "$@"
do 
	echo "the num is $i"
done
