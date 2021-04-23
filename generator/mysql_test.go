/*
@Time : 2021/4/23 下午3:12
@Author : MuYiMing
@File : mysql_test
@Software: GoLand
*/
package generator

import (
	"fmt"
	"github.com/MIM1993/sqltool/model"
	"testing"
)

func TestMysqlGenerateSql(t *testing.T) {
	qs := model.NewQueryStage(model.AND, false)
	qs.AddQuery("code", "1305", model.Equal, model.Nothing, false)
	qs1 := model.NewQueryStage(model.Nothing, false)
	qs1.AddQuery("administrator", "郭伟", model.Equal, model.Nothing, true)
	s, err := GenerateSql("mdr_community", []*model.QueryStage{qs,qs1})
	if err != nil {
		fmt.Println("GenerateSql err:", err)
		return
	}

	fmt.Println(s)

}
