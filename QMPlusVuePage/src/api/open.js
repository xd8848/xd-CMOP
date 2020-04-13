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
export const getLiveStreamList = (data) => {
    return service({
        url: "/livestream/getLiveStreamList",
        method: 'post',
        data
    })
}


// @Tags menu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.GetById true "根据id获取菜单"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getApiById [post]
export const getLiveStreamById = (data) => {
    return service({
        url: "/livestream/getLiveStreamById",
        method: 'post',
        data
    })
}




export const getAllLiveStreams = (data) => {
    return service({
        url: "/livestream/getAllLiveStreams",
        method: 'post',
        data
    })
}
