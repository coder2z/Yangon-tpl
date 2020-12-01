/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 14:40
 */
package R

type PageData struct {
	//当前页码
	PageNo int64
	//每页大小
	PageSize int64
	//一共的页数
	TotalPage int64
	//总条数
	TotalCount int64
	//是否是第一页
	FirstPage bool
	//是否是最后一页
	LastPage bool
	//数据
	List interface{}
}

//总条数  当前页码  每页大小   数据list
func Page(count int64, pageNo int64, pageSize int64, list interface{}) PageData {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	return PageData{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp, List: list}
}
