package transform

func init() {
	RegisterTranslator(&Pipeline{}, &Switcher{})
	RegisterTranslator(&RegExpReplacer{})
	RegisterTranslator((*StrCase)(nil), (*Formatter)(nil))
	RegisterTranslator((*Cast)(nil))
}
