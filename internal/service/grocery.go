package service

import (
	"context"
	"fmt"
	shopapi "kratos_test/api/shop/v1"
	"time"
)

// GroceryService 此函数连结到grpc.go上，因为需要进行gRPC的注册，由于NewGRPCServer支持可变参数，因此在后续继续添加服务模块即可
type GroceryService struct {
	shopapi.UnimplementedGroceryServer
}

func NewGroceryService() *GroceryService {
	return &GroceryService{}
}

// 这里是业务层书写地方

// CreateGrocery FIXME 商品-业务层-创建商品
func (s *GroceryService) CreateGrocery(ctx context.Context, req *shopapi.CreateGroceryRequest) (*shopapi.CreateGroceryReply, error) {
	// 接下来可以自定义修改groceryEntity
	req.Id = "10001"
	timestamp := time.Now().Unix()
	fmt.Printf("created groceryEntity: %s\n", req)
	reply := &shopapi.CreateGroceryReply{
		Status:    "200",
		Msg:       "Create Successfully!",
		Timestamp: timestamp,
		Id:        req.Id,
	}
	return reply, nil
}
func (s *GroceryService) UpdateGrocery(ctx context.Context, req *shopapi.UpdateGroceryRequest) (*shopapi.UpdateGroceryReply, error) {
	return &shopapi.UpdateGroceryReply{}, nil
}
func (s *GroceryService) DeleteGrocery(ctx context.Context, req *shopapi.DeleteGroceryRequest) (*shopapi.DeleteGroceryReply, error) {
	return &shopapi.DeleteGroceryReply{}, nil
}
func (s *GroceryService) GetGrocery(ctx context.Context, req *shopapi.GetGroceryRequest) (*shopapi.GetGroceryReply, error) {
	return &shopapi.GetGroceryReply{}, nil
}
func (s *GroceryService) ListGrocery(ctx context.Context, req *shopapi.ListGroceryRequest) (*shopapi.ListGroceryReply, error) {
	return &shopapi.ListGroceryReply{}, nil
}
