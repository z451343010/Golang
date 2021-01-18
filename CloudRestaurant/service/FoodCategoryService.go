package service

import (
	"gin/CloudRestaurant/dao"
	"gin/CloudRestaurant/model"
	"gin/CloudRestaurant/tool"
)

type FoodCategoryService struct {
}

// 获取食品类别
func (fcs *FoodCategoryService) Categories() ([]model.FoodCategory, error) {

	foodCategoryDao := dao.FoodCategoryDao{tool.DbEngine}
	return foodCategoryDao.QueryCategories()

}
