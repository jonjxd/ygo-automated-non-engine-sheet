package sections

import (
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

//---------------------------------------------------------------------------------
// HELPER FUNCTIONS
//---------------------------------------------------------------------------------

func Headings(wrkbook *excelize.File, row int, cardnames []string, decknames []string, start bool, cursheet string) {
	//start from E, go to end of thing
	curcell := ""
	i := 0
	if start {
		for i = 0; i < len(cardnames); i++ {
			curcell = lettermapping[4+i] + strconv.Itoa(row)
			wrkbook.SetCellValue(cursheet, curcell, cardnames[i])
		}
	} else {
		for i = 0; i < len(cardnames); i++ {
			curcell = lettermapping[4+i] + strconv.Itoa(row)
			formula := lettermapping[4+i] + "2"
			wrkbook.SetCellFormula(cursheet, curcell, formula)
		}
	}
}

func Matchups(wrkbook *excelize.File, row int, cardnames []string, decknames []string, cursheet string) {
	//set up matchups in B column
	wrkbook.SetCellValue(cursheet, "B"+strconv.Itoa(row), "Matchup")
	//setup bold style
	boldstyle, err := wrkbook.NewStyle(`{"font":{"bold":true}}`)
	if err == nil {
		wrkbook.SetCellStyle(cursheet, "B"+strconv.Itoa(row), "B"+strconv.Itoa(row), boldstyle)
	}
	row += 1
	index := row
	for index < row+len(decknames) {
		wrkbook.SetCellValue(cursheet, "B"+strconv.Itoa(index), decknames[index-row])
		index++
	}

}

func SetStandardFormulas(wrkbook *excelize.File, row int, cardnames []string, decknames []string, cursheet string) {
	//setup number format
	numstyle, err := wrkbook.NewStyle(`{"custom_number_format": "0.0000_ "}`)
	//separate counter for looping through formula
	index := 3
	//set row to next row
	row++
	//first loop: row
	for i := row; i < len(decknames)+row; i++ {
		//second loop: column
		for j := 4; j < len(cardnames)+4; j++ {
			curCell := lettermapping[j] + strconv.Itoa(i)
			formula := "D" + strconv.Itoa(i) + "*" + lettermapping[j] + strconv.Itoa(index)
			wrkbook.SetCellFormula(cursheet, curCell, formula)
			if err == nil {
				wrkbook.SetCellStyle(cursheet, curCell, curCell, numstyle)
			}
		}
		index++
	}
}

func WeightedSum(wrkbook *excelize.File, row int, cardnames []string, decknames []string, space bool, cursheet string) {
	thisrow := 0
	if !space {
		thisrow = row + len(decknames) + 1
	} else {
		thisrow = row + len(decknames) + 2
	}
	wrkbook.SetCellValue(cursheet, "B"+strconv.Itoa(thisrow), "Weighted Sum")
	//setup bold style
	boldstyle, err := wrkbook.NewStyle(`{"font":{"bold":true}}`)
	if err == nil {
		wrkbook.SetCellStyle(cursheet, "B"+strconv.Itoa(thisrow), "B"+strconv.Itoa(thisrow), boldstyle)
	}
	for j := 4; j < len(cardnames)+4; j++ {
		curCell := lettermapping[j] + strconv.Itoa(thisrow)
		formula := "SUM(" + lettermapping[j] + strconv.Itoa(row+1) + ":"
		if !space {
			formula = formula + lettermapping[j] + strconv.Itoa(thisrow-1) + ")"
		} else {
			formula = formula + lettermapping[j] + strconv.Itoa(thisrow-2) + ")"
		}
		wrkbook.SetCellFormula(cursheet, curCell, formula)
	}
}

func ScoreVsMean(wrkbook *excelize.File, row int, cardnames []string, decknames []string, space bool, cursheet string) {
	//setup number format
	numstyle, err := wrkbook.NewStyle(`{"custom_number_format": "0.0000_ "}`)
	//recalculate row to correct place
	if space {
		row += len(decknames) + 3
	} else {
		row += len(decknames) + 2
	}
	//set name of row
	wrkbook.SetCellValue(cursheet, "B"+strconv.Itoa(row), "Score Vs Mean")
	//setup bold style
	boldstyle, err := wrkbook.NewStyle(`{"font":{"bold":true}}`)
	if err == nil {
		wrkbook.SetCellStyle(cursheet, "B"+strconv.Itoa(row), "B"+strconv.Itoa(row), boldstyle)
	}
	//set row string
	formstr := strconv.Itoa(row - 1)
	//grab weighted sums and conver to formula
	formula := "SUM(E" + formstr + ":" + lettermapping[len(cardnames)+3] + formstr + ")/" + strconv.Itoa(len(cardnames))
	//save formula cell to save work later
	formcell := "C" + strconv.Itoa(row)
	wrkbook.SetCellFormula(cursheet, formcell, formula)
	//loop through and set score vs mean formula
	for j := 4; j < len(cardnames)+4; j++ {
		formula = lettermapping[j] + formstr + "/" + formcell
		wrkbook.SetCellFormula(cursheet, lettermapping[j]+strconv.Itoa(row), formula)
		if err == nil {
			wrkbook.SetCellStyle(cursheet, lettermapping[j]+strconv.Itoa(row), lettermapping[j]+strconv.Itoa(row), numstyle)
		}
	}

}

func Frequency(wrkbook *excelize.File, row int, cardnames []string, decknames []string, cursheet string) {
	//setup number format
	numstyle, err := wrkbook.NewStyle(`{"custom_number_format": "0.0000_ "}`)
	//set Total Number of Decks
	thisrow := row + len(decknames) + 1
	formula := "SUM(C" + strconv.Itoa(row+1) + ":C" + strconv.Itoa(thisrow-1) + ")"
	wrkbook.SetCellValue(cursheet, "B"+strconv.Itoa(thisrow), "Total # of Decks")
	//setup bold style
	boldstyle, err := wrkbook.NewStyle(`{"font":{"bold":true}}`)
	if err == nil {
		wrkbook.SetCellStyle(cursheet, "B"+strconv.Itoa(thisrow), "B"+strconv.Itoa(thisrow), boldstyle)
	}
	wrkbook.SetCellFormula(cursheet, "C"+strconv.Itoa(thisrow), formula)
	//set new frequency formulas
	row += 1
	curCell := ""
	for row < thisrow {
		curCell = "D" + strconv.Itoa(row)
		formula = "C" + strconv.Itoa(row) + "/C" + strconv.Itoa(thisrow)
		wrkbook.SetCellFormula(cursheet, curCell, formula)
		if err == nil {
			wrkbook.SetCellStyle(cursheet, curCell, curCell, numstyle)
		}
		row++
	}
}
