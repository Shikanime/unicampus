package main

import (
	"context"
	"fmt"

	"github.com/Shikanime/unicampus/api/admission/v1alpha1"

	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	b := unicampus_api_admission_v1alpha1.NewAdmissionServiceClient(conn)
	arts, _ := b.GetSchool(context.Background(), &unicampus_api_admission_v1alpha1.School{Uuid: "yo"})
	fmt.Println(arts)
}
