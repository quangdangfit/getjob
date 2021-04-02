package utils

import (
	"encoding/json"

	"github.com/quangdangfit/getjob/pkg/errors"
)

func Copy(dest interface{}, src interface{}) error {
	data, err := json.Marshal(src)
	if err != nil {
		return errors.New(errors.Marshal, err.Error())
	}

	err = json.Unmarshal(data, dest)
	if err != nil {
		return errors.New(errors.Marshal, err.Error())
	}

	return nil
}
