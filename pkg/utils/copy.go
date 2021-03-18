package utils

import (
	"encoding/json"

	"github.com/quangdangfit/gocommon/logger"

	"github.com/quangdangfit/getjob/pkg/errors"
)

func Copy(dest interface{}, src interface{}) error {
	data, err := json.Marshal(src)
	if err != nil {
		logger.Error("Failed to marshal data")
		return errors.New(errors.ECMarshal, err.Error())
	}

	err = json.Unmarshal(data, dest)
	if err != nil {
		return errors.New(errors.ECMarshal, err.Error())
	}

	return nil
}
