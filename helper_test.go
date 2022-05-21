package http_build_query

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	data := map[string]interface{}{
		"int":     1,
		"str":     "str",
		"arr_int": []int16{1, -2, 4},
		"m_arr": map[string][]int16{
			"test": []int16{1, -2, 4},
		},
		"m_m": []interface{}{
			map[string]string{"mo1": "v", "mo2": "v2"},
			map[string]string{"mo2": "v"},
		},
		"m_m_m": map[string]interface{}{
			"mm": struct{ Name string }{"张三"},
		},
	}

	str := Encode(data)
	// int=1&str=str&arr_int[]=1&arr_int[]=-2&arr_int[]=4&m_arr[test][0]=1&m_arr[test][1]=-2&m_arr[test][2]=4&m_m[0][mo1]=v&m_m[0][mo2]=v2&m_m[1][mo2]=v&m_m_m[mm][Name]=张三
	if 168 != len(str) {
		t.Error("解析不对")
	}
}

func TestEncode1(t *testing.T) {
	data := map[string]interface{}{
		"id":   1.0,
		"name": "john",
		"sub": map[string]interface{}{
			"0": map[string]interface{}{
				"1": map[string]interface{}{
					"2": 1.0,
				},
			},
			"1": 3.0,
			"2": "three",
		},
	}
	str := Encode(data)

	fmt.Println(str)
	got := "id=1&name=john&sub[0][1][2]=1&sub[1]=3&sub[2]=three"
	if str != got {
		t.Error("解析不对")
	}
}
