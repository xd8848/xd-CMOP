import service from '@/utils/request'
// @Tags api
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取用户列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
// {
//  page     int
//	pageSize int
// }
export const getChannelReportList = (data) => {
    return service({
        url: "/channel/getChannelReportList",
        method: 'post',
        data
    })
}

export const getAllChannelReports = (data) => {
    return service({
        url: "/channel/getAllChannelReports",
        method: 'post',
        data
    })
}
