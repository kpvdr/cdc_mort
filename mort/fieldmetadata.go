package mort

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

type MortFieldMetaDataer interface {
	LoadFieldMetaData(jsonfile string)
	field(string) map[string]interface{}
	Print()
}

type MortFieldMetaData struct {
	name     string
	year     int
	keyList  []string
	fieldmap map[string]interface{}
}

func NewMortFieldMetaData(jsonfilename string) *MortFieldMetaData {
	mfmd := new(MortFieldMetaData)
	mfmd.LoadFieldMetaData(jsonfilename)
	return mfmd
}

func (mfmd *MortFieldMetaData) LoadFieldMetaData(jsonfilename string) {
	jsonfile, err := os.Open(jsonfilename)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonfile.Close()

	var jsonmap map[string]interface{}
	jsonParser := json.NewDecoder(jsonfile)
	if err = jsonParser.Decode(&jsonmap); err != nil {
		log.Print("parsing config file", err.Error())
	}
	mfmd.name = jsonmap["title"].(string)
	mfmd.year = int(jsonmap["year"].(float64))
	mfmd.fieldmap = jsonmap["fields"].(map[string]interface{})

	// Create list of keys ordered by json "loc" key
	for k := range mfmd.fieldmap {
		mfmd.keyList = append(mfmd.keyList, k)
	}
	sort.Slice(mfmd.keyList,
		func(i, j int) bool {
			i_field := mfmd.field(mfmd.keyList[i])
			j_field := mfmd.field(mfmd.keyList[j])
			return int(i_field["loc"].(float64)) < int(j_field["loc"].(float64))
		})
}

func (mfmd MortFieldMetaData) field(k string) map[string]interface{} {
	val, ok := mfmd.fieldmap[k]
	if !ok {
		log.Fatal(k)
	}
	return val.(map[string]interface{})
}

func (mfmd MortFieldMetaData) Print() {
	fmt.Printf("name: \"%s\"\nyear: %d\nfields: \n", mfmd.name, mfmd.year)
	for i, key := range mfmd.keyList {
		fmt.Printf("  %3d %s\n", i, key)
	}
}
