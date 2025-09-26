package sections

import (
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

/**
* Setup Raw Score section
* Params: *excelize.File
* Returns: Row that it ends on
 */
func RawScoreSection(wrkbook *excelize.File, cursheet string, cardnames []string, decknames []string) {
	//setup section with repeatables
	Headings(wrkbook, 2, true, cursheet, cardnames, decknames)
	Matchups(wrkbook, 2, cursheet, cardnames, decknames)
	WeightedSum(wrkbook, 2, false, cursheet, cardnames, decknames)
	ScoreVsMean(wrkbook, 2, false, cursheet, cardnames, decknames)
	//setup other info
	//setup boldstyle for bold headings
	validboldstyle := true
	boldstyle, err := wrkbook.NewStyle(`{"font":{"bold":true}}`)
	if err != nil {
		validboldstyle = false
	}
	//standard headings and key
	wrkbook.SetCellValue(cursheet, "A1", "<Put Name of Sheet Here>")
	if validboldstyle {
		wrkbook.SetCellStyle(cursheet, "A1", "A1", boldstyle)
	}
	wrkbook.SetCellValue(cursheet, "A2", "Blowout = 4pts")
	wrkbook.SetCellValue(cursheet, "A3", "High Impact = 4pts")
	wrkbook.SetCellValue(cursheet, "A4", "Mid Impact = 2 pts")
	wrkbook.SetCellValue(cursheet, "A5", "Low Impact = 1pt")
	wrkbook.SetCellValue(cursheet, "A6", "Non-Factor = 0pts")
	wrkbook.SetCellValue(cursheet, "C2", "# represented")
	if validboldstyle {
		wrkbook.SetCellStyle(cursheet, "C2", "C2", boldstyle)
	}
	wrkbook.SetCellValue(cursheet, "D2", "% frequency")
	if validboldstyle {
		wrkbook.SetCellStyle(cursheet, "D2", "D2", boldstyle)
	}
	wrkbook.SetCellValue(cursheet, "B"+strconv.Itoa(3+len(decknames)), "Raw Score")
	if validboldstyle {
		wrkbook.SetCellStyle(cursheet, "B"+strconv.Itoa(3+len(decknames)), "B"+strconv.Itoa(3+len(decknames)), boldstyle)
	}
}
