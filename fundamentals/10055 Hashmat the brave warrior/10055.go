package main

import (
	"fmt"
)

func main(){
	var hashmat_army, enemy_army uint32
	_, err := fmt.Scan(&hashmat_army, &enemy_army) 
	for err == nil {
		println(enemy_army - hashmat_army)
		_, err = fmt.Scan(&hashmat_army, &enemy_army) 
	}
}