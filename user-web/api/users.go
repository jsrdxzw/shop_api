package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"shop_api/user-web/global/response"
	"shop_api/user-web/proto"
	"time"
)

func GetUserList(ctx *gin.Context) {
	ip := "127.0.0.1"
	port := 50051
	// 连接用户grpc服务
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接【用户服务】失败", "msg", err.Error())
	}
	client := proto.NewUserClient(userConn)
	resp, err := client.GetUserList(context.Background(), &proto.PageInfo{})
	if err != nil {
		zap.S().Errorw("[GetUserList] 查询【用户列表】失败", "msg", err.Error())
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	result := make([]response.UserResponse, 0)
	for _, value := range resp.Data {
		userResponse := response.UserResponse{
			Id:       value.Id,
			NickName: value.NickName,
			Birthday: response.JsonTime(time.Unix(int64(value.Birthday), 0)),
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		}
		result = append(result, userResponse)
	}
	ctx.JSON(http.StatusOK, result)
}

// HandleGrpcErrorToHttp 将grpc的code转化为http状码
func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "服务器内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "其他错误",
				})
			}
		}
	}
}
