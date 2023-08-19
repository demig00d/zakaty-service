package columns

import (
	"fmt"
	"github.com/demig00d/zakaty-service/config"
)

type Columns struct {
	TelegramId int
	Firstname  int
	Lastname   int
	Patronym   int
	Sum        int
}

func FromConfigColumns(cfgColumns config.Columns) Columns {
	return Columns{
		TelegramId: int(cfgColumns.TelegramId - 65),
		Firstname:  int(cfgColumns.Firstname - 65),
		Lastname:   int(cfgColumns.Lastname - 65),
		Patronym:   int(cfgColumns.Patronym - 65),
		Sum:        int(cfgColumns.Sum - 65),
	}
}

func (c Columns) FormatPerson(person []any) string {
	return fmt.Sprintf(
		"%s %s %s (%s)",
		person[c.Firstname],
		firstLetter(person[c.Lastname].(string)),
		firstLetter(person[c.Patronym].(string)),
		person[c.Sum].(string),
	)
}
func firstLetter(s string) string {
	if len(s) > 0 {
		return string(s[0] + '.')
	}

	return ""
}
