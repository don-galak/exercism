package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID, Parent int
	Children   []*Node
}

func Build(records []Record) (*Node, error) {
	sort.SliceStable(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make(map[int]*Node, len(records))

	for i, record := range records {
		if i != record.ID || record.Parent > record.ID || (record.Parent == record.ID && record.ID > 0) {
			return nil, errors.New("malformed records")
		}

		nodes[record.ID] = &Node{ID: record.ID, Parent: record.Parent}
		if record.ID > 0 {
			nodes[record.Parent].Children = append(nodes[record.Parent].Children, nodes[record.ID])
		}
	}

	return nodes[0], nil
}
