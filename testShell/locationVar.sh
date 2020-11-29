#!/bin/bash
#位置参数变量
# $n 代表1-n的第几个参数，十以上的参数需要用大括号包含，如：${10}
# $0 代表命令行 如：./locationVar.sh
echo $1
# $* 代表命令行中所有的参数，$* 把所有的参数看成一个整体
echo $*
# $@ 代表命令行中所有的参数，不过$@把每个参数区分对待
echo $@
# $# 代表命令行中所有参数的个数
echo $#
