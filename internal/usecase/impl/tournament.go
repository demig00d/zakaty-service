package impl

import (
	"fmt"
	"github.com/demig00d/zakaty-service/config"
	"github.com/demig00d/zakaty-service/internal/usecase/impl/columns"
	"sort"
	"strconv"
	"strings"

	"github.com/demig00d/zakaty-service/internal/usecase"
	"github.com/demig00d/zakaty-service/pkg/puzzlebot"
	"github.com/demig00d/zakaty-service/pkg/sheets"
)

type TournamentImpl struct {
	column      columns.Columns
	spreadsheet sheets.Spreadsheet
}

func NewTournamentImpl(spreadsheet sheets.Spreadsheet, cfgColumns config.Columns) TournamentImpl {
	return TournamentImpl{
		column:      columns.FromConfigColumns(cfgColumns),
		spreadsheet: spreadsheet,
	}
}

func (t TournamentImpl) GetRating(user puzzlebot.User) (usecase.Raiting, error) {

	resp, err := t.spreadsheet.Get()
	if err != nil {
		return "", err
	}

	persons := resp.Values
	sort.SliceStable(persons, func(i, j int) bool {
		value1, _ := strconv.Atoi(persons[i][t.column.Sum].(string))
		value2, _ := strconv.Atoi(persons[j][t.column.Sum].(string))

		return value1 > value2
	})

	isIn := false
	var rating strings.Builder
	for i, person := range persons {
		if isIn && i > 3 {
			break
		}

		personId, _ := strconv.Atoi(person[0].(string))
		if len(person) > 0 && personId == user.Id {
			if i > 3 {
				rating.WriteString(".\n.\n.\n")
			}
			rating.WriteString(fmt.Sprintf("<b>%d. %s</b>\n", i+1, t.column.FormatPerson(person)))

			isIn = true
			continue
		}

		if i < 3 {
			rating.WriteString(fmt.Sprintf("%d. %s\n", i+1, t.column.FormatPerson(person)))
		}
	}

	return rating.String(), nil

}
