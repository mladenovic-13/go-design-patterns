package abstractfactory

type FamilyCar struct{}

func (*FamilyCar) NumDoors() int {
	return 4
}
func (*FamilyCar) NumWheels() int {
	return 4
}
func (*FamilyCar) NumSeats() int {
	return 5
}
