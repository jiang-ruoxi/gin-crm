### 运行前准备
（1）将根目录下的gin-crm.sql 文件导入到你的本地数据库，数据库的名字为gin-crm

（2）在conf文件夹下面配置好ini文件
```
[database]

TYPE = mysql  #数据库类型为mysql

USER = "root" #数据库的用户名

PASSWORD = "123456" #数据库的密码

HOST = 127.0.0.1:3306 #数据库的ip以及端口

NAME = gin-crm #数据库的名称

TABLE_PREFIX = #表前缀
```

### 运行

首先需要在conf文件夹中配置好ini文件

```
go mod tidy
```

```
go mod vendor
```

```
go run main.go
```
由于值简单做了几个菜单，还为进行登录操作的验证
项目访问地址， 访问http://localhost:8080/admin/index
