package util

import (
	"fmt"
	"strconv"
	"strings"
)

type TrieNode struct {
	children     map[string]*TrieNode
	isEndOfRange bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[string]*TrieNode),
	}
}

func Insert(root *TrieNode, ipRange string) {
	ipParts := strings.Split(ipRange, "-")
	startIP := convertToBinary(ipParts[0])
	endIP := convertToBinary(ipParts[1])

	currentNode := root
	for i := 0; i < len(startIP) && i < len(endIP); i++ {
		bit1 := string(startIP[i])
		bit2 := string(endIP[i])
		if bit1 != bit2 {
			break
		}
		if _, ok := currentNode.children[bit1]; !ok {
			currentNode.children[bit1] = NewTrieNode()
		}
		currentNode = currentNode.children[bit1]
	}
	currentNode.isEndOfRange = true

}

func Find(root *TrieNode, ip string) bool {
	binaryIP := convertToBinary(ip)

	currentNode := root
	for _, bit := range binaryIP {
		bitStr := string(bit)
		if childNode, ok := currentNode.children[bitStr]; ok {
			currentNode = childNode
			if currentNode.isEndOfRange {
				fmt.Println("ip block:", bitStr)
				return true
			}
		} else {
			return false
		}
	}
	return false

}

func convertToBinary(ip string) string {
	binary := ""
	ipParts := strings.Split(ip, ".")
	for _, part := range ipParts {
		num, _ := strconv.Atoi(part)
		binary += fmt.Sprintf("%08b", num)
	}
	return binary
}

func ConvertIPToScore(ip string) (int64, error) {
	ipParts := strings.Split(ip, ".")
	if len(ipParts) != 4 {
		return 0, fmt.Errorf("Invalid IP address format")
	}

	var score int64 = 0
	for i := 0; i < 4; i++ {
		partValue, err := strconv.ParseInt(ipParts[i], 10, 64)
		if err != nil {
			fmt.Printf("err partValue:%v", partValue)
			return 0, err
		}
		score = (score << 8) + partValue // 将每个部分的值左移8位并相加
	}
	return score, nil
}
