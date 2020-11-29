#!/bin/bash
# Shell编程综合案例
# 需求分析：
# 1) 每天凌晨2：10备份日志文件到/opt/testShell/logBackup
# 2) 备份开始和备份结束后能够给出相应的提示信息
# 3）备份后的文件要求以备份时间为文件名，并打包成.tar.gz的形式，比如：2018-03-12_230201.tar.gz
# 4) 在备份的同时，检查是否有10天前备份的数据库文件，如果有就将其删除

# 文件备份的路径
BACKUPLOG=/opt/testShell/logBackup
# test.log 文件夹的路径
LOGPATH=/opt/testShell/log
# 以当前的时间作为文件名
DATETIME=`date +%Y_%m_%d_%H%M%S`
#echo $DATETIME
echo "======开始备份======"
echo "======备份的路径是：$BACKUPLOG/$DATETIME.tar.gz ======"
# 创建备份的路径
if [ ! -d "$BACKUPLOG/$DATETIME" ]
then
	mkdir -p "$BACKUPLOG/$DATETIME"
fi
# 复制日志到该路径
cp /opt/testShell/log/test.log $BACKUPLOG/$DATETIME
# 打包备份文件
cd $BACKUPLOG
tar -zcvf $DATETIME.tar.gz $DATETIME
# 删除临时目录
rm -rf $BACKUPLOG/$DATETIME
# 删除10天前的备份文件
find $BACKUPLOG -mtime +10 -name "*.tar.gz" -exec rm -rf {} \;
echo "======备份文件成功======"
