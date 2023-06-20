package main

import (
	"fmt"
	"log"
	"net"

	"github.com/skrevolve/auth-svc/pkg/config"
	"github.com/skrevolve/auth-svc/pkg/db"
	"github.com/skrevolve/auth-svc/pkg/pb"
	"github.com/skrevolve/auth-svc/pkg/services"
	"github.com/skrevolve/auth-svc/pkg/utils"
	"google.golang.org/grpc"
)

func main() {

	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config loading", err)
	}

	h := db.Init(c.DBUrl)

	jwt := utils.JwtWrapper{
		SecretKey: c.JWTSecretKey,
		Issuer: "auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listening:", err)
	}
	fmt.Println("Auth Svc on", c.Port)

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &services.Server{
		H: h,
		Jwt: jwt,
	})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}