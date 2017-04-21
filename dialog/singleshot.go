package dialog

import "github.com/yvasiyarov/php_session_decoder/php_serialize"

// SingleShot is a single shot
type SingleShot struct {
	*Entry
}

func newSingleShot(in *php_serialize.PhpObject) (*SingleShot, error) {
	return nil, nil
}
