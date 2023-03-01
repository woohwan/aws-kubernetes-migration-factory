package source

import (
	cluster "containers-migration-factory/app/cluster"
	resource "containers-migration-factory/app/resource"
)

type Source interface {
	Connect(sCluster *cluster.Cluster)
	GetSourceDetails(sCluster *cluster.Cluster) resource.Resources
	FormatSourceData(resource *resource.Resources, resToInclude []string) 	// trim / data clean up
}

func SetContext(source Source, sCluster *cluster.Cluster) {
	// Connect to source clusters
	source.Connect(sCluster)
}

func Invoke(source Source, sType string, sCluster *cluster.Cluster, dCluster *cluster.Cluster) resource.Resources {
	resources := source.GetSourceDetails(sCluster)
	source.FormatSourceData(&resources, sCluster.Resources)

	return resources
}