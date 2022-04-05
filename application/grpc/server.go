package grpc

import (
	"fmt"
	"github.com/Uallessonivo/codepix/application/grpc/pb"
	"github.com/Uallessonivo/codepix/application/usecase"
	"github.com/Uallessonivo/codepix/infrastructure/repository"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	pixUseCase := usecase.PixUseCase{PixKeyRepository: pixRepository}
	pixGrpcService := NewPixGrpcService(pixUseCase)

	pb.RegisterPixServiceServer(grpcServer, pixGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)

	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}

	log.Printf("grpc server started on port %d", port)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}
}
