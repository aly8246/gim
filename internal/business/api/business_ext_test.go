package api

import (
	"context"
	"fmt"
	"gim/pkg/protocol/pb"
	"strconv"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func getBusinessExtClient() pb.BusinessExtClient {
	conn, err := grpc.Dial("180.76.164.117:8020", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return pb.NewBusinessExtClient(conn)
}

func getCtx() context.Context {
	token := "0"
	return metadata.NewOutgoingContext(context.TODO(), metadata.Pairs(
		"user_id", "1",
		"device_id", "1",
		"token", token,
		"request_id", strconv.FormatInt(time.Now().UnixNano(), 10)))
}

func TestUserExtServer_SignIn(t *testing.T) {
	resp, err := getBusinessExtClient().SignIn(getCtx(), &pb.SignInReq{
		PhoneNumber: "15730809230",
		Code:        "0",
		DeviceId:    2,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestUserExtServer_GetUser(t *testing.T) {
	resp, err := getBusinessExtClient().GetUser(getCtx(), &pb.GetUserReq{UserId: 2})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", resp)
}
