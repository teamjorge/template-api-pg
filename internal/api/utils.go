package api

import (
	"encoding/json"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/exp/maps"
)

// parsePutColumns will retrieve the columns present in the request body for use with database Updates
//
// Returned Columns are in an indetermined order
func parsePutColumns(reqBody []byte) (boil.Columns, error) {
	colMap := make(map[string]interface{})

	if err := json.Unmarshal(reqBody, &colMap); err != nil {
		return boil.None(), err
	}

	columnsToUpdate := maps.Keys(colMap)

	return boil.Whitelist(columnsToUpdate...), nil
}
