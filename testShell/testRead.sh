#!/bin/bash
# 读取控制台的输入
# 1：读取控制台输入的一个值NUM

read -p "请输入一个数num1 = " NUM1
echo "你输入的数为：$NUM1"

# 2：读取控制台输入的一个值NUM，在10秒内输入

read -t 10 -p "请输入一个数num2 = " NUM2
echo "你输入的数为：$NUM2"
