package lists

import "github.com/pinem/server/models"

func GetSimpleLists(lists []models.List) []*models.SimpleList {
	var simpleLists []*models.SimpleList
	for _, list := range lists {
		simpleLists = append(simpleLists, GetSimpleList(&list))
	}
	return simpleLists
}

func GetSimpleList(list *models.List) *models.SimpleList {
	return &models.SimpleList{List: list}
}
