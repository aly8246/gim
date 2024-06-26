package api

import (
	"fmt"
	"gim/pkg/protocol/pb"
	"testing"

	"google.golang.org/grpc"
)

func getBusinessIntClient() pb.BusinessIntClient {
	conn, err := grpc.Dial("180.76.164.117:8020", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return pb.NewBusinessIntClient(conn)
}

func TestUserIntServer_Auth(t *testing.T) {
	res, err := getBusinessIntClient().Auth(getCtx(), &pb.AuthReq{
		UserId:   1069,
		DeviceId: 2,
		Token:    "0",
	})
	fmt.Println(err)
	fmt.Println(res)
}

func TestUserIntServer_GetUsers(t *testing.T) {
	resp, err := getBusinessIntClient().GetUsers(getCtx(), &pb.GetUsersReq{
		UserIds: map[int64]int32{1: 0, 2: 0, 3: 0},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range resp.Users {
		fmt.Printf("%+-5v  %+v\n", k, v)
	}
}
