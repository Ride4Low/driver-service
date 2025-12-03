package service

import (
	"fmt"
	"math/rand"
)

func GetRandomAvatar(index int) string {
	return fmt.Sprintf("https://randomuser.me/api/portraits/lego/%d.jpg", index)
}

func GenerateRandomPlate() string {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	plate := ""
	for i := 0; i < 3; i++ {
		plate += string(letters[rand.Intn(len(letters))])
	}

	return plate
}
