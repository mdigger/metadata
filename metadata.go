// Библиотека для разбора и работы с метаданными (YAML front matter) в файлах для создания
// статических сайтов и блогов.
package metadata

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Metadata описывает представление метаданных для документов и проектов.
type Metadata map[string]interface{}

// Get возвращает строковое представление значения, содержащегося в метаданных под указанным
// именем. В том случае, если данных с таким именем нет, то будет возвращена пустая строка.
func (self Metadata) Get(name string) string {
	if result := self[name]; result != nil {
		return strings.TrimSpace(fmt.Sprint(result))
	}
	return ""
}

// GetQuickList возвращает значение, хранящееся под указанным именем в виде списка строк.
// Если там изначально хранился именно список строк, то он и будет возвращен. Если же значение
// представлено в виде строки, то оно будет разбито на несколько строк: в качестве разделителя
// выступает любой символ, который не является буквой, цифрой, подчеркиванием или тире. Во всех
// остальных случаях будет возвращен пустой список.
func (self Metadata) GetQuickList(name string) []string {
	switch data := self[name].(type) {
	case []string:
		return data
	case []interface{}:
		list := make([]string, len(data))
		for i, value := range data {
			list[i] = fmt.Sprint(value)
		}
		return list
	case string:
		return strings.FieldsFunc(data, func(c rune) bool {
			return c != '_' && c != '-' && !unicode.IsLetter(c) && !unicode.IsNumber(c)
		})
	default:
		return nil
	}
}

// Разделитель, используемый для разделения строки на части.
var reSplitter = regexp.MustCompile(`\s*[;,]\s*`)

// GetList возвращает список строк, хранящихся под указанным именем. Если там хранится строка, то
// она разбивается на отдельные строки. В качестве разделителей выступают запятая и точка с запятой.
func (self Metadata) GetList(name string) []string {
	switch data := self[name].(type) {
	case []string:
		return data
	case []interface{}:
		list := make([]string, len(data))
		for i, value := range data {
			list[i] = fmt.Sprint(value)
		}
		return list
	case string:
		return reSplitter.Split(strings.TrimSpace(data), -1)
	default:
		return nil
	}
}

// SupportedDatetimeFormats содержит список поддерживаемых форматов даты и времени, которые
// используются для разбора даты.
var SupportedDatetimeFormats = []string{
	"2006-01-02T15:04:05Z",
	"2006-01-02T15:04:05",
	"2006-01-02T15:04Z",
	"2006-01-02T15:04",
	"2006-01-02 15:04:05Z",
	"2006-01-02 15:04:05",
	"2006-01-02 15:04Z",
	"2006-01-02 15:04",
	"2006-01-02Z",
	"2006-01-02",
	"2006-01",
	"2006",
}

// GetDate возвращает значение метаданных с указанным именем как дату. Если не удалось получить
// данные о дате или преобразовать их из исходного формата, то возвращается пустая дата.
func (self Metadata) GetDate(name string) time.Time {
	switch data := self[name].(type) {
	case time.Time:
		return data
	case int64:
		return time.Unix(data, 0)
	case string:
		for _, format := range SupportedDatetimeFormats {
			var loc *time.Location
			if strings.ContainsRune(format, 'Z') {
				loc = time.UTC
			} else {
				loc = time.Local
			}
			if pdate, err := time.ParseInLocation(format, data, loc); err == nil {
				return pdate
			}
		}
		return time.Time{}
	default:
		return time.Time{}
	}
}

// GetBool возвращает true, если значение определено и похоже на утверждение.
func (self Metadata) GetBool(name string) bool {
	switch data := self[name].(type) {
	case bool:
		return data
	case string:
		value, _ := strconv.ParseBool(data)
		return value
	case int:
		return data > 0
	default:
		return false
	}
}

// GetInt возвращает числовое значение или ноль.
func (self Metadata) GetInt(name string) int {
	switch data := self[name].(type) {
	case string:
		value, _ := strconv.Atoi(data)
		return value
	case int:
		return data
	default:
		return 0
	}
}

// GetSubMetadata возвращает значение с указанным ключем как метаданные.
func (self Metadata) GetSubMetadata(name string) Metadata {
	switch data := self[name].(type) {
	case Metadata:
		return data
	case map[interface{}]interface{}:
		metadata := make(Metadata, len(data))
		for key, value := range data {
			metadata[fmt.Sprint(key)] = value
		}
		return metadata
	default:
		return make(Metadata)
	}
}

// Предопределенные имена полей метаданных.
const (
	MetanameTitle       = "title"       // Заголовок
	MetanameSubtitle    = "subtitle"    // Подзаголовок
	MetanameDescription = "description" // Описание
	MetanameKeywords    = "keywords"    // Ключевые слова
	MetanameTags        = "tags"        // Теги
	MetanameCategories  = "categories"  // Категории
	MetanameDate        = "date"        // Дата
	MetanameAuthor      = "author"      // Автор
	MetanameTemplate    = "layout"      // Название шаблона
	MetanameLang        = "lang"        // Язык
	MetanameDraft       = "draft"       // Флаг черновика
)

// Title возвращает название.
func (self Metadata) Title() string {
	return self.Get(MetanameTitle)
}

// Subtitle возвращает подзаголовок.
func (self Metadata) Subtitle() string {
	return self.Get(MetanameSubtitle)
}

// Description возвращает описание.
func (self Metadata) Description() string {
	return self.Get(MetanameDescription)
}

// Keywords возвращает список ключевых слов.
func (self Metadata) Keywords() []string {
	return self.GetList(MetanameKeywords)
}

// Tags возвращает список тегов.
func (self Metadata) Tags() []string {
	return self.GetQuickList(MetanameTags)
}

// Categories возвращает список категорий.
func (self Metadata) Categories() []string {
	return self.GetList(MetanameCategories)
}

// Layout возвращает название шаблона.
func (self Metadata) Layout() string {
	return self.Get(MetanameTemplate)
}

// Authors возвращает список авторов.
func (self Metadata) Authors() []string {
	return self.GetList(MetanameAuthor)
}

// Author возвращает первого автора из списка авторов.
func (self Metadata) Author() string {
	return self.Authors()[0]
}

// Date возвращает дату из метаинформации.
func (self Metadata) Date() time.Time {
	return self.GetDate(MetanameDate)
}

// Lang возвращает язык документа.
func (self Metadata) Lang() string {
	return self.Get(MetanameLang)
}
