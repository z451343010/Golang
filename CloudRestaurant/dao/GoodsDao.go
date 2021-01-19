package dao

import (
	"fmt"
	"gin/CloudRestaurant/model"
	"gin/CloudRestaurant/tool"
)

type GoodsDao struct {
	*tool.Orm
}

// 根据商户的id查询商户下所拥有的食品
func (gd *GoodsDao) QueryFoods(shop_id int64) ([]model.Goods, error) {

	var goods []model.Goods

	err := gd.Orm.Where("shop_id = ? ", shop_id).Find(&goods)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return goods, nil

}
