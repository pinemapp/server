package locale

import (
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/i18n"
	"github.com/pinem/server/config"
)

type Translate struct {
	tfunc i18n.TranslateFunc
}

func InitLocale() error {
	conf := config.Get()

	defaultPath := path.Join(conf.I18n.Dir, conf.I18n.Default+".yaml")
	if err := i18n.LoadTranslationFile(defaultPath); err != nil {
		return err
	}

	paths, err := filepath.Glob(path.Join(conf.I18n.Dir, "*.yaml"))
	if err != nil {
		return err
	}
	for _, file := range paths {
		if err := i18n.LoadTranslationFile(file); err != nil {
			return err
		}
	}
	return nil
}

func Get(c *gin.Context) *Translate {
	accept := c.GetHeader("Accept-Language")
	lang := c.Query("lang")
	cookie, _ := c.Cookie("lang")
	defaultLang := config.Get().I18n.Default
	return &Translate{i18n.MustTfunc(accept, lang, cookie, defaultLang)}
}

func (tran *Translate) T(key string, args ...interface{}) string {
	return tran.tfunc(key, args...)
}
