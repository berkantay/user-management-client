package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/berkantay/user-management-client/proto"

	"github.com/berkantay/user-management-client/client"
)

func main() {

	client := client.NewClient(context.Background())

	///*----------------------Service Healthcheck Start----------------------*/
	healthcheck, err := client.HealthCheck(context.Background(), &pb.HealthcheckRequest{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Healthcheck response ->", healthcheck)
	///*----------------------Service Healthcheck End----------------------*/

	///*----------------------Create User Start----------------------*/
	resp, err := client.CreateUser(context.Background(), &pb.CreateUserRequest{
		FirstName: "berkant",
		LastName:  "ay",
		NickName:  "berkantay",
		Password:  "samplepasswd123",
		Email:     "berkantay.5@gmail.com",
		Country:   "Turkey",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Create response ->", resp)
	///*----------------------Create User End----------------------*/

	///*----------------------Delete User Start----------------------*/
	// deleteResponse, err := client.DeleteUser(context.Background(), &pb.DeleteUserRequest{
	// 	Id: resp.Payload.Id,
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Delete response ->", deleteResponse)
	///*----------------------Delete User End----------------------*/

	///*----------------------Update User Start----------------------*/
	// updateResponse, err := client.UpdateUser(context.Background(), &pb.UpdateUserRequest{
	// 	Id:        "bdbde9ad-84cb-43f4-8e9d-d559e26934cb",
	// 	FirstName: "jackson",
	// 	LastName:  "seiter",
	// 	NickName:  "jseiter23",
	// 	Email:     "jseiter@gmail.com",
	// 	Password:  "supersecretpasswd",
	// 	Country:   "Germany",
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Update response ->", updateResponse)
	///*----------------------Update User End----------------------*/

	///*----------------------Query User Start----------------------*/
	queriedUser, err := client.QueryUser(context.Background(), &pb.QueryUsersRequest{
		FirstName: stringPtr("berkant"),
		Page:      int64Ptr(1),
		Size:      int64Ptr(1),
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Query response ->", queriedUser)
	///*----------------------Query User End----------------------*/
}

func stringPtr(s string) *string {
	return &s
}

func int64Ptr(n int64) *int64 {
	return &n
}
