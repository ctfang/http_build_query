# http_build_query
go 版本的http_build_query 实现


install
~~~~shell
go get github.com/ctfang/http_build_query
~~~~
code
~~~~go
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

// echo 
// int=1&str=str&arr_int[]=1&arr_int[]=-2&arr_int[]=4&m_arr[test][0]=1&m_arr[test][1]=-2&m_arr[test][2]=4&m_m[0][mo1]=v&m_m[0][mo2]=v2&m_m[1][mo2]=v&m_m_m[mm][Name]=张三
~~~~
