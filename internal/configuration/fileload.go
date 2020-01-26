package configuration

import (
	"io/ioutil"
	"os"
)

func loadfile(s string) (d []byte, e error) {
	f, err := os.Open(s)
	if err != nil {
		return nil, err
	}

	d, err = ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return
}
