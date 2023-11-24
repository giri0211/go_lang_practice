package methods

import "fmt"

type rect struct {
	width, height int
}

func (rec *rect) area() int {
	return rec.height * rec.width
}

func (rec rect) perimeter() int {
	return 2*rec.width + 2*rec.height
}

func (rec rect) perimeter_n_times(n int) int {
	return (2*rec.width + 2*rec.height) * n
}

func Test_methods() {
	r := rect{width: 10, height: 5}
	fmt.Printf("area %v, perimeter %v\n, perimeter n times %v \n", r.area(), r.perimeter(), r.perimeter_n_times(2))
	rp := &r
	fmt.Printf("area %v, perimeter %v\n perimeter n times %v \n", rp.area(), rp.perimeter(), r.perimeter_n_times(2))
}
