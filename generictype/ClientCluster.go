package main

import "fmt"

/**
 * 知识点：一个方法，通过传入不同的接口实例，来返回不同的接口类型实例，这种也是多态
 */
// 定义出接口，接口方法是规范
type Cluster interface {
	ClusterName() string
}

// 定义一个类型
type ACluster struct {
	//...
	Title string
}

type BCluster struct {
	//...
	Age int
}

// 类型去实现接口中的方法
func (b BCluster) ClusterName() string {
	//...
	return "b.age"
}

func (b ACluster) ClusterName() string {
	//...
	return "a.title"
}

// ------------------------------------------------------------//
type HttpClient interface {
	PrintMethod()
}

type ClusterClient[C Cluster] struct {
	cluster C
}

func (c ClusterClient[C]) PrintMethod() {
	//...
	fmt.Println("PrintMethod:", c.cluster)
}

// 通过一个方法，传入一个类型实便，再返回另一个类型的实例
func NewClusterClient[C Cluster](cluster C) *ClusterClient[C] {
	return &ClusterClient[C]{cluster: cluster}
}

func main() {
	// 通过传入一个类型实例，返回另一个类型的实例
	clusterClient := NewClusterClient(ACluster{Title: "a.title"})
	clusterClient.PrintMethod()
	clusterClient1 := NewClusterClient(BCluster{Age: 20})
	clusterClient1.PrintMethod()
}
