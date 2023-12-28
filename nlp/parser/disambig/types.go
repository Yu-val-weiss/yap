package disambig

import (
	. "yu-val-weiss/yap/nlp/types"
)

type MorphologicalDisambiguator interface {
	Parse(LatticeSentence) (Mappings, interface{})
}
