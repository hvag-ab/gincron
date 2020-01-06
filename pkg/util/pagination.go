package util

import (
	"math"
)


func Paginator(nums int64, page int, pageSize int) map[string]interface{} {

	if page ==0{
		page = 1
	}
	if pageSize <1 {
		pageSize = 1
	}
	var
	(
		offset int
		prepage interface{} //前一页地址
		nextpage interface{}  //后一页地址
	)
	//根据nums总数，和prepage每页数量 生成分页总数
	totalpage := int(math.Ceil(float64(nums) / float64(pageSize))) //page总数
	if totalpage==1 {
		page = 1
		nextpage = nil
		prepage = nil
		offset = 0
	}else if totalpage ==0{
		page = 1
		nextpage = nil
		prepage = nil
		offset = 0
	}else if page>=totalpage{
		page = totalpage
		nextpage = nil
		prepage = page -1
		offset = (page-1) * pageSize
	}else if page==1{
		nextpage = page +1
		prepage = nil
		offset = 0
	}else{
		nextpage = page+1
		prepage = page -1
		offset = (page-1) * pageSize
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["count"] = nums
	paginatorMap["totalpages"] = totalpage
	if prepage ==nil{
		paginatorMap["prepage"] = nil
	}else{
		paginatorMap["prepage"] = prepage
	}
	if nextpage ==nil{
		paginatorMap["nextpage"] = nil
	}else{
		paginatorMap["nextpage"] = nextpage
	}
	paginatorMap["offset"] = offset
	paginatorMap["pageSize"] = pageSize
	paginatorMap["current"] = page

	return paginatorMap
}
