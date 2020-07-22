package util

import (
	"log"
	"os"

	"github.com/Azure/azure-kusto-go/kusto"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

// Connect establishes a connection to Data explorer
func Connect() *kusto.Client {
	clientID := os.Getenv("AZURE_SP_CLIENT_ID")
	clientSecret := os.Getenv("AZURE_SP_CLIENT_SECRET")
	tenantID := os.Getenv("AZURE_SP_TENANT_ID")
	kustoEndpoint := os.Getenv("KUSTO_ENDPOINT")

	if clientID == "" || clientSecret == "" || tenantID == "" || kustoEndpoint == "" {
		log.Fatal("missing one or more environment variables: AZURE_SP_CLIENT_ID, AZURE_SP_CLIENT_SECRET, AZURE_SP_TENANT_ID, KUSTO_ENDPOINT")
	}

	auth := kusto.Authorization{Config: auth.NewClientCredentialsConfig(clientID, clientSecret, tenantID)}
	dataxClient, err := kusto.New(kustoEndpoint, auth)
	if err != nil {
		log.Fatal("failed to create kusto client", err)
	}

	log.Println("Connected to Azure Data Explorer")
	return dataxClient
}
