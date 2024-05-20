package main

import (
	"fmt"

	"github.com/kpvdr/cdc_mort/mort"
)

func main() {
	//fmt.Println("starting")
	//var cod_39 = mort.NewKeyMap("data/cause_of_death_39_rev10.txt", 4)
	//cod_39.Print()
	//var cod_113 = mort.NewKeyMap("data/cause_of_death_113_rev10.txt", 4)
	//cod_113.Print()
	//var cod_130 = mort.NewKeyMap("data/cause_of_death_130_rev10.txt", 4)
	//cod_130.Print()
	//var cod_358 = mort.NewKeyMap("data/cause_of_death_358_rev10.txt", 4)
	//cod_358.Print()
	//var icd10 = mort.NewKeyMap("data/icd10cm-codes-2022.txt", 8)
	//icd10.Print()
	var mfmd = mort.NewMortFieldMetaData("data/mort_mucod_data_layout_2022.json")
	mfmd.Print()
	fmt.Println("done")
}
