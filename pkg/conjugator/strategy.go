package conjugator

import (
	"fmt"
	"github.com/kieranajp/habbly/ent"
	"github.com/kieranajp/habbly/ent/conjugation"
)

type ConjugationStrategy interface {
	// Conjugate the verb
	Conjugate(verb *ent.Verb) map[conjugation.Subject]string
}

func NewConjugation(verb *ent.Verb, tense conjugation.Tense) (ConjugationStrategy, error) {
	switch tense {
	case conjugation.TensePresent:
		if verb.IsStemchanger || verb.IsIrregular {
			return &PresentIrregularConjugationStrategy{
				verb: verb,
			}, nil
		}
		return &PresentConjugationStrategy{
			verb: verb,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported tense: %s", tense)
	}
}
