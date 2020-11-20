package main

import "fmt"

func main(){
	var slice1 []int
	slice2 := []string{"a","b","c","d","e"}
	slice3 := make([]string, 3)
	fmt.Println(len(slice1))
	slice1 = append(slice1, 65)
	fmt.Println(slice1)
	fmt.Println(slice3)
	slice2[2] = "Queen"
	slice2 = append(slice2, string(65))
	fmt.Println(slice3)
	fmt.Printf("%x\n", &slice3[2])
	fmt.Println(slice2)
	fmt.Printf("%x\n", &slice2[2])

	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)
	fmt.Println(floatSlice[8], boolSlice[6])
	myMusic := []string{"Brown eyes", "Exo", "AKMU", "BIGBANG"}
	urMusic := myMusic[:]
	urMusic[3] = "Sun"
	fmt.Printf("%x\n", &urMusic[3])
	fmt.Printf("%x\n", &myMusic[3])
	myMusic = append(myMusic, "Irin")
	fmt.Println(urMusic)
	fmt.Println(myMusic)
	fmt.Printf("%x\n", &urMusic[3])
	fmt.Printf("%x\n", &myMusic[3])
}
