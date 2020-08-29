package goutils

import (
	"encoding/json"
	"errors"
	"strconv"
)

func Interface2String(val interface{}) (ret string, err error) {
	if val == nil {
		return ret, errors.New("input is nil")
	}

	switch value := val.(type) {
	case string:
		ret = value
	case float64:
		ret = strconv.FormatFloat(value, 'f', -1, 64)
	case json.Number:
		ret = val.(json.Number).String()
	}
	return ret, err
}

func Interface2JsonNum(val interface{}) (ret json.Number, err error) {
	if val == nil {
		return ret, errors.New("input is nil")
	}

	switch value := val.(type) {
	case string:
		ret = json.Number(value)
	case float64:
		f := strconv.FormatFloat(value, 'f', -1, 64)
		ret = json.Number(f)
	case json.Number:
		ret, _ = val.(json.Number)
	}
	return ret, err
}

func Interface2Int(val interface{}) (ret int, err error) {
	if val == nil {
		return ret, errors.New("input is nil")
	}

	switch value := val.(type) {
	case string:
		ret, err = strconv.Atoi(value)
	case float32:
		f := strconv.FormatFloat(float64(value), 'f', -1, 64)
		ret, err = strconv.Atoi(f)
	case float64:
		f := strconv.FormatFloat(value, 'f', -1, 64)
		ret, err = strconv.Atoi(f)
	case int64:
		ret = int(value)
	case json.Number:
		i, err1 := val.(json.Number).Int64()
		if err1 == nil {
			ret = int(i)
		}
	}
	return ret, err
}

func Interface2Int64(val interface{}) (ret int64, err error) {
	if val == nil {
		return ret, errors.New("input is nil")
	}

	switch value := val.(type) {
	case string:
		i, err := strconv.Atoi(value)
		if err == nil {
			ret = int64(i)
		}
	case int:
		ret = int64(value)
	case int8:
		ret = int64(value)
	case int16:
		ret = int64(value)
	case int32:
		ret = int64(value)
	case uint:
		ret = int64(value)
	case uint8:
		ret = int64(value)
	case uint16:
		ret = int64(value)
	case uint32:
		ret = int64(value)
	case uint64:
		ret = int64(value)
	case int64:
		ret = value
	case json.Number:
		ret, err = val.(json.Number).Int64()
	default:
		return ret, errors.New("can parse to int64")
	}
	return ret, err
}

func Interface2Float64(val interface{}) (ret float64, err error) {
	if val == nil {
		return ret, errors.New("input is nil")
	}

	switch value := val.(type) {
	case string:
		ret, err = strconv.ParseFloat(value, 32)
	case int:
		ret = float64(value)
	case int8:
		ret = float64(value)
	case int16:
		ret = float64(value)
	case int32:
		ret = float64(value)
	case uint:
		ret = float64(value)
	case uint8:
		ret = float64(value)
	case uint16:
		ret = float64(value)
	case uint32:
		ret = float64(value)
	case uint64:
		ret = float64(value)
	case float32:
		ret = float64(value)
	case float64:
		ret = value
	case json.Number:
		ret, err = val.(json.Number).Float64()
	default:
		return ret, errors.New("can parse to float64")
	}
	return ret, err
}

func Interface2Bool(val interface{}) (ret bool, err error) {
	if val == nil {
		return ret, errors.New("input is nil")
	}

	switch value := val.(type) {
	case bool:
		ret = value
	}
	return ret, err
}
