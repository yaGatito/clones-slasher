package formatter

import (
	"cloneslasher/internal/app"
	"cloneslasher/internal/domain"
	"cloneslasher/pkg/slicex"
)

type ItemClonesDTO struct {
	Item   ItemDTO  `json:"item"`
	Clones []string `json:"clones"`
}

type ItemNamesakesDTO struct {
	Name      string        `json:"name"`
	Namesakes []NamesakeDTO `json:"namesakes"`
}

type ItemDTO struct {
	Path      string   `json:"path"`
	Name      string   `json:"name"`
	Size      int64    `json:"size"`
	IsFolder  bool     `json:"isFolder"`
	Extension string   `json:"ext,omitempty"`
	Content   []string `json:"contentPaths,omitempty"`
}

type NamesakeDTO struct {
	Path      string   `json:"path"`
	Size      int64    `json:"size"`
	Extension string   `json:"ext,omitempty"`
	Content   []string `json:"contentPaths,omitempty"`
}

func MapItemNamesakesToDTO(data app.ItemNamesakes) ItemNamesakesDTO {
	dtoNamesakes := slicex.Map(data.Namesakes, mapToNamesakeDTO)
	return ItemNamesakesDTO{
		Name:      string(data.Name),
		Namesakes: dtoNamesakes,
	}
}

func MapItemClonesToDTO(itemClones app.ItemClones) ItemClonesDTO {
	clonesPaths := slicex.Map(itemClones.Clones, mapItemToString)
	return ItemClonesDTO{
		Item:   mapToItemDTO(itemClones.Item),
		Clones: clonesPaths,
	}
}

func mapToNamesakeDTO(item domain.Item) NamesakeDTO {
	var contentIDs []string
	if len(item.Content) > 0 {
		contentIDs = slicex.Map(item.Content, mapItemIDToString)
	} else {
		contentIDs = nil
	}

	return NamesakeDTO{
		Path:      item.ItemID.UniquePath,
		Size:      item.ItemID.Size,
		Extension: item.ItemID.Extension,
		Content:   contentIDs,
	}
}

func mapToItemDTO(item domain.Item) ItemDTO {
	var contentIDs []string
	if len(item.Content) > 0 {
		contentIDs = slicex.Map(item.Content, mapItemIDToString)
	} else {
		contentIDs = nil
	}

	return ItemDTO{
		Path:      string(item.ItemID.UniquePath),
		Name:      string(item.ItemID.Name),
		Size:      item.ItemID.Size,
		Extension: item.ItemID.Extension,
		IsFolder:  item.ItemID.IsFolder,
		Content:   contentIDs,
	}
}

func mapItemIDToString(itemID domain.ItemID) string {
	return itemID.UniquePath
}

func mapItemToString(item domain.Item) string {
	return item.ItemID.UniquePath
}
