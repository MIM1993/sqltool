/*
@Time : 2021/4/23 下午12:17
@Author : MuYiMing
@File : model
@Software: GoLand
*/
package model

type QueryStage struct {
	QueryS []*Query
	IsNot  bool
	//Condition ConditionType
	Relation RelationType
}

type Query struct {
	Key   string
	Value string
	//ValueType ValueType
	QueryCondition ConditionType
	QueryRelation  RelationType
	IsNot          bool
	//Child     []*Query
}

type OueryField struct {
	FName string
	FType ValueType
}

// order>=0 asc    小 --》 大
// order<0  desc   大 --》 小
type OrederBy struct {
	Name  string
	Order int
}

//NewQueryStage create QueryStage
func NewQueryStage(relation RelationType, isNot bool) *QueryStage {
	return &QueryStage{
		QueryS:   make([]*Query, 0),
		IsNot:    isNot,
		Relation: RelationType(relation),
	}
}

//AddQuery add Query into QueryStage.QueryS , In order to add
func (qs *QueryStage) AddQuery(key, value string, condition ConditionType, relation RelationType, isNot bool) {
	q := newQuery(key, value, condition, relation, isNot)
	qs.QueryS = append(qs.QueryS, q)
}

//NewQuery create Query
func newQuery(key, value string, condition ConditionType, relation RelationType, isNot bool) *Query {
	return &Query{
		Key:            key,
		Value:          value,
		QueryCondition: ConditionType(condition),
		QueryRelation:  RelationType(relation),
		IsNot:          isNot,
	}
}
