package helpers

import (
	"slices"
	"strings"
)

const (
	SORT_TITLE_ASC        = "title_asc"
	SORT_TITLE_DESC       = "title_desc"
	SORT_BEST_BEFORE_ASC  = "bestbefore_asc"
	SORT_BEST_BEFORE_DESC = "bestbefore_desc"
	SORT_PRICE_ASC        = "price_asc"
	SORT_PRICE_DESC       = "price_desc"
	SORT_DEFAULT          = "created_at asc"
	SEARCH_TYPE_AND       = "and"
	SEARCH_TYPE_OR        = "or"
	SEARCH_TYPE_DEFAULT   = SEARCH_TYPE_AND
)

func sortWhiteList() map[string]string {
	return map[string]string{
		SORT_TITLE_ASC:        "title ASC",
		SORT_TITLE_DESC:       "title DESC",
		SORT_BEST_BEFORE_ASC:  "best_before ASC",
		SORT_BEST_BEFORE_DESC: "best_before DESC",
		SORT_PRICE_ASC:        "price ASC",
		SORT_PRICE_DESC:       "price DESC",
	}
}

func SearchQuery(params map[string]string) (whereCond string, bindParams []any) {
	var tmpWhere []string
	if params["color"] != "" {
		tmpWhere = append(tmpWhere, "color = ?")
		bindParams = append(bindParams, params["color"])
	}
	if params["price_min"] != "" {
		tmpWhere = append(tmpWhere, "price >= ?")
		bindParams = append(bindParams, params["price_min"])
	}
	if params["price_max"] != "" {
		tmpWhere = append(tmpWhere, "price <= ?")
		bindParams = append(bindParams, params["price_max"])
	}
	if params["title"] != "" {
		tmpWhere = append(tmpWhere, "title = ?")
		bindParams = append(bindParams, params["title"])
	}
	if params["manufacturer"] != "" {
		tmpWhere = append(tmpWhere, "manufacturer = ?")
		bindParams = append(bindParams, params["manufacturer"])
	}
	searchType := SEARCH_TYPE_DEFAULT
	if params["search_type"] != "" && slices.Contains([]string{SEARCH_TYPE_AND, SEARCH_TYPE_OR}, params["search_type"]) {
		searchType = params["search_type"]
	}
	whereCond = strings.Join(tmpWhere, " "+searchType+" ")
	return
}

func OrderQuery(params map[string]string) string {
	if params["sort"] == "" {
		return SORT_DEFAULT
	}
	parts := strings.Split(params["sort"], ",")

	var whiteParts []string
	whiteList := sortWhiteList()
	for _, v := range parts {
		if whiteList[v] != "" {
			whiteParts = append(whiteParts, whiteList[v])
		}
	}

	return strings.Join(whiteParts, ", ")
}
