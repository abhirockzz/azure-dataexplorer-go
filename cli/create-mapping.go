package cli

import (
	"github.com/abhirockzz/datax-go/cli/util"
	"github.com/abhirockzz/datax-go/setup"

	"github.com/spf13/cobra"
)

// azdatax create-mapping --dbname testdb
var createMappingCLICommand = cobra.Command{Use: "create-mapping", Short: "creates a mapping named StormEvents_CSV_Mapping", Long: "creates a mapping named StormEvents_CSV_Mapping for table StormEvents in database passed using --dbname", Run: createMapping}

func init() {
	createMappingCLICommand.Flags().String(dbnameFlag, "", "Database name in Azure Data Explorer")
	createMappingCLICommand.MarkFlagRequired(dbnameFlag)
	rootCmd.AddCommand(&createMappingCLICommand)
}

func createMapping(c *cobra.Command, args []string) {
	dbname := c.Flag(dbnameFlag).Value.String()
	setup.CreateMapping(util.Connect(), dbname)
}
