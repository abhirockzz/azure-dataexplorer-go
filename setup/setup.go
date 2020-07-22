package setup

import (
	"context"
	"io"
	"log"

	"github.com/Azure/azure-kusto-go/kusto"
)

const (
	kustoTable          = "StormEvents"
	testQuery           = ".show tables"
	kustoMappingRefName = "StormEvents_CSV_Mapping"

	createTableCommand = ".create table StormEvents (StartTime: datetime, EndTime: datetime, EpisodeId: int, EventId: int, State: string, EventType: string, InjuriesDirect: int, InjuriesIndirect: int, DeathsDirect: int, DeathsIndirect: int, DamageProperty: int, DamageCrops: int, Source: string, BeginLocation: string, EndLocation: string, BeginLat: real, BeginLon: real, EndLat: real, EndLon: real, EpisodeNarrative: string, EventNarrative: string, StormSummary: dynamic)"

	createMappingCommand = `.create table StormEvents ingestion csv mapping 'StormEvents_CSV_Mapping' '[{"Name":"StartTime","datatype":"datetime","Ordinal":0}, {"Name":"EndTime","datatype":"datetime","Ordinal":1},{"Name":"EpisodeId","datatype":"int","Ordinal":2},{"Name":"EventId","datatype":"int","Ordinal":3},{"Name":"State","datatype":"string","Ordinal":4},{"Name":"EventType","datatype":"string","Ordinal":5},{"Name":"InjuriesDirect","datatype":"int","Ordinal":6},{"Name":"InjuriesIndirect","datatype":"int","Ordinal":7},{"Name":"DeathsDirect","datatype":"int","Ordinal":8},{"Name":"DeathsIndirect","datatype":"int","Ordinal":9},{"Name":"DamageProperty","datatype":"int","Ordinal":10},{"Name":"DamageCrops","datatype":"int","Ordinal":11},{"Name":"Source","datatype":"string","Ordinal":12},{"Name":"BeginLocation","datatype":"string","Ordinal":13},{"Name":"EndLocation","datatype":"string","Ordinal":14},{"Name":"BeginLat","datatype":"real","Ordinal":16},{"Name":"BeginLon","datatype":"real","Ordinal":17},{"Name":"EndLat","datatype":"real","Ordinal":18},{"Name":"EndLon","datatype":"real","Ordinal":19},{"Name":"EpisodeNarrative","datatype":"string","Ordinal":20},{"Name":"EventNarrative","datatype":"string","Ordinal":21},{"Name":"StormSummary","datatype":"dynamic","Ordinal":22}]'`
)

// CreateTable creates a Kusto table in database
func CreateTable(kc *kusto.Client, kustoDB string) {
	_, err := kc.Mgmt(context.Background(), kustoDB, kusto.NewStmt(createTableCommand))
	if err != nil {
		log.Fatal("failed to create table", err)
	}

	log.Printf("table %s created\n", kustoTable)
}

// FindTables gets tables and their respective database name
func FindTables(kc *kusto.Client, kustoDB string) []TableInfo {
	var tables []TableInfo
	ri, err := kc.Mgmt(context.Background(), kustoDB, kusto.NewStmt(testQuery))
	if err != nil {
		log.Fatalf("failed to execute query %s - %s", testQuery, err)
	}
	var t TableInfo
	for {
		row, err := ri.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal("error iterating query results -", err)
			}
		}
		row.ToStruct(&t)
		tables = append(tables, t)
	}
	return tables
}

// TableInfo stores table and db name
type TableInfo struct {
	Name string `kusto:"TableName"`
	DB   string `kusto:"DatabaseName"`
}

// CreateMapping creates mapping
func CreateMapping(kc *kusto.Client, kustoDB string) {
	_, err := kc.Mgmt(context.Background(), kustoDB, kusto.NewStmt(createMappingCommand))
	if err != nil {
		log.Fatal("failed to create mapping - ", err)
	}
	log.Printf("mapping %s created\n", kustoMappingRefName)
}

const dropTableQ = ".drop table StormEvents"

// DropTable drops a table
func DropTable(kc *kusto.Client, kustoDB string) {
	_, err := kc.Mgmt(context.Background(), kustoDB, kusto.NewStmt(dropTableQ))
	if err != nil {
		log.Fatal("failed to drop table - ", err)
	}
	log.Println("Table StormEvents dropped")
}
