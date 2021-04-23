/*
@Time : 2021/4/23 上午11:23
@Author : MuYiMing
@File : definition
@Software: GoLand
*/
package model

type ValueType int32

const (
	Int    ValueType = 0
	Float  ValueType = 1
	String ValueType = 2
)

//ConditionType 操作符
type ConditionType int32

const (
	Equal        ConditionType = iota    //等于
	NotEqual                             //不等于
	Greater      						 //大于
	GreaterEqual 						 //大于等于
	Less         						 //小于
	LessEqual      						 //小于等于
	Like         						 //通配符
	In              					 //in查询
)

//RelationType 表示当前语句与后语句的关系
type RelationType int32

const (
	AND     RelationType = iota
	OR
	Nothing
)
