package shared

// for html source
type AttrInfo struct {
	Tag string
	Value string
	Line uint
}
type TagInfo struct {
	Attr []string
	Value []string
	Line uint
}
