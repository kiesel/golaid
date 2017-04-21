package dialog

import "testing"
import "reflect"
import "github.com/yvasiyarov/php_session_decoder/php_serialize"
import "github.com/davecgh/go-spew/spew"

func Test_NewPage(t *testing.T) {
	p, err := NewPage(php_serialize.PhpArray{
		php_serialize.PhpValue("total"):   php_serialize.PhpValue(int64(15)),
		php_serialize.PhpValue("perpage"): php_serialize.PhpValue(int64(5)),
		php_serialize.PhpValue("entries"): php_serialize.PhpArray{
			php_serialize.PhpValue("201504171010-filename.dat"): php_serialize.PhpValue("filename"),
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
		t.Errorf("Actual does not match expected.\n%s\n\n%s\n", spew.Sdump(expect), spew.Sdump(given))
	}
}
