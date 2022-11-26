package types

type Bookmark struct {
	Name     string     `json:"name"`
	Type     string     `json:"type"`
	Url      string     `json:"url"`
	Children []Bookmark `json:"children"`
}

type Folder struct {
	Name     string     `json:"name"`
	Children []Bookmark `json:"children"`
	Type     string     `json:"type"`
}

type Roots struct {
	BookmarkBar Folder `json:"bookmark_bar"`
}

type Data struct {
	Roots Roots `json:"roots"`
}
