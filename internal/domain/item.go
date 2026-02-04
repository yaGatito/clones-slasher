package domain

type ItemID string

type ItemName string

type Item struct {
	ID        ItemID
	Name      ItemName
	Size      int64
	Extension string
	IsFolder  bool
	Content   []ItemID
}

func NewItem(path ItemID, name ItemName, extension string, isFolder bool, size int64) *Item {
	item := Item{
		ID:        path,
		Name:      name,
		Extension: extension,
		IsFolder:  isFolder,
		Size:      size,
	}
	if isFolder {
		item.Content = make([]ItemID, 0)
	} else {
		item.Content = nil
	}

	return &item
}

func (i Item) Equals(other Item) bool {
	return i.Name == other.Name && i.Size == other.Size &&
		i.Extension == other.Extension && i.IsFolder == other.IsFolder
}

func (i Item) Same(other Item) bool {
	return i.Name == other.Name
}

func getNameFromPath(path string) string {
	lastSlash := -1
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' || path[i] == '\\' {
			lastSlash = i
			break
		}
	}
	return path[lastSlash+1:]
}
