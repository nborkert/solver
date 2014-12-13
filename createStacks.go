package solver

type PlayerStack struct {
	PlayerNames     string
	Team            string
	ProjectedPoints float64
	Value           float64
	Salary          int
}

type ByVal []PlayerStack

func (this ByVal) Len() int {
	return len(this)
}

func (this ByVal) Less(i, j int) bool {
	return this[i].Value > this[j].Value
}

func (this ByVal) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type ByPoints []PlayerStack

func (this ByPoints) Len() int {
	return len(this)
}

func (this ByPoints) Less(i, j int) bool {
	return this[i].ProjectedPoints > this[j].ProjectedPoints
}

func (this ByPoints) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
