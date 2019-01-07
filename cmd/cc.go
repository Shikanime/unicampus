package main

import (
	"context"
	"fmt"

	"github.com/Shikanime/unicampus/pkg/unicampus_admission"

	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	b := unicampus_admission.NewAdmissionServiceClient(conn)
	arts, _ := b.GetSchool(context.Background(), &unicampus_admission.School{Uuid: "yo"})
	fmt.Println(arts)
}
