package ma

import . "yu-val-weiss/yap/nlp/types"

type MorphologicalAnalyzer interface {
	Analyze(input []string) (LatticeSentence, interface{})
}
