package awesome

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
}

func NewCodec() *Codec {
	return &Codec{}
}

// Serializes a tree to a single string.
func (th *Codec) Serialize(root *TreeNode) string {
	return fmt.Sprintf("[%s]", strings.Join(LevelOrder(root), ","))
}

// Deserializes your encoded data to tree.
func (th *Codec) Deserialize(data string) (tree *TreeNode) {
	data = strings.TrimSpace(data)
	data = strings.TrimPrefix(data, "[")
	data = strings.TrimSuffix(data, "]")

	item := strings.Split(data, ",")
	if len(item) == 1 && item[0] == "" {
		return
	}

	// var level = 1
	var rootIndex = 0
	var treeList []*TreeNode
	// 遍历item
	for i, val := range item {
		val = strings.TrimSpace(val)

		node := CreateNode(val)
		if node != nil {
			treeList = append(treeList, node)
		}
		if i == 0 {
			continue
		}
		currentRoot := treeList[rootIndex]
		if currentRoot == nil {
			continue
		}
		// 说明是右节点
		if i%2 == 0 {
			currentRoot.Right = node
			rootIndex++
			// 说明是左节点
		} else {
			currentRoot.Left = node
		}
	}
	if len(treeList) > 0 {
		tree = treeList[0]
	}
	return tree
}

func CreateNode(val string) *TreeNode {
	if val == "null" {
		return nil
	}
	value, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		panic(err)
	}
	return &TreeNode{
		Val: int(value),
	}
}

func LevelOrder(root *TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

	for i := 0; len(queue) > 0; i++ {

		var temp []*TreeNode
		for j := 0; j < len(queue); j++ {
			node := queue[j]

			if node == nil {
				res = append(res, "null")
			} else {
				res = append(res, strconv.Itoa(node.Val))

				if node.Left != nil {
					temp = append(temp, node.Left)
				} else {
					temp = append(temp, nil)
				}

				if node.Right != nil {
					temp = append(temp, node.Right)
				} else {
					temp = append(temp, nil)
				}
			}

		}

		queue = temp
	}

	for res[len(res)-1] == "null" {
		res = res[:len(res)-1]
	}
	return res
}
