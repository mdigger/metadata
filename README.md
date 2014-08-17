# Работа с метаданными (YAML front matter)

    import "github.com/mdigger/metadata"

Библиотека для разбора и работы с метаданными (YAML front matter) в файлах для
создания статических сайтов и блогов.

## Использование

```go
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
}
```
SupportedDatetimeFormats содержит список поддерживаемых форматов даты и времени,
которые используются для разбора даты.

#### type Metadata

```go
type Metadata map[string]interface{}
```

Metadata описывает представление метаданных для документов и проектов.

#### func  ReadFile

```go
func ReadFile(filename string) (metadata Metadata, data []byte, err error)
```
ReadFile читает файл с метаданными, разбирает его и возвращает вместе с
оставшимся текстом.

#### func (Metadata) Author

```go
func (self Metadata) Author() string
```
Author возвращает первого автора из списка авторов.

#### func (Metadata) Authors

```go
func (self Metadata) Authors() []string
```
Authors возвращает список авторов.

#### func (Metadata) Categories

```go
func (self Metadata) Categories() []string
```
Categories возвращает список категорий.

#### func (Metadata) Date

```go
func (self Metadata) Date() time.Time
```
Date возвращает дату из метаинформации.

#### func (Metadata) Description

```go
func (self Metadata) Description() string
```
Description возвращает описание.

#### func (Metadata) Get

```go
func (self Metadata) Get(name string) string
```
Get возвращает строковое представление значения, содержащегося в метаданных под
указанным именем. В том случае, если данных с таким именем нет, то будет
возвращена пустая строка.

#### func (Metadata) GetBool

```go
func (self Metadata) GetBool(name string) bool
```
GetBool возвращает true, если значение определено и похоже на утверждение.

#### func (Metadata) GetDate

```go
func (self Metadata) GetDate(name string) time.Time
```
GetDate возвращает значение метаданных с указанным именем как дату. Если не
удалось получить данные о дате или преобразовать их из исходного формата, то
возвращается пустая дата.

#### func (Metadata) GetList

```go
func (self Metadata) GetList(name string) []string
```
GetList возвращает список строк, хранящихся под указанным именем. Если там
хранится строка, то она разбивается на отдельные строки. В качестве разделителей
выступают запятая и точка с запятой.

#### func (Metadata) GetQuickList

```go
func (self Metadata) GetQuickList(name string) []string
```
GetQuickList возвращает значение, хранящееся под указанным именем в виде списка
строк. Если там изначально хранился именно список строк, то он и будет
возвращен. Если же значение представлено в виде строки, то оно будет разбито на
несколько строк: в качестве разделителя выступает любой символ, который не
является буквой, цифрой, подчеркиванием или тире. Во всех остальных случаях
будет возвращен пустой список.

#### func (Metadata) GetSubMetadata

```go
func (self Metadata) GetSubMetadata(name string) Metadata
```
GetSubMetadata возвращает значение с указанным ключем как метаданные.

#### func (Metadata) Keywords

```go
func (self Metadata) Keywords() []string
```
Keywords возвращает список ключевых слов.

#### func (Metadata) Layout

```go
func (self Metadata) Layout() string
```
Layout возвращает название шаблона.

#### func (Metadata) Subtitle

```go
func (self Metadata) Subtitle() string
```
Subtitle возвращает подзаголовок.

#### func (Metadata) Tags

```go
func (self Metadata) Tags() []string
```
Tags возвращает список тегов.

#### func (Metadata) Title

```go
func (self Metadata) Title() string
```
Title возвращает название.
