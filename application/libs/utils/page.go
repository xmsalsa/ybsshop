/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-10 15:14:01
 */
package utils

import (
	"errors"
	"fmt"

	"github.com/kataras/iris/v12"
)

const (
	PAGE_MIN_LIMIT     = 1     // 分页最小值
	PAGE_MAX_LIMIT     = 1000  // 分页最大值
	PAGE_EXCEPT_LIMIT  = 99999 // 获取所有时使用
	PAGE_DEFAULT_LIMIT = 10
	PAGE_DEFAULT       = 1
	PAGE_MIN_PAGE      = 1 // 页码最小值

)

type Pages struct {
	Offset int `json:"offset" comment:""`
	Limit  int `json:"limit" comment:""`
}

// 获取所有
func PagesAll() Pages {
	return Pages{
		Offset: 0,
		Limit:  PAGE_MAX_LIMIT,
	}
}

func PageInit(ctx iris.Context) (Pages, error) {
	var page Pages
	var err error
	page, err = _init(ctx.URLParamIntDefault("page", PAGE_DEFAULT), ctx.URLParamIntDefault("limit", PAGE_DEFAULT_LIMIT))
	if err != nil {
		return page, err
	}

	return page, nil
}

func _init(page, limit int) (Pages, error) {
	var pages = Pages{
		Offset: limit*page - limit,
		Limit:  limit,
	}
	if pages.Offset < 0 {
		return pages, errors.New(fmt.Sprintf("page不能小于%d", PAGE_MIN_PAGE))
	}
	if pages.Limit < 1 {
		return pages, errors.New(fmt.Sprintf("limit不能超过%d", PAGE_MAX_LIMIT))
	}
	if pages.Limit < PAGE_MIN_LIMIT {
		return pages, errors.New(fmt.Sprintf("limit不能小于%d", PAGE_MIN_LIMIT))
	}
	if pages.Limit > PAGE_MAX_LIMIT && pages.Limit != PAGE_EXCEPT_LIMIT {
		return pages, errors.New(fmt.Sprintf("limit不能超过%d", PAGE_MAX_LIMIT))
	}

	return pages, nil
}
