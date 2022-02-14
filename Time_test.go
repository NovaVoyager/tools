package tools

import "testing"

func TestTimeStr(t *testing.T) {
	t.Log(TimeDateFormat, "-->", TimeStr(TimeDateFormat, cstZone))
	t.Log(DateFormat, "-->", TimeStr(DateFormat, cstZone))
}
