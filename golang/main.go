package main

import (
	"os"
	"vinayak/helpers"
)

func main() {
	projectID := os.Getenv("PROJECT_ID")
	clusterID := os.Getenv("CLUSTER_ID")
	zone := os.Getenv("ZONE")
	helpers.Connect(projectID, clusterID, zone)
}
