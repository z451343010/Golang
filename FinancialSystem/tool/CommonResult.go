/*
	定义公用的返回结果
*/
package tool

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS int = 0 // 操作成功
	FAILED  int = 1 // 操作失败
)

// 请求成功
func Success(context *gin.Context, v interface{}) {

	context.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"smg":  "成功",
		"data": v,
	})

}

// 请求失败
func Failed(context *gin.Context, v interface{}) {

	context.JSON(http.StatusOK, map[string]interface{}{
		"code": FAILED,
		"smg":  v,
	})

}
