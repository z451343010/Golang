package service

import (
	"gin/FinancialSystem/dao"
	"gin/FinancialSystem/model"
	"gin/FinancialSystem/tool"
)

type InStockService struct {
}

// 参看某件商品的库存
func (iss *InStockService) FindInStockByGoodsId(goodsId int64) *model.InStock {

	inStockDao := dao.InStockDao{tool.DbEngine}
	return inStockDao.SelectInStockByGoodsId(goodsId)

}
