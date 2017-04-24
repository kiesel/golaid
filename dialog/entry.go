package dialog

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/yvasiyarov/php_session_decoder/php_serialize"
)

const (
	// PhpDateStringFormat is the default date parsing format for XP Framework date objects
	PhpDateStringFormat = "2006-01-02 15:04:05-0700"
)

// Entry represents a base entry
type Entry struct {
	Name        string    `php:"name"`
	Title       string    `php:"title"`
	Description string    `php:"description"`
	CreatedAt   time.Time `php:"createdAt"`
}

// IEntry is the entries interface
type IEntry interface {
	GetName() string
	// GetTitle() string
	// GetDescription() string
	// GetCreatedAt() time.Time
}

// GetName retrieves the name of the entry
func (e *Entry) GetName() string {
	return e.Name
}

// NewEntry constructs a new dialog entry structure
func NewEntry(in *php_serialize.PhpObject) (IEntry, error) {
	switch in.GetClassName() {
	case "Album", "de.thekid.dialog.Album":
		_, _ = newObject(Album{}, in)
		return nil, nil
		// return newAlbum(in)

	case "EntryCollection", "de.thekid.dialog.EntryCollection":
		return newEntryCollection(in)

	case "SingleShot", "de.thekid.dialog.SingleShot":
		return newSingleShot(in)

	default:
		return nil, errors.New("Cannot convert class " + in.GetClassName())
	}
}

func newObject(orig interface{}, in *php_serialize.PhpObject) (interface{}, error) {
	// fmt.Printf("Entering with %T \n", orig)
	// defer fmt.Println("Leaving newObject")

	// orig contains the struct to be filled; but it is a value, not a pointer, so we cannot change it
	// through reflection. Instead, create a copy

	// Create copy, assign pointer to reflection object to out
	out := reflect.New(reflect.TypeOf(orig))

	// Fetch the fields of the original struct:
	t := reflect.TypeOf(orig)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// fmt.Printf("Inspecting type %v\n", field)

		if field.Anonymous {
			if field.Type == reflect.TypeOf(&Entry{}) {

				// Create new object; static type is interface{}
				entry, err := newObject(Entry{}, in)
				if err != nil {
					return nil, err
				}

				// Type assertion to convert static type to dialog.Entry
				pentry := entry.(Entry)

				// Assign value to ptr to entry to the field
				out.Elem().FieldByIndex(field.Index).Set(reflect.ValueOf(&pentry))
				continue
			}
		}
		// if _, ok := field.Tag.Lookup("recurse"); ok {

		// 	// Create a new instance of orig's underlying type
		// 	// spew.Dump(field.Type)
		// 	entry := reflect.New(reflect.TypeOf(orig))
		// 	// spew.Dump(entry.Elem().Interface())

		// 	entryRef, err := newObject(entry.Elem(), in)
		// 	// spew.Dump(entryRef)

		// 	if err != nil {
		// 		return nil, err
		// 	}

		// 	// out.Elem().FieldByIndex(field.Index).Set(reflect.ValueOf(entryRef))
		// 	continue
		// }

		if phpFieldName, ok := field.Tag.Lookup("php"); ok {

			// Acquire value from PHP object
			value, ok := in.GetPublic(phpFieldName)
			if !ok {
				fmt.Printf("Cannot set value on field %s from %s\n", field.Name, phpFieldName)
				continue
			}

			if value == nil {
				fmt.Println("Skipping setting nil value")
				continue
			}

			// Fetch the reflection object to assign the new value to; .Elem() is the
			// equivalent of dereferencing the pointer
			// This value must be addressable and setable (which it is, because we obtained
			// it through a pointer)
			assignTo := out.Elem().FieldByIndex(field.Index)

			switch field.Type.Kind() {
			case reflect.TypeOf(time.Time{}).Kind():
				dateString, ok := value.(*php_serialize.PhpObject).GetPublic("value")
				if !ok {
					fmt.Println("Could not get date")
					break
				}
				time, err := time.Parse(PhpDateStringFormat, dateString.(string))
				if err != nil {
					fmt.Printf("Could not parse date: %v\n", err)
				}
				assignTo.Set(reflect.ValueOf(time))
			case reflect.Int:
			case reflect.Int64:
				assignTo.SetInt(php_serialize.PhpValueInt64(value))

			case reflect.String:
				assignTo.SetString(php_serialize.PhpValueString(value))

			case reflect.Bool:
				assignTo.SetBool(php_serialize.PhpValueBool(value))
			}
		}
	}

	// Return the actual copied value to the caller
	// Elem() dereferences the reflection pointer, Interface() then retrieves
	// the actual interface value
	return out.Elem().Interface(), nil
}

func phpArrayFrom(in php_serialize.PhpValue, ok bool) *php_serialize.PhpArray {
	if !ok {
		return nil
	}

	out := in.(php_serialize.PhpArray)
	return &out
}

func getFieldString(p *php_serialize.PhpObject, field string) string {
	val, ok := p.GetPublic(field)
	if !ok {
		return ""
	}

	return php_serialize.PhpValueString(val)
}

func getFieldInt64(p *php_serialize.PhpObject, field string) int64 {
	val, ok := p.GetPublic(field)
	if !ok {
		return 0
	}

	return php_serialize.PhpValueInt64(val)
}

func orNil(in interface{}, ok bool) interface{} {
	if !ok {
		return nil
	}

	return in
}
