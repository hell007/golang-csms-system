package tool

import (
	"encoding/json"
	"errors"
	redisClient "go-mvc/framework/cache/redis"
	"go-mvc/framework/conf"
	models "go-mvc/framework/models/member"
)

/**
通过token获取用户
*/
func GetUserByToken(token string) (*models.LoginUser, error) {
	var (
		err         error
		jsonU string
		member      = new(models.LoginUser)
	)

	jsonU, err = redisClient.Get(conf.GlobalConfig.RedisPrefix + token).Result()
	if err = json.Unmarshal([]byte(jsonU), &member); err != nil {
		return member, errors.New("会员获取失败")
	}

	return member, nil
}
