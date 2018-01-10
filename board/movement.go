package board

type Movement struct {
	SourceRow    int
	SourceColumn int
	DestRow      int
	DestColumn   int
	CastTo       int
}
