package day8

import (
	"strings"
)

type TreeNode struct {
	Value string
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(value string) *TreeNode {
	return &TreeNode{Value: value}
}

func ParseData(data string) *TreeNode {
	return ParseDataToMap(data)["AAA"]
}

func ParseDataWithEndA(data string) []*TreeNode {
	mapData := ParseDataToMap(data)
	var endA []*TreeNode
	for key := range mapData {
		if strings.HasSuffix(key, "A") {
			endA = append(endA, mapData[key])
		}
	}
	return endA
}

func ParseDataToMap(data string) map[string]*TreeNode {
	lines := strings.Split(data, "\n")
	mapTreeNodes := make(map[string]*TreeNode)
	for _, line := range lines[2:] {
		if line == "" {
			continue
		}
		val := line[0:3]
		newTreeNode := NewTreeNode(val)
		mapTreeNodes[val] = newTreeNode
	}
	for _, line := range lines[2:] {
		if line == "" {
			continue
		}
		val := line[0:3]
		left := line[7:10]
		right := line[12:15]

		mapTreeNodes[val].Left = mapTreeNodes[left]
		mapTreeNodes[val].Right = mapTreeNodes[right]

	}
	return mapTreeNodes
}

func traverseTree(root *TreeNode, instructions string) int {
	steps := 0
	currentNode := root
	instructionIndex := 0

	for currentNode.Value != "ZZZ" {
		if instructionIndex >= len(instructions) {
			instructionIndex = 0 // Restart instructions if at the end
		}

		if instructions[instructionIndex] == 'R' {
			currentNode = currentNode.Right
		} else {
			currentNode = currentNode.Left
		}

		instructionIndex++
		steps++
	}

	return steps
}

func traverseTreeEndWithZ(root *TreeNode, instructions string) int {
	steps := 0
	currentNode := root
	instructionIndex := 0

	for !strings.HasSuffix(currentNode.Value, "Z") {
		if instructionIndex >= len(instructions) {
			instructionIndex = 0 // Restart instructions if at the end
		}

		if instructions[instructionIndex] == 'R' {
			currentNode = currentNode.Right
		} else {
			currentNode = currentNode.Left
		}

		instructionIndex++
		steps++
	}

	return steps
}

func traverseTreeWithMultipleNode(roots []*TreeNode, instructions string) int {
	mapStepsPerRoot := []int{}

	for _, treeNode := range roots {
		steps := traverseTreeEndWithZ(treeNode, instructions)
		mapStepsPerRoot = append(mapStepsPerRoot, steps)
	}
	return lcmOfList(mapStepsPerRoot)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Function to calculate LCM of two numbers
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// Function to find LCM of a list of numbers
func lcmOfList(arr []int) int {
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result = lcm(result, arr[i])
	}
	return result
}
