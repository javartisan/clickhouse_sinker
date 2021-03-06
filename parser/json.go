package parser

import (
	"encoding/json"

	"github.com/housepower/clickhouse_sinker/model"
)

type JsonParser struct {
}

func (c *JsonParser) Parse(bs []byte) model.LogKV {
	v := make(map[string]interface{})
	json.Unmarshal(bs, &v)
	return v
}
