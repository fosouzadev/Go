package main

import "fmt"

func main() {
	pessoas := map[string][]string{
		"satoLuiza": {"comer", "dormir"},
	}

	pessoas["souzaFelipe"] = []string{"estudar", "tocar"}
	pessoas["souzaSabrina"] = []string{"fotos", "dormir"}

	for k, v := range pessoas {
		fmt.Println(k, v)
	}

	delete(pessoas, "souzaFelipe")

	for k, v := range pessoas {
		fmt.Println(k, v)
	}
}
