package cli

import (
	"log"

	"github.com/abhirockzz/datax-go/cli/util"
	"github.com/abhirockzz/datax-go/setup"

	"github.com/spf13/cobra"
)

// azdatax list-tables --dbname testdb
var listTablesCLICommand = cobra.Command{Use: "list-tables", Short: "lists tables", Long: "list tables in database passed using --dbname", Run: listTables}

func init() {
	listTablesCLICommand.Flags().String(dbnameFlag, "", "Database name in Azure Data Explorer")
	listTablesCLICommand.MarkFlagRequired(dbnameFlag)
	rootCmd.AddCommand(&listTablesCLICommand)
}

func listTables(c *cobra.Command, args []string) {
	dbname := c.Flag(dbnameFlag).Value.String()
	tables := setup.FindTables(util.Connect(), dbname)
	for _, t := range tables {
		log.Printf("Table name: %s, Database name: %s\n", t.Name, t.DB)
	}
}
