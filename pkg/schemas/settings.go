package interfaces

type MultiLangSchema struct {
	Value string
	Lang  string
}

type SuggestedLocations string

var SuggestedLocationMap = map[SuggestedLocations]struct{}{
	"hubs": {}, "warehouse": {}, "island": {}, "zones": {}, "shelf": {},
	"platform": {}, "shelfSection": {}, "shelfFloor": {}, "palet": {},
	"drawer": {}, "containerBox": {},
}

func IsValidSuggestedLocation(value SuggestedLocations) bool {
	_, exists := SuggestedLocationMap[value]
	return exists
}
