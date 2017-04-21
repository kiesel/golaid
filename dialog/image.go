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
	fmt.Println(in)
	return Image{
		Name:   getFieldString(in, "name"),
		Width:  getFieldInt64(in, "width"),
		Height: getFieldInt64(in, "height"),
		// TODO ExifData: ...
		// TODO IptcData: ...
	}, nil
}

func newHighlights(in *php_serialize.PhpArray) ([]Image, error) {
	if in == nil {
		return []Image{}, nil
	}

	out := make([]Image, len(*in))

	for _, item := range *in {
		image, err := newImage(item.(*php_serialize.PhpObject))
		if err != nil {
			return nil, err
		}

		out = append(out, image)
	}

	return out, nil
}
