package main

import (
	"github.com/smbrave/data-collector/idcard"
	"fmt"
)

func main() {
	info, err := idcard.Parse("500382198612277074")
	if err !=nil {
		panic(err)
	}
	fmt.Printf("%+v\n", info)

	info, err = idcard.Parse("120225198703263918")
	if err !=nil {
		panic(err)
	}
	fmt.Printf("%+v\n", info)


	info, err = idcard.Parse("500101198602198330")
	if err !=nil {
		panic(err)
	}
	fmt.Printf("%+v\n", info)


	info, err = idcard.Parse("510283198301017171")
	if err !=nil {
		panic(err)
	}
	fmt.Printf("%+v\n", info)

	info, err = idcard.Parse("51022619660308332X")
	if err !=nil {
		panic(err)
	}
	fmt.Printf("%+v\n", info)

}
