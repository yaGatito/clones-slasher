package domain

type ItemNamesakes struct {
	Name      ItemName
	Namesakes []Item
}

type ItemClones struct {
	Item   Item
	Clones []Item
}

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

// TODO: WHAT WILL BE IF CHANGE STRUCT KEY OF THE MAP AFTER IT WAS PASSED IN MAP.
func NewItem(id string, name string, extension string, isFolder bool, size int64) *Item {
	if id == "" || name == "" {
		panic("invalid item creation")
	}
	item := Item{
		ItemID:  ItemID{
			UniquePath: id,
			Name:       ItemName(name),
			IsFolder:   isFolder,
			Size:       size,
		},
	}
	if isFolder {
		item.Content = make([]ItemID, 0)
	} else {
		item.ItemID.Extension = extension
	}

	return &item
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
