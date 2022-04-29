package service

import "github.com/feiyangderizi/ginServer/model/result"

type AppService struct{}

func (service *AppService) GetSecret(appkey string) result.Result {
	if appkey == "" {
		return result.FailWithMsg("appkey为空")
	}
	return result.Success()
}
