/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-9
 * Time: 下午 16:02
 */
package _package

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

func CreateInBatches(tables string, keys []string, data [][]string) string {
	sql := fmt.Sprintf("INSERT INTO `%s` (%s) values  ", tables, strings.Join(keys, ","))
	for _, v := range data {
		sky := ""
		for _, k := range v {
			sky = fmt.Sprintf("%s\"%s\",", sky, k)
		}
		sky = strings.Trim(sky, ",")
		sql = fmt.Sprintf("%s (%s),", sql, sky)
	}
	return strings.Trim(sql, ",")
}

//智能拼接Sql, Sql *xorm.Session
func IntelligenceSql(where map[string]interface{}, Sql *gorm.DB, effect bool) *gorm.DB {
	for country := range where {
		Sql = splicingSql(where[country], Sql)
	}
	if effect == true {
		Sql.Where("effect = ? ", EFFECT)
	}
	return Sql
}

func splicingSql(t interface{}, Sql *gorm.DB) *gorm.DB {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)
		tagName := s.Index(0).Interface()
		switch s.Index(1).Interface() {
		case "FIND_IN_SET":
			//	Sql.Where(fmt.Sprintf("find_in_set(%s,%s)", tagName, s.Index(1).Interface()) , s.Index(2).Interface())
			//	Sql.Where("find_in_set('22', `cate_id`)")
			break
		case "in":
			Sql.Where(fmt.Sprintf("%s in (%s)", tagName, s.Index(1).Interface()), s.Index(2).Interface())
		default:
			Sql.Where(fmt.Sprintf("%s %s", tagName, s.Index(1).Interface())+" ? ", s.Index(2).Interface())
			break
		}
	}
	return Sql
}

type Page struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func Paging(Sql *gorm.DB, page Page) *gorm.DB {
	Sql.Limit(page.Limit).Offset(page.Limit*page.Page - page.Limit)
	return Sql
}
