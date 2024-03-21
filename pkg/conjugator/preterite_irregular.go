package conjugator

import (
	"github.com/kieranajp/habbly/ent"
	"github.com/kieranajp/habbly/ent/conjugation"
)

type PreteriteIrregularConjugationStrategy struct {
	verb *ent.Verb
}

func (s *PreteriteIrregularConjugationStrategy) Conjugate(verb *ent.Verb) map[conjugation.Subject]string {
	switch verb.Spanish {
	case "ser":
		return s.mapSubject([]string{"fui", "fuiste", "fue", "fuimos", "fuisteis", "fueron"})
	case "estar":
		return s.mapSubject([]string{"estuve", "estuviste", "estuvo", "estuvimos", "estuvisteis", "estuvieron"})
	case "tener":
		return s.mapSubject([]string{"tuve", "tuviste", "tuvo", "tuvimos", "tuvisteis", "tuvieron"})
	case "hacer":
		return s.mapSubject([]string{"hice", "hiciste", "hizo", "hicimos", "hicisteis", "hicieron"})
	case "ir":
		return s.mapSubject([]string{"fui", "fuiste", "fue", "fuimos", "fuisteis", "fueron"})
	case "poder":
		return s.mapSubject([]string{"pude", "pudiste", "pudo", "pudimos", "pudisteis", "pudieron"})
	case "dar":
		return s.mapSubject([]string{"di", "diste", "dio", "dimos", "disteis", "dieron"})
	case "saber":
		return s.mapSubject([]string{"supe", "supiste", "supo", "supimos", "supisteis", "supieron"})
	case "poner":
		return s.mapSubject([]string{"puse", "pusiste", "puso", "pusimos", "pusisteis", "pusieron"})
	case "pedir":
		return s.mapSubject([]string{"pedí", "pediste", "pidió", "pedimos", "pedisteis", "pidieron"})
	case "seguir":
		return s.mapSubject([]string{"seguí", "seguiste", "siguió", "seguimos", "seguisteis", "siguieron"})
	case "servir":
		return s.mapSubject([]string{"serví", "serviste", "sirvió", "servimos", "servisteis", "sirvieron"})
	case "elegir":
		return s.mapSubject([]string{"elegí", "elegiste", "eligió", "elegimos", "elegisteis", "eligieron"})
	case "caer":
		return s.mapSubject([]string{"caí", "caíste", "cayó", "caímos", "caísteis", "cayeron"})
	case "venir":
		return s.mapSubject([]string{"vine", "viniste", "vino", "vinimos", "vinisteis", "vinieron"})
	case "traer":
		return s.mapSubject([]string{"traje", "trajiste", "trajo", "trajimos", "trajisteis", "trajeron"})
	case "leer":
		return s.mapSubject([]string{"leí", "leíste", "leyó", "leímos", "leísteis", "leyeron"})
	case "conocer":
		return s.mapSubject([]string{"conocí", "conociste", "conoció", "conocimos", "conocisteis", "conocieron"})
	case "querer":
		return s.mapSubject([]string{"quise", "quisiste", "quiso", "quisimos", "quisisteis", "quisieron"})
	case "dirigir":
		return s.mapSubject([]string{"dirigí", "dirigiste", "dirigió", "dirigimos", "dirigisteis", "dirigieron"})
	case "construir":
		return s.mapSubject([]string{"construí", "construiste", "construyó", "construimos", "construisteis", "construyeron"})
	case "exigir":
		return s.mapSubject([]string{"exigí", "exigiste", "exigió", "exigimos", "exigisteis", "exigieron"})
	case "jugar":
		return s.mapSubject([]string{"jugué", "jugaste", "jugó", "jugamos", "jugasteis", "jugaron"})
	case "morir":
		return s.mapSubject([]string{"morí", "moriste", "murió", "morimos", "moristeis", "murieron"})
	case "conseguir":
		return s.mapSubject([]string{"conseguí", "conseguiste", "consiguió", "conseguimos", "conseguisteis", "consiguieron"})
	}

	return nil
}

func (s *PreteriteIrregularConjugationStrategy) mapSubject(c []string) map[conjugation.Subject]string {
	m := map[conjugation.Subject]string{}
	m[conjugation.SubjectYo] = c[0]
	m[conjugation.SubjectTu] = c[1]
	m[conjugation.SubjectEl] = c[2]
	m[conjugation.SubjectNosotros] = c[3]
	m[conjugation.SubjectVosotros] = c[4]
	m[conjugation.SubjectEllos] = c[5]

	return m
}
