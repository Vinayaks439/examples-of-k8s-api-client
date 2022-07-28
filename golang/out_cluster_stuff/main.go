package main

import (
	"os"
	"vinayak/helpers"
)

func main() {
	var projectID = os.Getenv("PROJECT_ID")
	var clusterID = os.Getenv("CLUSTER_ID")
	var zone = os.Getenv("ZONE")
	var namespace = os.Getenv("NAMESPACE")
	con := helpers.Connect(projectID, clusterID, zone)
	helpers.DeletePo(con, projectID, namespace)
}
