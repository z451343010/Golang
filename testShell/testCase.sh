#!/bin/bash
# 如果命令行参数是1.输出“周一”；2，输出“周二”，其它输出“other”
case $1 in
"1")
echo "周一"
;;
"2")
echo "周二"
;;
*)
echo "other"
;;
esac

