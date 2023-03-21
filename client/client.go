package client

import (
	"context"
	"log"

	pb "github.com/berkantay/user-management-client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	client pb.UserAPIClient
}

func NewClient(ctx context.Context) *Client {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewUserAPIClient(conn)

	return &Client{
		client: client,
	}
}

func (c *Client) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	r, err := c.client.Create(context.Background(), req)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return r, nil
}

func (c *Client) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {

	r, err := c.client.Update(context.Background(), req)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return r, nil
}

func (c *Client) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {

	r, err := c.client.Delete(context.Background(), req)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return r, nil
}

func (c *Client) QueryUser(ctx context.Context, req *pb.QueryUsersRequest) (*pb.QueryUsersResponse, error) {

	r, err := c.client.Query(context.Background(), req)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return r, nil
}

func (c *Client) HealthCheck(ctx context.Context, req *pb.HealthcheckRequest) (*pb.HealthcheckResponse, error) {
	r, err := c.client.HealthCheck(context.Background(), req)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return r, nil

}
