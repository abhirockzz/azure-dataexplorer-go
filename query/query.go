package query

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/Azure/azure-kusto-go/kusto"
)

const query = "StormEvents | where EventType == 'Flood' and State == 'WASHINGTON' | sort by DamageProperty desc | project StartTime, EndTime, Source, DamageProperty"

// Get query results from kusto
func Get(kc *kusto.Client, kustoDB string) []StormDetail {
	var events []StormDetail
	ri, err := kc.Query(context.Background(), kustoDB, kusto.NewStmt(query))

	if err != nil {
		log.Fatalf("failed to execute query %s - %s", query, err)
	}
	for {
		row, err := ri.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal("error iterating query results -", err)
			}
		}
		var event StormDetail
		row.ToStruct(&event)
		events = append(events, event)
	}
	return events
}

// StormDetail stores storm details
type StormDetail struct {
	Start  time.Time `kusto:"StartTime"`
	End    time.Time `kusto:"EndTime"`
	From   string    `kusto:"Source"`
	Damage int32     `kusto:"DamageProperty"`
}
