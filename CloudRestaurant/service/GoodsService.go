package service

import (
	"fmt"
	"gin/CloudRestaurant/dao"
	"gin/CloudRestaurant/model"
	"gin/CloudRestaurant/tool"
	"strconv"
)

type GoodsService struct {
}

/*
	获取商家的食品信息列表
*/
func (gs *GoodsService) GetFoods(shop_id string) []model.Goods {

	goodsDao := dao.GoodsDao{tool.DbEngine}

	shopId, err := strconv.Atoi(shop_id)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	goods, err := goodsDao.QueryFoods(int64(shopId))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return goods

}
