package services

import (
	"blog/app/http/dao"
	"blog/app/http/requests"
	"blog/pkg/enum"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

func (s *Service) TestService(cxt *gin.Context, req *requests.TestReq) *response.JsonResponse {
	testData, err := dao.IDao.GetTestById(22)
	if err != nil {
		return response.Fail(enum.HttpFail, err.Error())
	}
	if testData == nil {
		return response.Fail(enum.InvalidArgument, "数据不存在！")
	}

	data := &requests.TestResp{Name: testData.Name}
	return response.Success("", data)
}
