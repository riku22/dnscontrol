// Code generated by "stringer -type=Verb"; DO NOT EDIT.

package diff2

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CREATE-1]
	_ = x[CHANGE-2]
	_ = x[DELETE-3]
}

const _Verb_name = "CREATECHANGEDELETE"

var _Verb_index = [...]uint8{0, 6, 12, 18}

func (i Verb) String() string {
	i -= 1
	if i < 0 || i >= Verb(len(_Verb_index)-1) {
		return "Verb(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Verb_name[_Verb_index[i]:_Verb_index[i+1]]
}