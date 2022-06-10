package awesome

import (
	"fmt"
	"testing"
)

func Test_exist(t *testing.T) {
//[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]
	oParam := [][]string{{"A", "B", "C", "E"}, {"S", "F", "E", "S"}, {"A", "D", "E", "E"}}

	param := make([][]byte, len(oParam), len(oParam))
	for i, v := range oParam {
		param[i] = make([]byte, len(v), len(v))
		for j, s := range v {
			param[i][j] = []byte(s)[0]
		}
	}

	got := exist(param, "ABCEFSADEESE")
	fmt.Println(got)
}
