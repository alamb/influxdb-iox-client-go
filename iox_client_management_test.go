package influxdbiox_test

import (
	"context"

	"github.com/influxdata/influxdb-iox-client-go"
)

func ExampleClient_CreateDatabase() {
	config, _ := influxdbiox.ClientConfigFromAddressString("localhost:8082")
	client, _ := influxdbiox.NewClient(context.Background(), config)

	_ = client.CreateDatabase(context.Background(), "mydb")
}

func ExampleClient_ListDatabases() {
	config, _ := influxdbiox.ClientConfigFromAddressString("localhost:8082")
	client, _ := influxdbiox.NewClient(context.Background(), config)

	databases, _ := client.ListDatabases(context.Background())
	for _, database := range databases {
		println(database)
	}
}
