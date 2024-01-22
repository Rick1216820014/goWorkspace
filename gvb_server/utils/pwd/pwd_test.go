package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("1234"))
}
func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$tHanE7jwkTDlc0npM1Cwk.QaD.Qe0GZLpRfVZsHdvby.X/d4wlDtW", "1234"))
}
