#!/bin/bash

# 遍历数组
testArray=(A B C D)
# 方式一
echo "方式一"
for(( i=0;i<${#testArray[@]};i++))
do
	echo ${testArray[i]}
done
# 方式二
echo "方式二"
for element in ${testArray[@]}
do
	echo $element
done
# 方式三 （带数组下标的遍历）
echo "方式三"
for i in "${!testArray[@]}"
do
	printf "%s\t%s\n" "$i" "${testArray[i]}"
done
# 方式四 （While循环遍历数组）
echo "方式四"
i=0
while [ $i -lt ${#testArray[@]} ]
do
	echo ${testArray[i]}
	# printf "%s\t%s\n" "$i" "${testArray[i]}" 带数组下标输出
	let i++
done
