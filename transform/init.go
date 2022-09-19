package transform

func init() {
	RegisterTranslator(&Pipeline{}, &Switcher{}, &Foreach{})
	RegisterTranslator(&RegExpReplacer{})
	RegisterTranslator((*StrCase)(nil), (*Formatter)(nil))
	RegisterTranslator((*Cast)(nil))
}
