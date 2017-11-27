package binarysearchtree

// API:
//
type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

func Bst(data int) *SearchTreeData {
	return &SearchTreeData{nil, data, nil}
}

func (bst *SearchTreeData) Insert(data int) {
	if bst.data >= data {
		if bst.left != nil {
			bst.left.Insert(data)
		} else {
			bst.left = Bst(data)
		}
	} else {
		if bst.right != nil {
			bst.right.Insert(data)
		} else {
			bst.right = Bst(data)
		}
	}
}

func (bst *SearchTreeData) MapString(mapFunc func(int) string) []string {
	elements := bst.Walk()

	result := make([]string, len(elements))
	for i, elem := range elements {
		result[i] = mapFunc(elem)
	}
	return result
}

func (bst *SearchTreeData) MapInt(mapFunc func(int) int) []int {
	elements := bst.Walk()

	result := make([]int, len(elements))
	for i, elem := range elements {
		result[i] = mapFunc(elem)
	}
	return result
}

func (bst *SearchTreeData) Walk() []int {
	result := []int{}
	if bst.left != nil {
		result = append(result, bst.left.Walk()...)
	}
	result = append(result, bst.data)
	if bst.right != nil {
		result = append(result, bst.right.Walk()...)
	}
	return result
}
