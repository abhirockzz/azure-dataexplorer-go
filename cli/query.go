package cli

import (
	"log"
	"os"
	"strconv"

	"github.com/abhirockzz/datax-go/cli/util"
	"github.com/abhirockzz/datax-go/query"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// azdatax get --dbname testdb
var queryCLICommand = cobra.Command{Use: "get", Short: "queries data from StormEvents", Long: "queries data from StormEvents table in database specified using --dbname", Run: get}

func init() {
	queryCLICommand.Flags().String(dbnameFlag, "", "Database name in Azure Data Explorer")
	queryCLICommand.MarkFlagRequired(dbnameFlag)
	rootCmd.AddCommand(&queryCLICommand)
}

func get(c *cobra.Command, args []string) {
	dbname := c.Flag(dbnameFlag).Value.String()
	details := query.Get(util.Connect(), dbname)

	data := [][]string{}

	for _, detail := range details {
		data = append(data, []string{detail.Start.String(), detail.End.String(), detail.From, strconv.Itoa(int(detail.Damage))})
	}
	log.Println("StormEvents data....")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Start Time", "End Time", "From", "Damage"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
