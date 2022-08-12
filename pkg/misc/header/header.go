package header

// CTX context
type CTX interface {
	GetHeader(key string) string
}
