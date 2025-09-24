package main

import (
	"fmt"
	"math/rand"
	"time"	
)

func echanger(T []int, i, j int) {
	T[i], T[j] = T[j], T[i]
}

func partitionner(T []int, premier, dernier, pivot int) int {
	echanger(T, pivot, dernier)
	j := premier

	for i := premier; i < dernier; i++ {
		if T[i] <= T[dernier] {
			echanger(T, i, j)
			j++
		}
	}
	echanger(T, j, dernier)
	return j
}

func triRapide(T []int, premier, dernier int) {
	if premier < dernier {
		pivot := (premier + dernier) / 2 
		pivot = partitionner(T, premier, dernier, pivot)
		triRapide(T, premier, pivot-1)
		triRapide(T, pivot+1, dernier)
	}
}

func genererTableauUnique(n int) []int {
	T := make([]int, n)
	for i := 0; i < n; i++ {
		T[i] = i
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(n, func(i, j int) {
		T[i], T[j] = T[j], T[i]
	})

	return T
}

func main() {
	n := 70000000
	T := genererTableauUnique(n)


	//fmt.Println("Avant tri :", T)

	// Démarrage du timer
	start := time.Now()

	triRapide(T, 0, len(T)-1)

	// Mesure du temps écoulé
	elapsed := time.Since(start)

	//fmt.Println("Après tri :", T)
	fmt.Println("Temps d’exécution :", elapsed)
}