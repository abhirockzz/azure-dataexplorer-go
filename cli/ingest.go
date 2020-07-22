package cli

import (
	"github.com/abhirockzz/datax-go/cli/util"
	"github.com/abhirockzz/datax-go/load"
	"github.com/spf13/cobra"
)

// azdatax ingest --dbname testdb
var ingestCLICommand = cobra.Command{Use: "ingest", Short: "ingests CSV file from blob store", Long: "ingests data from blob storage container. full path - https://kustosamplefiles.blob.core.windows.net/samplefiles/StormEvents.csv?st=2018-08-31T22%3A02%3A25Z&se=2020-09-01T22%3A02%3A00Z&sp=r&sv=2018-03-28&sr=b&sig=LQIbomcKI8Ooz425hWtjeq6d61uEaq21UVX7YrM61N4%3D", Run: ingest}

const (
	kustoTable          = "StormEvents"
	kustoMappingRefName = "StormEvents_CSV_Mapping"

	blobStoreAccountName = "kustosamplefiles"
	blobStoreContainer   = "samplefiles"
	blobStoreToken       = "?st=2018-08-31T22%3A02%3A25Z&se=2020-09-01T22%3A02%3A00Z&sp=r&sv=2018-03-28&sr=b&sig=LQIbomcKI8Ooz425hWtjeq6d61uEaq21UVX7YrM61N4%3D"
	blobStoreFileName    = "StormEvents.csv"
	blobStorePathFormat  = "https://%s.blob.core.windows.net/%s/%s%s"
)

func init() {
	ingestCLICommand.Flags().String(dbnameFlag, "", "Database name in Azure Data Explorer")
	ingestCLICommand.MarkFlagRequired(dbnameFlag)
	rootCmd.AddCommand(&ingestCLICommand)
}

func ingest(c *cobra.Command, args []string) {
	dbname := c.Flag(dbnameFlag).Value.String()
	load.CSVFromBlob(util.Connect(), blobStoreAccountName, blobStoreContainer, blobStoreToken, blobStoreFileName, kustoMappingRefName, dbname, kustoTable)
}
