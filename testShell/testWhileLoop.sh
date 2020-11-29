#!/bin/bash
# 从命令行输入n，统计从1到n的值
SUM=0
i=0
while [ $i -le $1 ]
do
	SUM=$[ $SUM + $i ]
	i=$[ $i + 1 ]
done
echo "SUM = $SUM"
