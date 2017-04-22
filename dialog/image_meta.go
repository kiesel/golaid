package dialog

import (
	"fmt"
	"reflect"
	"time"

	"github.com/yvasiyarov/php_session_decoder/php_serialize"
)

// ExifData represents exif data
type ExifData struct {
	ApertureFNumber string    `php:"apertureFNumber"`
	Software        string    `php:"software"`
	ExposureTime    string    `php:"exposureTime"`
	MeteringMode    int       `php:"meteringMode"`
	Flash           int       `php:"flash"`
	Orientation     int       `php:"orientation"`
	FileSize        int       `php:"fileSize"`
	DateTime        time.Time `php:"dateTime"`
	Model           string    `php:"model"`
	MimeType        string    `php:"mimeType"`
	WhiteBalance    int       `php:"whiteBalance"`
	FocalLength     int       `php:"focalLength"`
	Make            string    `php:"make"`
	ExposureProgram int       `php:"exposureProgram"`
}

// IptcData represents iptc data
type IptcData struct {
}

func newExifData(in *php_serialize.PhpObject) (ExifData, error) {
	out, err := reflectiveNewObject(&ExifData{}, in)
	if err != nil {
		return ExifData{}, err
	}

	return out.(ExifData), nil
}

func reflectiveNewObject(orig interface{}, in *php_serialize.PhpObject) (interface{}, error) {

	// orig contains the struct to be filled; but it is a value, not a pointer, so we cannot change it
	// through reflection. Instead, create a copy

	// Create copy, assign pointer to reflection object to out
	out := reflect.New(reflect.TypeOf(orig))

	// Fetch the fields of the original struct:
	t := reflect.TypeOf(orig)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
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
