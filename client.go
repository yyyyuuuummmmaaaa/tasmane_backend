package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    "log"
    user "packages/proto/pb"
)

func main() {
    conn, err := grpc.Dial("go:3001", grpc.WithInsecure())
    if err != nil {
        log.Fatal("connection error:", err)
    }
    defer conn.Close()

    client := user.NewUserClient(conn)
    message := &user.GetUserRequest{Id: "1"}
    res, err := client.GetUser(context.Background(), message)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("result:%s\n", res)
}