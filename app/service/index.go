package service

import "zcw-gin/global"

type IndexService struct {
}

func Index() *IndexService {
	return global.Golbal.GetOrSetService("index", func() interface{} {
		return new(IndexService)
	}).(*IndexService)
}
