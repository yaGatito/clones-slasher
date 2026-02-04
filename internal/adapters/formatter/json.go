package formatter

import (
	"cloneslasher/internal/app"
	"cloneslasher/internal/domain"
)

type ItemClonesDTO struct {
	Item   ItemDTO   `json:"item"`
	Clones []ItemDTO `json:"clones"`
}

type ItemNamesakesDTO struct {
	Name      string    `json:"name"`
	Namesakes []ItemDTO `json:"namesakes"`
}

type ItemDTO struct {
	Path      string   `json:"path"`
	Name      string   `json:"name"`
	Size      int64    `json:"size"`
	IsFolder  bool     `json:"isFolder"`
	Extension string   `json:"ext,omitempty"`
	Content   []string `json:"contentPaths,omitempty"`
}

func MapItemNamesakesToDTO(data []app.ItemNamesakes) []ItemNamesakesDTO {
	res := make([]ItemNamesakesDTO, len(data))

	for i, itemNamesakes := range data {
		dtoNamesakes := make([]ItemDTO, len(itemNamesakes.Namesakes))

		for j, item := range itemNamesakes.Namesakes {
			dtoNamesakes[j] = MapToDTO(item)
		}

		res[i] = ItemNamesakesDTO{
			Name:      string(itemNamesakes.Name),
			Namesakes: dtoNamesakes,
		}
	}
	return res
}

func MapItemClonesToDTO(data []app.ItemClones) []ItemClonesDTO {
	res := make([]ItemClonesDTO, len(data))

	for i, itemClones := range data {
		clonesDTO := make([]ItemDTO, len(itemClones.Clones))

		for j, item := range itemClones.Clones {
			clonesDTO[j] = MapToDTO(item)
		}

		res[i] = ItemClonesDTO{
			Item:   MapToDTO(itemClones.Item),
			Clones: clonesDTO,
		}
	}
	return res
}

func MapToDTO(item domain.Item) ItemDTO {
	contentIDs := make([]string, len(item.Content))
	for i, it := range item.Content {
		contentIDs[i] = string(it)
	}

	return ItemDTO{
		Path:      string(item.ID),
		Name:      string(item.Name),
		Size:      item.Size,
		Extension: item.Extension,
		IsFolder:  item.IsFolder,
		Content:   contentIDs,
	}
}
