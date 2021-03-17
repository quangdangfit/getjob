package utils

import (
	"github.com/quangdangfit/getjob/pkg/logger"
)

func Copy(dest interface{}, src interface{}) error {
	data, err := json.Marshal(src)
	if err != nil {
		logger.Error("Failed to marshal data")
		return err
	}

	return json.Unmarshal(data, dest)
}
