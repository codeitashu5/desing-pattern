package Creational_Patterns

/*
   we can have different type of gun's that can be produced and those guns are feed into an war zone so that they can be used by the solders to fire
*/

type IGun interface {
	AssignBullets(amount int)
	FireBullet() string
	GetCurrentBullets() int
}

type MachineGun struct {
	bullets int
	IGun
}

func (mg MachineGun) FireBullet() string {
	if mg.bullets > 100 {
		mg.bullets -= 100
	} else if mg.bullets > 0 {
		mg.bullets = 0
	} else {
		return "nothing to fire"
	}

	return "soldier fired a bullet from machine gun"
}

func (mg MachineGun) GetCurrentBullets() int {
	return mg.bullets
}

func (mg MachineGun) AssignBullets(amount int) {
	// some value of bullets that can be assigned to an ak47 suppose that's only 100
	const mgCapacity = 10000
	// no of bullets that can be added into the ak
	bulletsCanBeAdded := mgCapacity - mg.bullets
	if amount > bulletsCanBeAdded {
		mg.bullets += bulletsCanBeAdded
	} else {
		mg.bullets = amount
	}
}

type Ak47 struct {
	bullets int
	IGun
}

func (ak Ak47) AssignBullets(amount int) {
	// some value of bullets that can be assigned to an ak47 suppose that's only 100
	const akCapacity = 10000
	// no of bullets that can be added into the ak
	bulletsCanBeAdded := akCapacity - ak.bullets
	if amount > bulletsCanBeAdded {
		ak.bullets += bulletsCanBeAdded
	} else {
		ak.bullets = amount
	}
}

func (ak Ak47) FireBullet() string {
	if ak.bullets > 1 {
		ak.bullets -= 1
		return "fired a bullet at an speed of 1400 km/h"
	}
	return "no bullets left to fire"
}

func (ak Ak47) GetCurrentBullets() int {
	return ak.bullets
}

type Battalion struct {
}
