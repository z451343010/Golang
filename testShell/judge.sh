#!/bin/bash
# 判断"ok"是否等于"ok"
if [ "ok" = "ok" ]
then
	echo "ok = ok"
fi

# 判断23是否大于22
if [ 23 -gt 22 ] 
then
	echo "23 > 22"
fi

# 判断 /opt/testShell/test.log 这个文件是否存在
if [ -e /opt/testShell/test.log ]
then
	echo "test.log is exist"
fi

# 判断 /opt/testShell/test.log 这个文件有没有读权限
if [ -r /opt/testShell/test.log ]
then
	echo "test.log can be read"
fi 
