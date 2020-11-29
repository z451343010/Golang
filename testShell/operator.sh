#!/bin/bash
# Shell 运算符,方式一 $(())
RESULT=$(((2 + 3) * 4))
# Shell 运算符,方式二 $[] (推荐使用方式二)
RESULT1=$[(2 + 3) * 4]
# Shell 运算符,方式三 expr
TEMP=`expr 2 + 3`
RESULT2=`expr $TEMP \* 4` 
echo "RESULT = $RESULT"
echo "RESULT1 = $RESULT1"
echo "RESULT2 = $RESULT2"
