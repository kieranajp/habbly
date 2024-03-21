package conjugator

import (
	"github.com/kieranajp/habbly/ent"
	"github.com/kieranajp/habbly/ent/conjugation"
)

type PreteriteConjugationStrategy struct {
	verb *ent.Verb
}

func (s *PreteriteConjugationStrategy) Conjugate(verb *ent.Verb) map[conjugation.Subject]string {
	switch verb.Ending {
	case "ar":
		return map[conjugation.Subject]string{
			conjugation.SubjectYo:       verb.Stem + "é",
			conjugation.SubjectTu:       verb.Stem + "aste",
			conjugation.SubjectEl:       verb.Stem + "ó",
			conjugation.SubjectNosotros: verb.Stem + "amos",
			conjugation.SubjectVosotros: verb.Stem + "asteis",
			conjugation.SubjectEllos:    verb.Stem + "aron",
		}
	case "er":
	case "ir":
		return map[conjugation.Subject]string{
			conjugation.SubjectYo:       verb.Stem + "í",
			conjugation.SubjectTu:       verb.Stem + "iste",
			conjugation.SubjectEl:       verb.Stem + "ió",
			conjugation.SubjectNosotros: verb.Stem + "imos",
			conjugation.SubjectVosotros: verb.Stem + "isteis",
			conjugation.SubjectEllos:    verb.Stem + "ieron",
		}
	}

	return nil
}
