package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/sachinjangam/programs/golang/food_delivery_app/foodsvc/api/proto"
	"github.com/sachinjangam/programs/golang/food_delivery_app/foodsvc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedFoodServiceServer
	Svc *service.Service
}

func NewServer(svc *service.Service) *Server {
	return &Server{Svc: svc}
}

func (s *Server) GetMenuItemSummary(ctx context.Context, req *pb.GetMenuItemRequest) (*pb.MenuItemSummary, error) {
	menu, rName, supplier, err := s.Svc.GetMenuItemSummary(ctx, req.MenuItemId)
	if err != nil {
		return nil, err
	}

	summary := &pb.MenuItemSummary{
		Id:         menu.ID,
		Name:       menu.Name,
		Restaurant: rName,
		Availability: func() string {
			if menu.Available {
				return "Available"
			}
			return "UnAvailable"
		}(),
		Price:          fmt.Sprintf("₹%.2f", menu.Price),
		SupplierStatus: supplier,
	}
	return summary, nil
}

func (s *Server) BulkUpdateAvailability(ctx context.Context, req *pb.BulkUpdateRequest) (*pb.BulkUpdateResponse, error) {
	success, failed := s.Svc.BulkUpdateAvailability(ctx, req.MenuItemIds, req.Available, int(req.Concurrency))
	return &pb.BulkUpdateResponse{
		Success: int32(success),
		Failed:  int32(failed),
	}, nil
}

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	id, err := s.Svc.CreateOrder(ctx, req.MenuItemId, req.CustomerName)
	if err != nil {
		return nil, err
	}
	return &pb.OrderResponse{
		OrderId: id,
		Status:  "pending",
	}, nil
}

func (s *Server) GetPendingOrderSummary(ctx context.Context, _ *emptypb.Empty) (*pb.PendingOrderSummary, error) {
	counts, err := s.Svc.GetPendingOrderSummary(ctx)
	if err != nil {
		return nil, err
	}
	resp := &pb.PendingOrderSummary{}
	for rid, cnt := range counts {
		resp.Items = append(resp.Items, &pb.PendingRestaurant{
			RestaurantId:   rid,
			RestaurantName: fmt.Sprintf("Restaurant %d", rid),
			PendingCount:   int32(cnt),
		})
	}
	return resp, nil
}

func ServeGrpc(addr string, s *Server) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterFoodServiceServer(grpcServer, s)
	log.Printf("grpc server listening on %s ", addr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("serve: %v", err)
	}
}
