package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type FireGun struct {
}

func (f *FireGun) Fire() error {
	if time.Now().Second()%10 == 0 {
		return errors.New("not firing at 10s seconds")
	}

	fmt.Println("SHOT!")

	return nil
}

type LaserGun struct {
	Ammo int
}

func (f *LaserGun) Fire() error {
	if f.Ammo == 0 {
		time.Sleep(500 * time.Millisecond)
		f.Ammo = 5
		return errors.New("need ammo, reloading")
	}

	fmt.Println("SHOT!")
	f.Ammo = f.Ammo - 1

	return nil
}

type Weapon interface {
	Fire() error
}

type Spaceship struct {
	Name   string
	Weapon Weapon // DI - Dependency Injection
}

func (s *Spaceship) Run() {
	for {
		err := s.Weapon.Fire()
		if err != nil {
			fmt.Printf("Damn, captain, weapon is not woring: %s\n", err.Error())
		}

		fmt.Println("Moving...")

		time.Sleep(2 * time.Second)
	}
}

type WeatherClient interface {
	GetWeather() (int, error)
}

type Bot struct {
	wc WeatherClient
}

type GismeteoClient struct {
}

func (g GismeteoClient) GetWeather() (int, error) {
	return rand.Int(), nil
}

type YandexWeather struct {
}

func (y *YandexWeather) GetWeather() (int, error) {
	return rand.Int() / 10 * 100 / 10, nil
}

func main() {
	/*b := &Bot{
		wc: &YandexWeather{},
	}*/

	fire := &FireGun{}
	laser := &LaserGun{Ammo: 5}

	var w Weapon
	if rand.Int()%2 == 0 {
		w = fire
	} else {
		w = laser
	}

	a := Spaceship{
		Name:   "Eagle1000",
		Weapon: w,
	}

	a.Run()
}
