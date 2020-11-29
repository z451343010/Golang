#!/bin/bash

# 从 1 累加到 100
SUM=0
for ((i=1;i<=100;i++))
do
	SUM=$[SUM + i]
done
echo "SUM = $SUM"
