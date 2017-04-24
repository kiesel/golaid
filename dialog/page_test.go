package dialog

import "testing"
import "reflect"

import "github.com/davecgh/go-spew/spew"
import "github.com/yvasiyarov/php_session_decoder/php_serialize"

func Test_Parse_to_Page(t *testing.T) {
	actual, err := ParsePage(`a:3:{s:5:"total";i:98;s:7:"perpage";i:5;s:7:"entries";a:5:{s:25:"20131125000000-zurich.dat";s:6:"zurich";s:34:"20120925000000-frankreich-2012.dat";s:15:"frankreich-2012";s:29:"20120528000000-birkweiler.dat";s:10:"birkweiler";s:34:"20111003155912-fallout-shelter.dat";s:15:"fallout-shelter";s:37:"20110820123913-dresden-2011-08-20.dat";s:18:"dresden-2011-08-20";}}`)

	if err != nil {
		t.Fail()
	}

	entries := make([]EntryRef, 0, 8)
	entries = append(entries,
		NewEntryRef("20131125000000-zurich.dat", "zurich"),
		NewEntryRef("20120925000000-frankreich-2012.dat", "frankreich-2012"),
		NewEntryRef("20120528000000-birkweiler.dat", "birkweiler"),
		NewEntryRef("20111003155912-fallout-shelter.dat", "fallout-shelter"),
		NewEntryRef("20110820123913-dresden-2011-08-20.dat", "dresden-2011-08-20"),
	)
	expected := Page{
		Total:   int64(98),
		Perpage: int64(5),
		Entries: entries,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Error("Actual value did not meet expected", spew.Sdump(actual), spew.Sdump(expected))
	}
}

func ParsePage(in string) (Page, error) {
	data, err := php_serialize.UnSerialize(in)
	if err != nil {
		return Page{}, err
	}

	return NewPage(data.(php_serialize.PhpArray))
}
