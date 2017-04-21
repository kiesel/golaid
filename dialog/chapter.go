package dialog

import "github.com/yvasiyarov/php_session_decoder/php_serialize"

// Chapter represents a chapter
type Chapter struct {
}

func newChapter(in *php_serialize.PhpArray) ([]Chapter, error) {
	return []Chapter{}, nil
}
