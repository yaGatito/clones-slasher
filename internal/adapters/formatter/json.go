package formatter

import (
	"bytes"
	"cloneslasher/internal/domain"
	"encoding/json"
)

type FolderClonesDTO struct {
	Name        string   `json:"name"`
	Path        string   `json:"path"`
	Size        string   `json:"size"`
	ClonesPaths []string `json:"clones,omitzero"`
}

type FolderNamesakesDTO struct {
	Name      string    `json:"name"`
	Namesakes []ItemDTO `json:"namesakes"`
}

type ItemDTO struct {
	Path      string   `json:"path"`
	Name      string   `json:"name"`
	Size      int64    `json:"size"`
	Extension string   `json:"ext"`
	IsFolder  bool     `json:"isFolder"`
	Content   []string `json:"contentPaths,omitzero"`
}

func MapCollectionToDTO(data map[string][]domain.Item) []FolderNamesakesDTO {
	res := make([]FolderNamesakesDTO, len(data))
	for k, v := range data {
		dtoClones := make([]ItemDTO, len(v))

		for _, cl := range v {
			dtoClones = append(dtoClones, MapToDTO(cl))
		}

		res = append(res, FolderNamesakesDTO{
			Name:      k,
			Namesakes: dtoClones,
		})
	}
	return res
}

func MapToDTO(item domain.Item) ItemDTO {
	contentPaths := make([]string, len(item.Content))
	for _, it := range item.Content {
		contentPaths = append(contentPaths, it.Path)
	}

	return ItemDTO{
		Path:      item.Path,
		Name:      item.Name,
		Size:      item.Size,
		Extension: item.Extension,
		IsFolder:  item.IsFolder,
		Content:   contentPaths,
	}
}

func MapToJson(data []FolderNamesakesDTO) (*bytes.Buffer, error) {
	buffer := bytes.NewBuffer(make([]byte, 0))
	err := json.NewEncoder(buffer).Encode(data)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
