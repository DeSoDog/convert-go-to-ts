package types


type DeSoNode struct {
	// Name of the node, displayed to users
	Name string

	// HTTPs URL to the node or app
	URL string

	// DeSo username of the node owner
	Owner string
}