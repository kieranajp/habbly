package conjugator

import (
	"github.com/kieranajp/habbly/ent"
	"github.com/kieranajp/habbly/ent/conjugation"
)

type PresentConjugationStrategy struct {
	verb *ent.Verb
}

func (s *PresentConjugationStrategy) Conjugate(verb *ent.Verb) map[conjugation.Subject]string {
	switch verb.Ending {
	case "ar":
		return map[conjugation.Subject]string{
			conjugation.SubjectYo:       verb.Stem + "o",
			conjugation.SubjectTu:       verb.Stem + "as",
			conjugation.SubjectEl:       verb.Stem + "a",
			conjugation.SubjectNosotros: verb.Stem + "amos",
			conjugation.SubjectVosotros: verb.Stem + "ais",
			conjugation.SubjectEllos:    verb.Stem + "an",
		}
	case "er":
		return map[conjugation.Subject]string{
			conjugation.SubjectYo:       verb.Stem + "o",
			conjugation.SubjectTu:       verb.Stem + "es",
			conjugation.SubjectEl:       verb.Stem + "e",
			conjugation.SubjectNosotros: verb.Stem + "emos",
			conjugation.SubjectVosotros: verb.Stem + "éis",
			conjugation.SubjectEllos:    verb.Stem + "en",
		}
	case "ir":
		return map[conjugation.Subject]string{
			conjugation.SubjectYo:       verb.Stem + "o",
			conjugation.SubjectTu:       verb.Stem + "es",
			conjugation.SubjectEl:       verb.Stem + "e",
			conjugation.SubjectNosotros: verb.Stem + "imos",
			conjugation.SubjectVosotros: verb.Stem + "ís",
			conjugation.SubjectEllos:    verb.Stem + "en",
		}
	}

	return nil
}
