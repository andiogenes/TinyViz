package input

import (
	"encoding/json"
	"io/ioutil"
)

// ArrangementLoader ...
type ArrangementLoader func(string) (interface{}, error)

type coordRepresentation struct {
	Coords [][]int `json:"coordinates"`
}

// CoordinatesLoader ...
func CoordinatesLoader(fileName string) (interface{}, error) {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	// Unmarshall data
	var cRep coordRepresentation
	if err = json.Unmarshal(f, &cRep); err != nil {
		return nil, err
	}

	return cRep.Coords, nil
}
