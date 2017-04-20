package dialog

import "testing"
import "reflect"

func Test_NewPage(t *testing.T) {
	p, err := NewPage(map[interface{}]interface{}{
		"total":   int64(15),
		"perpage": int64(5),
		"entries": map[interface{}]interface{}{
			"201504171010-filename.dat": "filename",
		},
	})

	if err != nil {
		t.Fail()
	}

	assertDeepEqual(Page{
		Total:   int64(15),
		Perpage: int64(5),
		Entries: []EntryRef{
			NewEntryRef("201504171010-filename.dat", "filename"),
		},
	}, p, t)
}

func assertDeepEqual(expect, given interface{}, t *testing.T) {
	if !reflect.DeepEqual(expect, given) {
		t.Errorf("Expected [%s] did not match given [%s].", expect, given)
	}
}
