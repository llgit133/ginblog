
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


用户密码加密

bcrypt
scrypt