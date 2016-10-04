package main

import "fmt"

func main() {
	fmt.Println("Starts..")

	// res, err := TestClient.PostForm()
	// if err != nil {
	// 	fmt.Println("error..")
	// }
	// fmt.Println(res.StatusCode)

	// res, err := TestClient.Post()
	// if err != nil {
	// 	fmt.Println("error..")
	// }
	// fmt.Println(res.StatusCode)

	// res, err := TestClient.PostFileAndData()
	// if err != nil {
	// 	fmt.Println("error..")
	// }
	// fmt.Println(res.StatusCode)

	res, err := TestClient.PostJSONData()
	if err != nil {
		fmt.Println("error..")
	}
	fmt.Println(res.StatusCode)

}
