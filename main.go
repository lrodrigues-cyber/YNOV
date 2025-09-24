package main

func main() {
	T := []int{8, 3, 1, 7, 0, 10, 2}
	fmt.Println("Avant tri :", T)

	triRapide(T, 0, len(T)-1)

	fmt.Println("Après tri :", T)
}