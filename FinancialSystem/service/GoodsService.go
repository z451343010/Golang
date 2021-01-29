package service

import (
	"gin/FinancialSystem/dao"
	"gin/FinancialSystem/model"
	"gin/FinancialSystem/tool"
)

type GoodsService struct {
}

func (gs *GoodsService) AddGoods(goods model.Goods) int64 {

	goodsDao := dao.GoodsDao{tool.DbEngine}
	return goodsDao.InsertGoods(goods)

}

func (gs *GoodsService) FindAllGoods() ([]model.Goods, error) {

	goodsDao := dao.GoodsDao{tool.DbEngine}
	return goodsDao.SelectAllGoods()

}

func (gs *GoodsService) FindGoodsById(goodsId int64) *model.Goods {

	goodsDao := dao.GoodsDao{tool.DbEngine}
	return goodsDao.SelectGoodsById(goodsId)

}
