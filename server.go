package main

import (
    "context"
    "errors"
    "google.golang.org/grpc"
    "log"
    "net"
    user "packages/proto/pb"
)

type userService struct{}

func (s *userService) GetUser(ctx context.Context, message *user.GetUserRequest) (*user.GetUserResponse, error) {
    switch message.Id {
    case "1":
        return &user.GetUserResponse{
            Name: "tama",
            Email: "tama@terra-drone.co.jp",
            Phone: "123456789",
            OrgCode: "TRD",
            LisenceNumber: "141421356",
        }, nil
    default:
        return nil, errors.New("Not Found User..")
    }
}

func main() {
    port, err := net.Listen("tcp", ":3001")
    if err != nil {
        log.Println(err.Error())
        return
    }
    s := grpc.NewServer()
    user.RegisterUserServer(s, &userService{})
    s.Serve(port)
}