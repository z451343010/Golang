package service

import (
	"fmt"
	"gin/CloudRestaurant/dao"
	"gin/CloudRestaurant/model"
	"gin/CloudRestaurant/tool"
	"strconv"
)

type ShopService struct {
}

// 查询商铺列表数据
func (shopService *ShopService) ShopList(long, lati string) []model.Shop {

	longitude, err := strconv.ParseFloat(long, 10)
	if err != nil {
		fmt.Println("转换经度时出错")
		return nil
	}

	latitude, err := strconv.ParseFloat(lati, 10)
	if err != nil {
		fmt.Println("转换纬度时出错")
		return nil
	}

	shopDao := dao.ShopDao{tool.DbEngine}
	return shopDao.QueryShops(longitude, latitude)

}
