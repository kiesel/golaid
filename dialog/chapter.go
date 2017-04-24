package dialog

import "github.com/yvasiyarov/php_session_decoder/php_serialize"

// Chapter represents a chapter
type Chapter struct {
	Name   string
	Images []Image
}

func newChapters(in *php_serialize.PhpArray) ([]Chapter, error) {
	out := []Chapter{}
	for _, val := range *in {
		obj := val.(*php_serialize.PhpObject)
		chapter, err := newChapter(obj)
		if err != nil {
			return nil, err
		}
		out = append(out, chapter)
	}

	return out, nil
}

func newChapter(in *php_serialize.PhpObject) (Chapter, error) {
	images, err := newImages(phpArrayFrom(in.GetPublic("images")))
	if err != nil {
		return Chapter{}, err
	}

	out := Chapter{
		Name:   getFieldString(in, "name"),
		Images: images,
	}

	return out, nil
}
