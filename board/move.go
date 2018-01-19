package board

type Move struct {
	SourceRow    int
	SourceColumn int
	DestRow      int
	DestColumn   int
	CastTo       int
}
