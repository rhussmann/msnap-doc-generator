// Copyright 2017 FoxyUtils ehf. All rights reserved.
package main

import (
	"fmt"
	"log"
	"os"
	"encoding/json"
	"io/ioutil"

	"github.com/unidoc/unioffice/document"
)

const (
	Morgantown = iota
	Westover = iota
	Blacksville = iota
	BrucetonMills = iota
	Burton = iota
	Cassville = iota
	Core = iota
	Desllslow = iota
	Fairmont = iota
	Fairview = iota
	Granville = iota
	Hundred = iota
	Independence = iota
	Maidsville = iota
	Masontown = iota
	Osage = iota
	Pentress = iota
	Pursglove = iota
	Reedsville = iota
	Rivesville = iota
	StarCity = iota
	Wadestown = iota
	Wana = iota
)

const (
	no = iota
	yes = iota
	maybe = iota
	na = iota
)

const (
	wom = iota
	fvr = iota
	retails = iota
	vet = iota
	internet = iota
	member = iota
	event = iota
	mccac = iota
	mchs = iota
	newspaper = iota
	radio = iota
	business = iota
	bannerSign = iota
	donorJar = iota
	other = iota
)

const (
	cat = iota
	dog = iota
)

const (
	female = iota
	male = iota
)

const (
	n = iota
	y = iota
)

const (
	dogWeightNA = iota
	_1to60 = iota
	_61to100 = iota
	_greaterThan100 = iota
)

const (
	none = iota
	crypt = iota
	heat = iota
	preg = iota
)

type Body struct {
	Responses []map[string]interface{}
}

type selectEntryFn func(map[string]interface{}) bool
type selectStringFn func(string) bool

func FindEntry(vs []map[string]interface{}, fn selectEntryFn) map[string]interface{} {
	for _, v := range vs {
		if (fn(v)) {
			return v
		}
	}
	return nil
}

func FindStringIndex(vs []string, fn selectStringFn) int {
	for i, v := range vs {
		if (fn(v)) {
			return i
		}
	}
	return -1
}

func FindById(id int) selectEntryFn {
	return func(e map[string]interface{}) bool { return int(e["index"].(float64)) == id }
}

func StringEquals(s string) selectStringFn {
	return func(other string) bool { return s == other }
}

func GetValue(entry map[string]interface{}) string {
	return entry["value"].(string)
}

const (
	FirstName = "Text2"
	LastName = "Text15"
	Address = "Text4"
	City = "Dropdown1"
	WithinCityLimits = "Dropdown2"
	HowDidYouHear = "Dropdown3"
	Zip = "Text5"
	Phone = "Text3"
	VoucherFirstName = "Text7"
	VoucherLastName = "Text16"
	Pet1Name = "Text8"
	Pet1Breed = "Text9"
	Pet1Color = "Text10"
	Pet1Species = "Dropdown4"
	Pet1Gender = "Dropdown5"
	Pet1Stray = "Dropdown6"
	Pet1OlderThan5 = "Dropdown7"
	Pet1DogWeight = "Dropdown10"
	Pet1Special = "Dropdown9"
	Pet1DistinguishingCharacteristics = "Text11"
	Pet2Name = "Text8"
	Pet2Breed = "Text9"
	Pet2Color = "Text10"
	Pet2DistinguishingCharacteristics = "Text11"
)

var cities = []string{
	"Morgantown",
	"Westover",
	"Blacksville",
	"Bruceton Mills",
	"Burton",
	"Cassville",
	"Core",
	"Dellslow",
	"Fairmont",
	"Fairview",
	"Granville",
	"Hundred",
	"Independence",
	"Maidsville",
	"Masontown",
	"Osage",
	"Pentress",
	"Pursglove",
	"Reedsville",
	"Rivesville",
	"Star City",
	"Wadestown",
	"Wana",
}

var withinCityLimitOptions = []string{
	"No",
	"Yes",
	"Maybe",
	"N/A",
}

var species = []string{
	"Cat",
	"Dog",
}

var gender = []string{
	"Female",
	"Male",
}

var stray = []string{
	"No",
	"Yes",
}

var olderThan5Years = []string{
	"No",
	"Yes",
}

var dogWeightRange = []string{
	"n/a",
	"1-60",
	"60-100",
	"101 or more pounds",
}

var special = []string{
	"none",
	"crypt",
	"heat",
	"preg",
}

