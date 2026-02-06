package domain

type ItemName string

type ItemID struct {
	UniquePath string
	Name       ItemName
	Size       int64
	Extension  string
	IsFolder   bool
}

type Item struct {
	ItemID  ItemID
	Content []ItemID
}

// TODO: WHAT WILL BE IF TO CHANGE STRUCT KEY OF THE MAP AFTER IT WILL BE PASSED.
func NewItem(id string, name string, extension string, isFolder bool, size int64) *Item {
	if id == "" || name == "" {
		panic("invalid item creation")
	}
	itemID := ItemID{
		UniquePath: id,
		Name:       ItemName(name),
		IsFolder:   isFolder,
		Size:       size,
	}
	var content []ItemID
	if isFolder {
		content = make([]ItemID, 0)
	} else {
		content = nil

		itemID.Extension = extension
	}

	return &Item{
		ItemID:  itemID,
		Content: content,
	}
}

func (i Item) Equal(other Item) bool {
	return i.ItemID == other.ItemID
}

func (i Item) IsClone(other Item) bool {
	return i.ItemID.IsClone(other.ItemID)
}

func (i ItemID) IsClone(other ItemID) bool {
	// Important part. i and other shouldn't have the same UniquePath.
	if i.UniquePath == other.UniquePath {
		return false
	}

	return i.Name == other.Name &&
		i.Size == other.Size && i.Extension == other.Extension &&
		i.IsFolder == other.IsFolder
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
