package domain

type Item struct {
	Path      string
	Name      string
	Size      int64
	Extension string
	IsFolder  bool
	Content   []Item
}

func NewFile(path, name string, size int64, extension string) *Item {
	return &Item{
		Path:      path,
		Name:      name,
		Extension: extension,
		IsFolder:  false,
		Size:      size,
	}
}

func NewFolder(path, name string, size int64, ptrs []Item) *Item {
	return &Item{
		Path:      path,
		Name:      name,
		Extension: "",
		IsFolder:  true,
		Size:      size,
		Content:   ptrs,
	}
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
