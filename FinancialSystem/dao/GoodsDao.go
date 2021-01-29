package dao

import (
	"gin/FinancialSystem/model"
	"gin/FinancialSystem/tool"
)

type GoodsDao struct {
	*tool.Orm
}

func (gd *GoodsDao) InsertGoods(goods model.Goods) int64 {

	result, err := gd.InsertOne(&goods)
	if err != nil {
		cfg := tool.GetConfig()
		tool.LogRecoder(cfg, "GoodsDao InsertGoods", err)
		return 0
	}

	return result

}

func (gd *GoodsDao) SelectAllGoods() ([]model.Goods, error) {

	var goods []model.Goods

	err := gd.Orm.Find(&goods)
	if err != nil {
		cfg := tool.GetConfig()
		tool.LogRecoder(cfg, "GoodsDao SelectAllGoods", err)
		return nil, err
	}

	return goods, nil

}

// 通过 goodsId 来获取商品
func (gd *GoodsDao) SelectGoodsById(goodsId int64) *model.Goods {

	var goods model.Goods

	_, err := gd.Where("goods_id = ?", goodsId).Get(&goods)
	if err != nil {
		cfg := tool.GetConfig()
		tool.LogRecoder(cfg, "GoodsDao SelectGoodsById", err)
		return nil
	}

	return &goods

}
