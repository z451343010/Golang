package dao

import (
	"fmt"
	"gin/CloudRestaurant/model"
	"gin/CloudRestaurant/tool"
)

const DEFAULT_RANGE = 5

type ShopDao struct {
	*tool.Orm
}

// 从数据库中查询店铺
func (shopDao *ShopDao) QueryShops(longitude, latitude float64) []model.Shop {

	var shops []model.Shop
	err := shopDao.Engine.Where("longitude > ? and longitude < ? and latitude > ? and latitude < ?",
		longitude-DEFAULT_RANGE, longitude+DEFAULT_RANGE,
		latitude-DEFAULT_RANGE, latitude+DEFAULT_RANGE).Find(&shops)

	fmt.Println(shops)

	if err != nil {
		fmt.Println("查询数据库 Shop 出错")
		return nil
	}

	return shops

}
