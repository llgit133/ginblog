package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCategory 查询分类是否存在
// select * from categories where name = 'newName';
func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED //2001
	}
	return errmsg.SUCCESS
}

// CreateCate 新增分类
// insert into categories (name) values ('newName');
func CreateCate(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS
}

// GetCateInfo 查询单个分类信息
// SELECT * FROM categories where id = ?;
func GetCateInfo(id int) (Category, int) {
	var cate Category
	db.Where("id = ?", id).First(&cate)
	return cate, errmsg.SUCCESS
}

// GetCate 查询分类列表
// SELECT * FROM categories LIMIT 10 OFFSET 10;
func GetCate(pageSize int, pageNum int) ([]Category, int64) {
	var cate []Category
	var total int64
	err = db.Find(&cate).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	db.Model(&cate).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}

// EditCate 编辑分类信息
// UPDATE categories SET name = 'newName', age = 30 WHERE id = 5;
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	err = db.Model(&cate).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteCate 删除分类
// delete from categories where id = 5;
func DeleteCate(id int) int {
	var cate Category
	err = db.Where("id = ? ", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
