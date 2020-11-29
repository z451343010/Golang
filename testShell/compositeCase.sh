#!/bin/bash
# Shell编程综合案例
# 需求分析：
# 1) 每天凌晨2：10备份数据库到/data/backup/db
# 2) 备份开始和备份结束后能够给出相应的提示信息
# 3）备份后的文件要求以备份时间为文件名，并打包成.tar.gz的形式，比如：2018-03-12_230201.tar.gz
# 4) 在备份的同时，检查是否有10天前备份的数据库文件，如果有就将其删除

# 文件备份的路径
BACKUPLOG=/opt/testShell/logBackup
# 以当前的时间作为文件名
DATETIME=`date +%Y_%m_%d_%H%M%S`
echo $DATETIME
