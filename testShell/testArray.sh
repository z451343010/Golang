#!/bin/bash
# 数组 方式一
my_array=(A B "C" D)
echo ${my_array[2]}

# 数组 方式二
array_name[0]="zhang"
array_name[1]="jun"
array_name[2]="yi"
echo ${array_name[0]}

# 获取数组的长度
test_array[0]=A
test_array[1]=B
test_array[2]=C
test_array[3]=D
echo "数组元素的个数为： ${#test_array[*]}"
echo "数组元素的个数为： ${#test_array[@]}"


