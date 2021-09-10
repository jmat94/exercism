package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func checkRoot(record Record) bool {
	return record.ID == 0 && record.Parent == 0
}

func checkValidValue(record Record) bool {
	return record.ID > record.Parent
}

func createNode(ID int) *Node {
	return &Node{
		ID: ID,
	}
}

func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(value1, value2 int) bool {
		record1 := records[value1]
		record2 := records[value2]

		return record1.ID < record2.ID
	})

	nodes := make([]*Node, len(records))

	root := records[0]
	records = records[1:]

	if !checkRoot(root) {
		return nil, errors.New("no root found")
	}

	nodes[0] = createNode(root.ID)

	for index, record := range records {
		if !checkValidValue(record) {
			return nil, errors.New("invalid value")
		}

		if index+1 != record.ID {
			return nil, errors.New("non-continuous Ids")
		}

		if index > 0 {

			newRecord := records[index-1]
			if newRecord.ID == record.ID && newRecord.Parent == record.Parent {
				return nil, errors.New("duplicates found")
			}

		}

		nodeParent := nodes[record.Parent]
		node := createNode(record.ID)
		nodes[node.ID] = node
		nodeParent.Children = append(nodeParent.Children, node)

	}

	return nodes[0], nil

}
