package main

import (
	"os"
	"vinayak/helpers"
)

func main() {
	var projectID = os.Getenv("PROJECT_ID")
	var clusterID = os.Getenv("CLUSTER_ID")
	var zone = os.Getenv("ZONE")
	con := helpers.Connect(projectID, clusterID, zone)
	helpers.Createcontext(con, projectID)
}
