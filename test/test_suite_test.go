package test

import (
	"context"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"latentlab.cc/easyhousing/api"
	"latentlab.cc/easyhousing/pkg/service"
	"log"
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	ctx    context.Context
	Client api.EasyHousingClient
)

const (
	IntegrationTestLabel = "integration"
)

func TestTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite")
}

var _ = BeforeSuite(func() {
	ctx = context.Background()
	err := godotenv.Load()
	if err != nil {
		Fail(err.Error())
	}

	go func() {
		service.Start()
	}()

	grpcEndpoint := os.Getenv("GRPC_HOST")
	if grpcEndpoint == "" {
		log.Println("No GRPC_HOST environment variable found. Defaulting to localhost:8889")
		grpcEndpoint = "localhost:8889"
	}

	conn, err := grpc.NewClient(grpcEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	Expect(err).ToNot(HaveOccurred())
	Client = api.NewEasyHousingClient(conn)
})
