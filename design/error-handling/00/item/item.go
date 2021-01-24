package item

type List struct {
	id    int
	items []Item
}

type Item struct {
	index int
	name  string
}
