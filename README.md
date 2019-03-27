# dataarchiver
数据库备份

### 使用方法

##### 清理历史记录

``` Bash Shell

set MYSQL_HOME=/usr/local/mysql
set D_HOME=$(cd `dirname $0`; pwd)

${D_HOME}/bin/dbarchiver clearold -o "<待清理文件夹路径>" -d <清理多少天前的>

```

```

# 清理 /mnt/dump_data 下的 31 天之前的数据。
/usr/local/dataarchiver/bin/dbarchiver clearold -o "/mnt/dump_data" -d 31

```



##### 数据库备份

``` Bash

set MYSQL_HOME=/usr/local/mysql
set D_HOME=$(cd `dirname $0`; pwd)

${D_HOME}/bin/dbarchiver backup -H <数据库地址> -P 数据库端口 -u <用户名> -p <密码> -o "<保存文件夹路径>" [-d <指定库名> [-t <指定表名>]] [-c]

-
```

##### 
