package dialog

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/juju/loggo"
	"github.com/yvasiyarov/php_session_decoder/php_serialize"
)

const (
	// PhpDateStringFormat is the default date parsing format for XP Framework date objects
	PhpDateStringFormat = "2006-01-02 15:04:05-0700"
)

var logger = loggo.GetLogger("dialog.Entry")

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
	case "Album":
		album, err := newObject(Album{}, in)
		if err != nil {
			return nil, err
		}

		return album.(Album), nil

	case "EntryCollection":
		collection, err := newObject(EntryCollection{}, in)
		if err != nil {
			return nil, err
		}
		return collection.(EntryCollection), nil

	case "SingleShot":
		shot, err := newObject(SingleShot{}, in)
		if err != nil {
			return nil, err
		}
		return shot.(SingleShot), nil

	default:
		return nil, errors.New("Cannot convert class " + in.GetClassName())
	}
}

func newObject(orig interface{}, in *php_serialize.PhpObject) (interface{}, error) {
	logger.Debugf("Entering newObject() with a %T @ %p", orig, in)
	defer logger.Debugf("Leaving newObject() of %T @ %p", orig, in)

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
				logger.Infof("Cannot set value on field %s from %s\n", field.Name, phpFieldName)
				continue
			}

			if value == nil {
				logger.Debugf("Skipping setting nil value")
				continue
			}

			// Fetch the reflection object to assign the new value to; .Elem() is the
			// equivalent of dereferencing the pointer
			// This value must be addressable and setable (which it is, because we obtained
			// it through a pointer)
			assignTo := out.Elem().FieldByIndex(field.Index)

			logger.Debugf("Current field [%s] is a %s backed by %s",
				field.Name,
				field.Type.String(),
				field.Type.Kind().String(),
			)

			switch field.Type.Kind() {
			case reflect.Slice:
				switch field.Type {
				case reflect.TypeOf([]Image{}):
					input := value.(php_serialize.PhpArray)
					images := make([]Image, 0, len(input))

					logger.Debugf("Creating %d elements of type %T", cap(images), images)
					for i := 0; i < len(input); i++ {
						if input[i] == nil {
							logger.Debugf("Element %d is nil, skipping", i)
							continue
						}
						data := input[i].(*php_serialize.PhpObject)
						object, err := newObject(Image{}, data)
						if err != nil {
							return nil, err
						}

						images = append(images, object.(Image))
					}
					assignTo.Set(reflect.ValueOf(images))

				case reflect.TypeOf([]Chapter{}):
					input := value.(php_serialize.PhpArray)
					chapters := make([]Chapter, 0, len(input))

					logger.Debugf("Creating %d elements of type %T", cap(chapters), chapters)
					for i := 0; i < len(input); i++ {
						if input[i] == nil {
							logger.Debugf("Element %d is nil, skipping", i)
							continue
						}
						data := input[i].(*php_serialize.PhpObject)
						object, err := newObject(Chapter{}, data)
						if err != nil {
							return nil, err
						}

						chapters = append(chapters, object.(Chapter))
					}
					assignTo.Set(reflect.ValueOf(chapters))

				default:
					return nil, fmt.Errorf("Cannot convert structure, have %v", value)
				}

			case reflect.Struct:
				switch field.Type {
				case reflect.TypeOf(time.Time{}):
					dateString, ok := value.(*php_serialize.PhpObject).GetPublic("value")
					if !ok {
						logger.Infof("Could not convert date, have %v", value)
						break
					}

					time, err := time.Parse(PhpDateStringFormat, dateString.(string))
					if err != nil {
						logger.Infof("Could not parse date: %v\n", err)
					}
					assignTo.Set(reflect.ValueOf(time))

				case reflect.TypeOf(ExifData{}):
					data, ok := value.(*php_serialize.PhpObject)
					if !ok {
						logger.Infof("Could not convert exit data, have %v", value)
						break
					}

					exif, err := newObject(ExifData{}, data)
					if err != nil {
						return nil, err
					}
					assignTo.Set(reflect.ValueOf(exif))

				case reflect.TypeOf(IptcData{}):
					// TBI

				default:
					return nil, fmt.Errorf("Cannot convert type %T into %s", value, field.Type)

				}

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
