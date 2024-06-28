package main

func main() {
	robot := Robot{true, 2, 4}
	testStruct := &robot
	testStruct.RideBike()
}

type Robot struct {
	On          bool
	Ammo, Power int
}

func (testStruct *Robot) Shoot() bool {
	if testStruct.On == false || testStruct.Ammo == 0 {
		return false
	} else {
		testStruct.Ammo -= 1
		return true
	}
}
func (testStruct *Robot) RideBike() bool {
	if testStruct.On == false || testStruct.Power == 0 {
		return false
	} else {
		testStruct.Power -= 1
		return true
	}
}
