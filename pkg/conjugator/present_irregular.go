package conjugator

import (
	"github.com/kieranajp/habbly/ent"
	"github.com/kieranajp/habbly/ent/conjugation"
)

type PresentIrregularConjugationStrategy struct {
	verb *ent.Verb
}

func (s *PresentIrregularConjugationStrategy) Conjugate(verb *ent.Verb) map[conjugation.Subject]string {
	switch verb.Spanish {
	case "ser":
		return s.mapSubject([]string{"soy", "eres", "es", "somos", "sois", "son"})
	case "estar":
		return s.mapSubject([]string{"estoy", "estás", "está", "estamos", "estáis", "están"})
	case "tener":
		return s.mapSubject([]string{"tengo", "tienes", "tiene", "tenemos", "tenéis", "tienen"})
	case "hacer":
		return s.mapSubject([]string{"hago", "haces", "hace", "hacemos", "hacéis", "hacen"})
	case "ir":
		return s.mapSubject([]string{"voy", "vas", "va", "vamos", "vais", "van"})
	case "poder":
		return s.mapSubject([]string{"puedo", "puedes", "puede", "podemos", "podéis", "pueden"})
	case "dar":
		return s.mapSubject([]string{"doy", "das", "da", "damos", "dais", "dan"})
	case "saber":
		return s.mapSubject([]string{"sé", "sabes", "sabe", "sabemos", "sabéis", "saben"})
	case "poner":
		return s.mapSubject([]string{"pongo", "pones", "pone", "ponemos", "ponéis", "ponen"})
	case "pedir":
		return s.mapSubject([]string{"pido", "pides", "pide", "pedimos", "pedís", "piden"})
	case "seguir":
		return s.mapSubject([]string{"sigo", "sigues", "sigue", "seguimos", "seguís", "siguen"})
	case "servir":
		return s.mapSubject([]string{"sirvo", "sirves", "sirve", "servimos", "servís", "sirven"})
	case "elegir":
		return s.mapSubject([]string{"elijo", "eliges", "elige", "elegimos", "elegís", "eligen"})
	case "caer":
		return s.mapSubject([]string{"caigo", "caes", "cae", "caemos", "caéis", "caen"})
	case "venir":
		return s.mapSubject([]string{"vengo", "vienes", "viene", "venimos", "venís", "vienen"})
	case "traer":
		return s.mapSubject([]string{"traigo", "traes", "trae", "traemos", "traéis", "traen"})
	case "haber":
		return s.mapSubject([]string{"he", "has", "ha", "hemos", "habéis", "han"})
	case "leer":
		return s.mapSubject([]string{"leo", "lees", "lee", "leemos", "leéis", "leen"})
	case "conocer":
		return s.mapSubject([]string{"conozco", "conoces", "conoce", "conocemos", "conocéis", "conocen"})
	case "querer":
		return s.mapSubject([]string{"quiero", "quieres", "quiere", "queremos", "queréis", "quieren"})
	case "dirigir":
		return s.mapSubject([]string{"dirijo", "diriges", "dirige", "dirigimos", "dirigís", "dirigen"})
	case "construir":
		return s.mapSubject([]string{"construyo", "construyes", "construye", "construimos", "construís", "construyen"})
	case "exigir":
		return s.mapSubject([]string{"exijo", "exiges", "exige", "exigimos", "exigís", "exigen"})
	case "jugar":
		return s.mapSubject([]string{"juego", "juegas", "juega", "jugamos", "jugáis", "juegan"})
	case "pensar":
		return s.mapSubject([]string{"pienso", "piensas", "piensa", "pensamos", "pensáis", "piensan"})
	case "empezar":
		return s.mapSubject([]string{"empiezo", "empiezas", "empieza", "empezamos", "empezáis", "empiezan"})
	case "entender":
		return s.mapSubject([]string{"entiendo", "entiendes", "entiende", "entendemos", "entendéis", "entienden"})
	case "volver":
		return s.mapSubject([]string{"vuelvo", "vuelves", "vuelve", "volvemos", "volvéis", "vuelven"})
	case "doler":
		return s.mapSubject([]string{"duelo", "dueles", "duele", "dolemos", "doléis", "duelen"})
	case "sentir":
		return s.mapSubject([]string{"siento", "sientes", "siente", "sentimos", "sentís", "sienten"})
	case "morir":
		return s.mapSubject([]string{"muero", "mueres", "muere", "morimos", "morís", "mueren"})
	case "conseguir":
		return s.mapSubject([]string{"consigo", "consigues", "consigue", "conseguimos", "conseguís", "consiguen"})
	}

	return nil
}

func (s *PresentIrregularConjugationStrategy) mapSubject(c []string) map[conjugation.Subject]string {
	m := map[conjugation.Subject]string{}
	m[conjugation.SubjectYo] = c[0]
	m[conjugation.SubjectTu] = c[1]
	m[conjugation.SubjectEl] = c[2]
	m[conjugation.SubjectNosotros] = c[3]
	m[conjugation.SubjectVosotros] = c[4]
	m[conjugation.SubjectEllos] = c[5]

	return m
}
