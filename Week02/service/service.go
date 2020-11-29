package service

import (
	"Week02/dao"
)

func Service()(int, error){
	// 直接返回err
	return dao.Dao()
}
