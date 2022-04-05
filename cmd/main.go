package main

import (
  "log"
  "net"
  "errors"
  "context"
  "google.golang.org/grpc"
  "google.golang.org/grpc/reflection"
  pb "packages/proto/pb"
  postgres "packages/database"
  "gorm.io/gorm"
  "fmt"
)

type userService struct {
  pb.UnimplementedUserServer
  gorm.Model
  Name string `gorm:"<-"`
  Description string `gorm:"<-"`
  Role string `gorm:"<-"`
  Email string `gorm:"<-,unique"`
  Phone string `gorm:"<-"`
  Team string `gorm:"<-"`
}

func NewUser() *userService {
	return &userService{}
}

func (s *userService) RegisterUser(ctx context.Context, message *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {

  fmt.Println(message, "message")

  db := postgres.Connect()
  result := db.Create(&userService{
      Name: message.Name,
      Description: message.Description,
      Role: message.Role,
      Email: message.Email,
      Phone: message.Phone,
      Team: message.Team,
    })

  if result.Error != nil {
    log.Fatalf("Data Creation is failed: %v", result.Error)
  }

  return &pb.RegisterUserResponse{
      Name: message.Name,
      Description: message.Description,
      Role: message.Role,
      Email: message.Email,
      Phone: message.Phone,
      Team: message.Team,
    }, nil
}

func (s *userService) GetUser(ctx context.Context, message *pb.GetUserRequest) (*pb.GetUserResponse, error) {
  switch message.Id {
  case "1":
      return &pb.GetUserResponse{
          Name: "tama",
          Description: "My name is Tama.",
          Role: "Admin",
          Email: "tama@terra-drone.co.jp",
          Phone: "123456789",
          Team: "TerraDrone",
      }, nil
  default:
      return nil, errors.New("Not Found User..")
  }
}

func main() {
  lis, err := net.Listen("tcp", ":3000")
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }

  db := postgres.Connect()
	db.AutoMigrate(NewUser())
  fmt.Println("here-----------------------------")

  grpcServer := grpc.NewServer()
  pb.RegisterUserServer(grpcServer, NewUser())
  reflection.Register(grpcServer)

  if err := grpcServer.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %s", err)
  }
}