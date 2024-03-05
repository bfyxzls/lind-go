package generictype

import "fmt"

type ACluster struct {
	//...
	Title string
}

type BCluster struct {
	//...
	Age int
}

func (c ClusterClient[C]) PrintMethod() {
	//...
	fmt.Println("PrintMethod:", c.cluster)
}

func (b BCluster) ClusterName() string {
	//...
	return "b.age"
}

func (b ACluster) ClusterName() string {
	//...
	return "a.title"
}

func NewClusterClient[C Cluster](cluster C) *ClusterClient[C] {
	return &ClusterClient[C]{cluster: cluster}
}

type Cluster interface {
	ClusterName() string
}

type ClusterClient[C Cluster] struct {
	cluster C
}

type HttpClient interface {
	PrintMethod()
}
