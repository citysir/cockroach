// generated by stringer -type Field; DO NOT EDIT

package log

import "fmt"

const _Field_name = "ClusterIDNodeIDStoreIDRaftIDMethodClientErrDetailKeymaxField"

var _Field_index = [...]uint8{0, 9, 15, 22, 28, 34, 40, 43, 49, 52, 60}

func (i Field) String() string {
	if i < 0 || i+1 >= Field(len(_Field_index)) {
		return fmt.Sprintf("Field(%d)", i)
	}
	return _Field_name[_Field_index[i]:_Field_index[i+1]]
}
