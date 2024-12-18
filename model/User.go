package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20)" : json:"username"`
	Password string `gorm:"type:varchar(20)" : json:"password"`
	Role     int    `gorm:"type:int" : json:"role"`
}

// json 前后端数据交互

// 检查用户名是否已经存在
// select id from user where username = userName limit 1
func CheckUser(Username string) int {
	var user User
	db.Select("id").Where("username = ?", Username).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 添加用户
// insert into user values(null, userName, password, role)
func CreateUser(data *User) int {
	//data.Password = ScryptPassword(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询用户列表
// 返回User切片
// select * from user limit 10 offset 10
func GetUsers(pageSize int, pageNum int) []User {

	var users []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

func EditUser(id int, data *User) (code int) {

	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

// 1.密码加密--直接调用
func ScryptPassword(password string) string {

	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}
	hashPwd, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(hashPwd)
	return fpw
}

// 2.密码加密 --钩子函数

func (u *User) BeforeSave() {
	u.Password = ScryptPassword(u.Password)
}

// DeleteUser 删除用户
// DELETE FROM users WHERE id = 1;
func DeleteUser(id int) int {

	var user User
	err := db.Where("id = ?", 1).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// CheckLogin 后台登录验证
func CheckLogin(username string, password string) (User, int) {
	var user User
	var PasswordErr error

	db.Where("username = ?", username).First(&user)

	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if PasswordErr != nil {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return user, errmsg.ERROR_USER_NO_RIGHT
	}
	return user, errmsg.SUCCESS
}

// CheckLoginFront 前台登录
func CheckLoginFront(username string, password string) (User, int) {
	var user User
	var PasswordErr error

	db.Where("username = ?", username).First(&user)

	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if PasswordErr != nil {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	return user, errmsg.SUCCESS
}
