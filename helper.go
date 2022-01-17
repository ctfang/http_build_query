package http_build_query

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func Encode(data map[string]interface{}) string {
	return encodeMap(data)
}

func encodeMap(data map[string]interface{}, keys ...string) string {
	key := ""
	for _, k := range keys {
		key = k
	}

	paramsArr := make([]string, 0, len(data))
	for k, v := range data {
		if key != "" {
			k = "[" + k + "]"
		}

		switch v.(type) {
		case string, uint, int, uint8, int8, uint16, int16, uint32, int32, uint64, int64, float32, float64:
			paramsArr = append(paramsArr, fmt.Sprintf("%s%s=%v", key, k, v))
		case map[string]string:
			for ks, s := range v.(map[string]string) {
				paramsArr = append(paramsArr, fmt.Sprintf("%s%s[%s]=%v", key, k, ks, s))
			}
		case []string:
			for _, s := range v.([]string) {
				paramsArr = append(paramsArr, fmt.Sprintf("%s%s[]=%v", key, k, s))
			}
		case []float64:
			for _, s := range v.([]float64) {
				paramsArr = append(paramsArr, fmt.Sprintf("%s%s[]=%v", key, k, s))
			}
		default:
			bt, _ := json.Marshal(v)
			var iData interface{}
			_ = json.Unmarshal(bt, &iData)

			if iv, ok := iData.([]interface{}); ok {
				for i, i2 := range iv {
					switch i2.(type) {
					case string:
						paramsArr = append(paramsArr, fmt.Sprintf("%s%s[]=%v", key, k, i2))
					case float64:
						paramsArr = append(paramsArr, fmt.Sprintf("%s%s[]=%v", key, k, i2))
					default:
						vk := key + k + "[" + strconv.Itoa(i) + "]"
						paramsArr = append(paramsArr, encodeMap(i2.(map[string]interface{}), vk))
					}
				}
			} else if im, ok := iData.(map[string]interface{}); ok {
				for ik, i2 := range im {
					switch i2.(type) {
					case string:
						paramsArr = append(paramsArr, fmt.Sprintf("%s%s[%v]=%v", key, k, ik, i2))
					case float64:
						paramsArr = append(paramsArr, fmt.Sprintf("%s%s[%v]=%v", key, k, ik, i2))
					case []interface{}:
						vk := key + k + "[" + ik + "]"
						for k3, v3 := range i2.([]interface{}) {
							switch v3.(type) {
							case string:
								paramsArr = append(paramsArr, fmt.Sprintf("%s%s[%v]=%v", key, vk, k3, v3))
							case float64:
								paramsArr = append(paramsArr, fmt.Sprintf("%s%s[%v]=%v", key, vk, k3, v3))
							default:
								if key != "" {
									k = key + k
								}

								paramsArr = append(paramsArr, encodeMap(i2.(map[string]interface{}), k))
							}
						}
					default:
						k = key + k + "[" + ik + "]"

						paramsArr = append(paramsArr, encodeMap(i2.(map[string]interface{}), k))
					}
				}
			}
		}
	}
	return strings.Join(paramsArr, "&")
}
