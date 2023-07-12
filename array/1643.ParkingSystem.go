package main

type ParkingSystem struct {
	big, medium, small int
}

// T(O) = 75% S(O) = 93%
func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big, medium, small}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		this.big = this.big - 1
		return this.big >= 0
	case 2:
		this.medium = this.medium - 1
		return this.medium >= 0
	case 3:
		this.small = this.small - 1
		return this.small >= 0
	default:
		return false
	}
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
