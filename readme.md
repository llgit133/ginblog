go get -u github.com/gin-gonic/gin

数据模型驱动开发
domain drive des

config.go 改为ini
go get gopkg.in/ini.v1

gorm orm框架
go get -u github.com/jinzhu/gorm

异常信息处理
errmsg

controller路由搭建
user , article, category

gorm与数据库交互查询

gobatis 框架 类似于mybatis

code data msg 信息封装
{
"data": {
"ID": 2,
"CreatedAt": "2024-07-05T15:48:25.6184412+08:00",
"UpdatedAt": "2024-07-05T15:48:25.6184412+08:00",
"DeletedAt": null,
"Username": "111",
"Password": "111",
"Role": 111
},
"message": "OK",
"status": 200
}

https://gitee.com/wejectchan/ginblog/blob/master/model/User.go
https://www.bilibili.com/video/BV1oA411e7kM/?spm_id_from=333.788&vd_source=b88b8b3483f8d9fb78e502aa8b359b50

----用户模块编写
用户密码加密

bcrypt
scrypt

AddUser
GetUser
DeleteUser
EditUser

----分类模块编写

常见的
-------------------------------------------------gorm单表查询对应sql：
// First 获取第一条记录（主键升序）
// SELECT * FROM `user` ORDER BY `user`.`id` LIMIT 1
models.DB.Debug().First(&user)

// Last 获取最后一条记录（主键降序）
// SELECT * FROM `user` WHERE `user`.`id` = 1 ORDER BY `user`.`id` DESC LIMIT 1
models.DB.Debug().Last(&user)

// Find 获取所有记录
//SELECT * FROM `user`
users := []models.User{}
models.DB.Debug().Find(&users)

-------------------------------------------------gorm多表查询对应sql：

// 联合查询(left join)
db.Table("go_service_info")
.Select("go_service_info.serviceId as service_id, go_service_info.serviceName as service_name, go_system_info.systemId
as system_id, go_system_info.systemName as system_name")
.Joins("left join go_system_info on go_service_info.systemId = go_system_info.systemId")
.Scan(&results)
fmt.Println(mapToJson(results))

//.Scan(&results) 用于将查询结果扫描到 results 变量中

SELECT
go_service_info.serviceId AS service_id,
go_service_info.serviceName AS service_name,
go_system_info.systemId AS system_id,
go_system_info.systemName AS system_name
FROM
go_service_info
LEFT JOIN
go_system_info
ON
go_service_info.systemId = go_system_info.systemId;

// where
db.Table("go_service_info")
. Select("go_service_info.serviceId as service_id, go_service_info.serviceName as service_name, go_system_info.systemId
as system_id, go_system_info.systemName as system_name")
. Joins("left join go_system_info on go_service_info.systemId = go_system_info.systemId where
go_service_info.serviceId <> ? and go_system_info.systemId = ?", "xxx", "xxx")
.Scan(&results)
fmt.Println(mapToJson(results))

SELECT
go_service_info.serviceId AS service_id,
go_service_info.serviceName AS service_name,
go_system_info.systemId AS system_id,
go_system_info.systemName AS system_name
FROM
go_service_info
LEFT JOIN
go_system_info
ON
go_service_info.systemId = go_system_info.systemId
WHERE
go_service_info.serviceId <> 'xxx'
AND go_system_info.systemId = 'xxx';

----文章模块编写

----用户登录 jwt
go get github.com/golang-jwt/jwt/v4




对content理解
请求的上下文，封装了请求的参数，请求的响应，请求的url参数，请求的header信息，http的cookie信息