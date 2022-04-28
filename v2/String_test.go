package tools

import "testing"

func TestKrand(t *testing.T) {
	t.Log(KcRandKindNum, Krand(6, KcRandKindNum))
	t.Log(KcRandKindLower, Krand(6, KcRandKindLower))
	t.Log(KcRandKindUpper, Krand(6, KcRandKindUpper))
	t.Log(KcRandKindAll, Krand(6, KcRandKindAll))
}
