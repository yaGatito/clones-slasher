package domain

type ItemID string

type ItemName string

type Comparer[T any] interface {
	Equal(other T) bool
}

type Item struct {
	ID        ItemID
	Name      ItemName
	Size      int64
	Extension string
	IsFolder  bool
	Content   []ItemID
}

func NewItem(id ItemID, name ItemName, extension string, isFolder bool, size int64) *Item {
	if id == "" || name == "" {
		panic("invalid item creation")
	}
	item := Item{
		ID:       id,
		Name:     name,
		IsFolder: isFolder,
		Size:     size,
	}
	if isFolder {
		item.Content = make([]ItemID, 0)
	} else {
		item.Extension = extension
		item.Content = nil
	}

	return &item
}

func (i Item) Equal(other Item) bool {
	return i.ID == other.ID && i.Same(other)
}

func (i Item) Same(other Item) bool {
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