func main() {
	inputJson := os.Args[1]
	inputDocx := os.Args[2]
	outputDocx := os.Args[3]
	fmt.Printf("Input JSON: %+v\n", inputJson)
	fmt.Printf("Input DOCX: %+v\n", inputDocx)
	fmt.Printf("Output DOCX: %+v\n", outputDocx)

	data, err := ioutil.ReadFile(inputJson)
	entryJson := string(data)
	var result Body
	json.Unmarshal([]byte(entryJson), &result)
	responses := result.Responses

	doc, err := document.Open(inputDocx)
	if err != nil {
		log.Fatalf("error opening form: %s", err)
	}

	// FindAllFields is a helper function that traverses the document
	// identifying fields
	fields := doc.FormFields()
	//	fmt.Println("found", len(fields), "fields")

	for _, fld := range fields {
		//		fmt.Println("- Name:", fld.Name(), "Type:", fld.Type(), "Value:", fld.Value())

		switch fld.Name() {
		case FirstName:
			fld.SetValue(GetValue(FindEntry(responses, FindById(0))))
		case LastName:
			fld.SetValue(GetValue(FindEntry(responses, FindById(1))))
		case Address:
			fld.SetValue(GetValue(FindEntry(responses, FindById(2))))
		case Zip:
			fld.SetValue(GetValue(FindEntry(responses, FindById(3))))
		case City:
			var selectedCity = GetValue(FindEntry(responses, FindById(4)))
			var index = FindStringIndex(cities, StringEquals(selectedCity))
			fld.SetValue(fld.PossibleValues()[index])
		case WithinCityLimits:
			var selectedOption = GetValue(FindEntry(responses, FindById(5)))
			var index = FindStringIndex(withinCityLimitOptions, StringEquals(selectedOption))
			fld.SetValue(fld.PossibleValues()[index])
		case Phone:
			fld.SetValue(GetValue(FindEntry(responses, FindById(6))))
		case VoucherFirstName:
			fld.SetValue(GetValue(FindEntry(responses, FindById(7))))
		case VoucherLastName:
			fld.SetValue(GetValue(FindEntry(responses, FindById(8))))
		case Pet1Name:
			fld.SetValue(GetValue(FindEntry(responses, FindById(10))))
		case Pet1Breed:
			fld.SetValue(GetValue(FindEntry(responses, FindById(13))))
		case Pet1Color:
			fld.SetValue(GetValue(FindEntry(responses, FindById(14))))
		case Pet1DistinguishingCharacteristics:
			fld.SetValue(GetValue(FindEntry(responses, FindById(15))))
		case Pet1Species:
			var selectedOption = GetValue(FindEntry(responses, FindById(11)))
			var index = FindStringIndex(species, StringEquals(selectedOption))
			fld.SetValue(fld.PossibleValues()[index])
		case Pet1Gender:
			var selectedOption = GetValue(FindEntry(responses, FindById(12)))
			var index = FindStringIndex(gender, StringEquals(selectedOption))
			fld.SetValue(fld.PossibleValues()[index])
		case Pet1Stray:
			var selectedOption = GetValue(FindEntry(responses, FindById(16)))
			var index = FindStringIndex(stray, StringEquals(selectedOption))
			fld.SetValue(fld.PossibleValues()[index])
		case Pet1OlderThan5:
			var selectedOption = GetValue(FindEntry(responses, FindById(17)))
			var index = FindStringIndex(olderThan5Years, StringEquals(selectedOption))
			fld.SetValue(fld.PossibleValues()[index])
		case Pet1DogWeight:
			var selectedOption = GetValue(FindEntry(responses, FindById(18)))
			var index = FindStringIndex(dogWeightRange, StringEquals(selectedOption))
			fld.SetValue(fld.PossibleValues()[index])
		case Pet1Special:
			var selectedOption = GetValue(FindEntry(responses, FindById(19)))
			var index = FindStringIndex(special, StringEquals(selectedOption))
			fld.SetValue(fld.PossibleValues()[index])
		}
		// switch fld.Type() {
		// case document.FormFieldTypeText:
		// 	// you can directly set values on text fields
		// 	fld.SetValue("testing 123")
		// case document.FormFieldTypeCheckBox:
		// 	// you can check check boxes
		// 	fld.SetChecked(true)
		// case document.FormFieldTypeDropDown:
		// 	// and select items in a dropdown, here the value must be one of the
		// 	// fields possible values
		// 	lpv := len(fld.PossibleValues())
		// 	if lpv > 0 {
		// 		fld.SetValue(fld.PossibleValues()[lpv-1])
		// 	}
		// }
	}

	doc.SaveToFile(outputDocx)
}
