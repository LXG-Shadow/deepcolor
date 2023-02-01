package transform

func init() {
	_RegisteredTranslator["Pipeline"] = &Pipeline{}
	_RegisteredTranslator["Switcher"] = &Switcher{}
	_RegisteredTranslator["Foreach"] = &Foreach{}
	_RegisteredTranslator["RegExpReplacer"] = &RegExpReplacer{}
	_RegisteredTranslator["StrCase"] = &StrCase{}
	_RegisteredTranslator["Formatter"] = &Formatter{}
	_RegisteredTranslator["Cast"] = &Cast{}
	_RegisteredTranslator["Setter"] = &Setter{}
	//RegisterTranslator(NewPipeline(), NewSwitcher(), &Foreach{})
	//RegisterTranslator(&RegExpReplacer{})
	//RegisterTranslator((*StrCase)(nil), (*Formatter)(nil))
	//RegisterTranslator((*Cast)(nil))
}
