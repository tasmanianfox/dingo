package board

type Movement struct {
	SourceHorizontal int
	SourceVertical   int
	DestHorizontal   int
	DestVertical     int
	CastTo           int
}
