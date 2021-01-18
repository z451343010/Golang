package dao

import (
	"fmt"
	"gin/CloudRestaurant/model"
	"gin/CloudRestaurant/tool"
)

type FoodCategoryDao struct {
	*tool.Orm
}

// 从数据库中查询所有的食品种类，并返回
func (fcd *FoodCategoryDao) QueryCategories() ([]model.FoodCategory, error) {

	var categories []model.FoodCategory
	err := fcd.Engine.Find(&categories)
	if err != nil {
		fmt.Println("查询所有的食品种类失败")
		return nil, err
	}

	return categories, nil

}
