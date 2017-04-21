package dialog

import (
	"fmt"

	"github.com/yvasiyarov/php_session_decoder/php_serialize"
)

// Image represents Dialog Image
type Image struct {
	Name     string
	Width    int64
	Height   int64
	ExifData ExifData
	IptcData IptcData
}

func newImage(in *php_serialize.PhpObject) (Image, error) {
	exif, err := newExifData(orNil(in.GetPublic("exifData")).(*php_serialize.PhpObject))
	if err != nil {
		return Image{}, err
	}

	return Image{
		Name:     getFieldString(in, "name"),
		Width:    getFieldInt64(in, "width"),
		Height:   getFieldInt64(in, "height"),
		ExifData: exif,
		// TODO ExifData: ...
		// TODO IptcData: ...
	}, nil
}

func newImages(in *php_serialize.PhpArray) ([]Image, error) {
	out := []Image{}

	for _, item := range *in {
		if item == nil {
			fmt.Println("Given image was nil, cannot import")

			// Still insert empty image
			out = append(out, Image{})
			continue
		}

		image, err := newImage(item.(*php_serialize.PhpObject))
		if err != nil {
			return nil, err
		}

		out = append(out, image)
	}

	return out, nil
}
