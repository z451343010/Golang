#!/bin/bash
readonly A=100
echo "A = $A"
#readonly 类型的变量不能unset，否则程序会报错
unset A
echo "A = $A"
