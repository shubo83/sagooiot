package system

import (
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type GetJobReq struct {
	g.Meta `path:"/job_log/get" method:"get" summary:"获取任务日志详情" tags:"任务日志"`
	Id     int `json:"id" description:"任务日志ID" v:"required#任务日志ID不能为空"`
}

type GetJobLogListReq struct {
	g.Meta  `path:"/job_log/list" method:"get" summary:"获取任务日志列表" tags:"任务日志"`
	JobName string `json:"jobName" description:"任务名称"`
	Status  string `json:"status" description:"状态（0正常 1暂停）"`
	*common.PaginationReq
}

type GetJobLogListRes struct {
	Data []*model.SysJobLogListRes
	common.PaginationRes
}

type DeleteJobLogReq struct {
	g.Meta `path:"/job_log/delete" method:"delete" summary:"删除任务日志" tags:"任务日志"`
	Ids    []int `json:"ids" description:"任务日志ID列表" v:"required#任务日志ID不能为空"`
}

type DeleteJobLogRes struct {
}

type JobExportReq struct {
	g.Meta  `path:"/job_log/export" method:"get" summary:"导出任务日志" tags:"任务日志"`
	JobName string `json:"jobName" description:"任务名称"`
	Status  string `json:"status" description:"状态（0正常 1暂停）"`
	*common.PaginationReq
}

type JobLogExportRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
