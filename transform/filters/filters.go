package filters

import "github.com/aynakeya/deepcolor/transform"

func init() {
	transform.RegisterFilter(&RegExpFilter{})
}
