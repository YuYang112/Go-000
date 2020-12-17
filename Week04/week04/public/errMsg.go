package public

func ErrMsg(code int) string{
	var msg string
	switch code {
	case -1:
		msg = "系统异常"
	case 999:
		msg = "系统维护中"
	case 0:
		msg = "请求成功"
	case 1000:
		msg = "缺少参数"
	case 1001:
		msg = "参数验证失败"
	case 1002:
		msg = "没有相关数据"
	case 1003:
		msg = "保存失败"
	case 1004:
		msg = "删除失败"
	case 1005:
		msg = "获取失败"
	case 1006:
		msg = "创建错误"
	case 1007:
		msg = "没有权限"
	}

	return msg
}
