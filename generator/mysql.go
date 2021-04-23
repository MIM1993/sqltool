/*
@Time : 2021/4/23 下午1:41
@Author : MuYiMing
@File : mysql
@Software: GoLand
*/
package generator

import (
	"fmt"
	"github.com/MIM1993/sqltool/internal"
	"github.com/MIM1993/sqltool/model"
	"strings"
)

//OrderBy 处理排序字段
func OrderBy(order []model.OrederBy) {
	//if len(order) <= 0 {
	//	return db
	//}
	//
	//for _, v := range order {
	//	tmpStr := ""
	//	if v.Order >= 0 {
	//		tmpStr = fmt.Sprintf("%s ASC", v.Name)
	//	} else {
	//		tmpStr = fmt.Sprintf("%s DESC", v.Name)
	//	}
	//	db = db.Order(tmpStr)
	//}
	//return db
}

//GenerateSql 生成sql语句
func GenerateSql(tableName string, qs []*model.QueryStage, fields ...model.OueryField) (string, error) {
	if len(qs) <= 0 || tableName == "" {
		return "", internal.ErrParameter
	}

	//生成查询指定字段语句
	fieldsStr := "*"
	if len(fields) > 0 {
		fArr := []string{}
		for _, v := range fields {
			fArr = append(fArr, v.FName)
		}
		fieldsStr = strings.Join(fArr, ",")
	}

	str := fmt.Sprintf("SELECT %s FROM %s WHERE ", fieldsStr, tableName)

	//循环阶段
	for _, v := range qs {
		//将阶段中的查询条件生成字符串
		tmpStr := queryStageGenerateSql(v.QueryS)
		shellStr := "(%s)"
		if v.IsNot {
			shellStr = fmt.Sprintf("NOT%s", shellStr)
		}

		tmpStr = fmt.Sprintf(shellStr, tmpStr)
		str = fmt.Sprintf("%s %s ", str, tmpStr)

		switch v.Relation {
		case model.AND:
			str += " AND "
		case model.OR:
			str += " OR "
		case model.Nothing:
			str += " "
		default:
		}
	}

	//循环，判断
	//for _, v := range qs {
	//
	//	if v.Not {
	//		str += " NOT "
	//	}
	//
	//	switch v.Condition {
	//	case Equal:
	//		str += fmt.Sprintf(" %s = '%s' ", v.Key, v.Value)
	//	case Like:
	//		str += fmt.Sprintf(" %s LIKE '%s' ", v.Key, v.Value)
	//	case Greater:
	//		str += fmt.Sprintf(" %s > '%s' ", v.Key, v.Value)
	//	case GreaterEqual:
	//		str += fmt.Sprintf(" %s >= '%s' ", v.Key, v.Value)
	//	case Less:
	//		str += fmt.Sprintf(" %s < '%s' ", v.Key, v.Value)
	//	case LessEqual:
	//		str += fmt.Sprintf(" %s <= '%s' ", v.Key, v.Value)
	//	case In:
	//		arr := strings.SplitN(v.Value, ";", -1)
	//		inValue := "("
	//		if len(arr) > 0 {
	//			for i, _ := range arr {
	//				inValue += fmt.Sprintf("\"%s\" ", arr[i])
	//				if i == len(arr)-1 {
	//					break
	//				}
	//				inValue += ","
	//			}
	//		}
	//		inValue += ")"
	//		str += fmt.Sprintf(" %s IN %s ", v.Key, inValue)
	//	default:
	//	}
	//
	//	switch v.Relation {
	//	case AND:
	//		str += " AND "
	//	case OR:
	//		str += " OR "
	//	case Nothing:
	//		str += " "
	//	}
	//}
	return str, nil
}

func queryStageGenerateSql(qs []*model.Query) string {
	str := ""
	for _, q := range qs {
		tmpStr := ""
		switch q.QueryCondition {
		case model.Equal:
			tmpStr += fmt.Sprintf(" %s = '%s' ", q.Key, q.Value)
		case model.NotEqual:
			tmpStr += fmt.Sprintf(" %s != '%s' ", q.Key, q.Value)
		case model.Like:
			tmpStr += fmt.Sprintf(" %s LIKE '%s' ", q.Key, q.Value)
		case model.Greater:
			tmpStr += fmt.Sprintf(" %s > '%s' ", q.Key, q.Value)
		case model.GreaterEqual:
			tmpStr += fmt.Sprintf(" %s >= '%s' ", q.Key, q.Value)
		case model.Less:
			tmpStr += fmt.Sprintf(" %s < '%s' ", q.Key, q.Value)
		case model.LessEqual:
			tmpStr += fmt.Sprintf(" %s <= '%s' ", q.Key, q.Value)
		case model.In:
			arr := strings.SplitN(q.Value, ";", -1)
			inValue := "("
			if len(arr) > 0 {
				for i, _ := range arr {
					inValue += fmt.Sprintf("\"%s\" ", arr[i])
					if i == len(arr)-1 {
						break
					}
					inValue += ","
				}
			}
			inValue += ")"
			tmpStr += fmt.Sprintf(" %s IN %s ", q.Key, inValue)
		default:
		}

		if q.IsNot {
			tmpStr = fmt.Sprintf(" NOT(%s) ", tmpStr)
		}

		switch q.QueryRelation {
		case model.AND:
			tmpStr += " AND "
		case model.OR:
			tmpStr += " OR "
		case model.Nothing:
			tmpStr += " "
		default:
		}

		str = fmt.Sprintf("%s %s", str, tmpStr)
	}
	return str
}
