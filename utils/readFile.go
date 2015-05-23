package utils

import (
    "io/ioutil"
)

func ReadFile(location string) string{
    data, err := ioutil.ReadFile(location)
	GeneralCheckError(err)
	return string(data)
}