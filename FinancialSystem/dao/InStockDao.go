package dao

import (
	"gin/FinancialSystem/model"
	"gin/FinancialSystem/tool"
)

type InStockDao struct {
	*tool.Orm
}

func (isd *InStockDao) SelectInStockByGoodsId(goodsId int64) *model.InStock {

	var inStock model.InStock
	_, err := isd.Where("goods_id = ?", goodsId).Get(&inStock)
	if err != nil {
		cfg := tool.GetConfig()
		tool.LogRecoder(cfg, "InStockDao SelectInStockByGoodsId", err)
		return nil
	}

	return &inStock

}
