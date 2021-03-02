package statistic

type StatisticType int

const (
	VAOBinds = StatisticType(iota)
	VBOBinds
	IBOBinds
	FBOBinds
	DrawCalls
	VerticesDrawn
	VertexUpload
	SpritesDrawn
)
