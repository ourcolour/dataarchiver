# dataarchiver
数据库备份

### 使用方法

##### 清理历史记录

格式：

```
set MYSQL_HOME=/usr/local/mysql
set D_HOME=$(cd `dirname $0`; pwd)

${D_HOME}/bin/dbarchiver clearold -o "<待清理文件夹路径>" -d <清理多少天前的>
```

例子：

```
# 清理 /mnt/dump_data 下 31 天之前的数据。
/usr/local/dataarchiver/bin/dbarchiver clearold -o "/mnt/dump_data" -d 31
```

##### 数据库备份

格式：

```
set MYSQL_HOME=/usr/local/mysql
set D_HOME=$(cd `dirname $0`; pwd)

${D_HOME}/bin/dbarchiver backup -H <数据库地址> -P 数据库端口 -u <用户名> -p <密码> -o "<保存文件夹路径>" [-d <指定库名> [-t <指定表名>]] [-c]
```

例子：

```
# 备份所有库
/usr/local/dataarchiver/bin/dbarchiver backup -H "127.0.0.1" -P 3306 -u root -p toor -o "/mnt/dump_data"

# 压缩备份
/usr/local/dataarchiver/bin/dbarchiver backup -H "127.0.0.1" -P 3306 -u root -p toor -o "/mnt/dump_data" -c

# 压缩备份指定库
/usr/local/dataarchiver/bin/dbarchiver backup -H "127.0.0.1" -P 3306 -u root -p toor -o "/mnt/dump_data" -c -d test

# 压缩备份指定库指定表
/usr/local/dataarchiver/bin/dbarchiver backup -H "127.0.0.1" -P 3306 -u root -p toor -o "/mnt/dump_data" -c -d test -t mytable1
```


