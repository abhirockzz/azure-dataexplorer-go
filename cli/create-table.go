package cli

import (
	"github.com/abhirockzz/datax-go/cli/util"
	"github.com/abhirockzz/datax-go/setup"

	"github.com/spf13/cobra"
)

// azdatax create-table --dbname testdb
var createTableCLICommand = cobra.Command{Use: "create-table", Short: "creates a table named StormEvents", Long: "creates a table named StormEvents in database passed using --dbname", Run: createTable}

func init() {
	createTableCLICommand.Flags().String(dbnameFlag, "", "Database name in Azure Data Explorer")
	createTableCLICommand.MarkFlagRequired(dbnameFlag)
	rootCmd.AddCommand(&createTableCLICommand)
}

func createTable(c *cobra.Command, args []string) {
	dbname := c.Flag(dbnameFlag).Value.String()
	setup.CreateTable(util.Connect(), dbname)
}
