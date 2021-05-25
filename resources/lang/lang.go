package lang

import (
	"github.com/ArtisanCloud/go-framework/resources/lang/en_US"
	"github.com/ArtisanCloud/go-framework/resources/lang/zh_TW"
)

func LoadLanguages()  {
	en_US.LoadLang()
	en_TW.LoadLang()
}