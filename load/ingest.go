package load

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-kusto-go/kusto"
	"github.com/Azure/azure-kusto-go/kusto/ingest"
)

const blobStorePathFormat = "https://%s.blob.core.windows.net/%s/%s%s"

// CSVFromBlob ingests a CSV file from a Blob store
func CSVFromBlob(kc *kusto.Client, blobStoreAccountName, blobStoreContainer, blobStoreToken, blobStoreFileName, kustoMappingRefName, kustoDB, kustoTable string) {
	kIngest, err := ingest.New(kc, kustoDB, kustoTable)
	if err != nil {
		log.Fatal("failed to create ingestion client", err)
	}
	blobStorePath := fmt.Sprintf(blobStorePathFormat, blobStoreAccountName, blobStoreContainer, blobStoreFileName, blobStoreToken)
	err = kIngest.FromFile(context.Background(), blobStorePath, ingest.FileFormat(ingest.CSV), ingest.IngestionMappingRef(kustoMappingRefName, ingest.CSV))

	if err != nil {
		log.Fatal("failed to ingest file", err)
	}
	log.Println("Ingested file from -", blobStorePath)
}
