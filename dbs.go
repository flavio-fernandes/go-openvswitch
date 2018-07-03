package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/digitalocean/go-openvswitch/ovsdb"
)

// This example demonstrates basic usage of a Client.  The Client connects to
// ovsdb-server and requests a list of all databases known to the server.
func main() {
	// Dial an OVSDB connection and create a *ovsdb.Client.
	c, err := ovsdb.Dial("unix", "/var/run/openvswitch/db.sock")
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	// Be sure to close the connection!
	defer c.Close()

	// Ask ovsdb-server for all of its databases, but only allow the RPC
	// a limited amount of time to complete before timing out.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	dbs, err := c.ListDatabases(ctx)
	if err != nil {
		log.Fatalf("failed to list databases: %v", err)
	}

	for _, d := range dbs {
		fmt.Println(d)
	}
}

