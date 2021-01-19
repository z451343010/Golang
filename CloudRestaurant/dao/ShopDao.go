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
func (shopDao *ShopDao) QueryShops(longitude, latitude float64, keyword string) []model.Shop {

	var shops []model.Shop

	if keyword == "" {

		err := shopDao.Engine.Where("longitude > ? and longitude < ? and latitude > ? and latitude < ?",
			longitude-DEFAULT_RANGE, longitude+DEFAULT_RANGE,
			latitude-DEFAULT_RANGE, latitude+DEFAULT_RANGE).Find(&shops)
		if err != nil {
			fmt.Println("查询数据库 Shop 出错")
			return nil
		}

	} else {

		err := shopDao.Engine.Where("longitude > ? and longitude < ? and latitude > ? and latitude < ? and name like '%"+keyword+"%'",
			longitude-DEFAULT_RANGE, longitude+DEFAULT_RANGE,
			latitude-DEFAULT_RANGE, latitude+DEFAULT_RANGE).Find(&shops)
		if err != nil {
			fmt.Println("查询数据库 Shop 出错")
			return nil
		}

	}

	return shops

}

// 根据店铺ID,查询该店铺提供的服务
func (shopDao *ShopDao) QueryServiceByShopId(shopId int64) []model.Service {

	var services []model.Service
	err := shopDao.Table("service").Join("INNER", "shop_service", "service.id = shop_service.service_id and shop_service.shop_id = ?", shopId).Find(&services)
	if err != nil {
		fmt.Println("根据店铺ID，查询该店铺提供的服务错误")
		return nil
	}

	return services

}
