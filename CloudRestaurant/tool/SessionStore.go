/*
	session
*/
package tool

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// 初始化 session 操作
func InitSession(engine *gin.Engine) {

	redisCfg := GetConfig().RedisConfig

	store, err := redis.NewStore(10, "tcp", redisCfg.Addr+":"+redisCfg.Port, "", []byte("secret"))
	if err != nil {
		fmt.Println(err.Error())
	}

	engine.Use(sessions.Sessions("mysession", store))

}

// set session 操作
func SetSession(context *gin.Context, key interface{}, value interface{}) error {

	session := sessions.Default(context)

	if session == nil {
		return nil
	}

	session.Set(key, value)
	return session.Save()

}

// get session 操作
func GetSession(context *gin.Context, key interface{}) interface{} {

	session := sessions.Default(context)
	return session.Get(key)

}
