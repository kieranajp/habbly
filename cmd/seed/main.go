package main

import (
	"context"
	"fmt"
	"github.com/kieranajp/habbly/ent/conjugation"
	"github.com/kieranajp/habbly/pkg/conjugator"
	"log"

	"github.com/kieranajp/habbly/ent"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=localhost port=55432 user=habbly dbname=habbly password=habbly sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Seed the database
	if err := seedVerbs(client); err != nil {
		log.Fatalf("failed seeding the database: %v", err)
	}

	log.Println("Successfully seeded the database")
}

func seedVerbs(client *ent.Client) error {
	verbs, err := client.Verb.CreateBulk(
		client.Verb.Create().SetSpanish("ser").SetEnglish("to be").SetStem("s").SetEnding("er").SetIsStemchanger(false).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("estar").SetEnglish("to be").SetStem("est").SetEnding("ar").SetIsStemchanger(false).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("tener").SetEnglish("to have").SetStem("ten").SetEnding("er").SetIsStemchanger(true).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("hacer").SetEnglish("to do").SetStem("hac").SetEnding("er").SetIsStemchanger(false).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("ir").SetEnglish("to go").SetStem("v").SetEnding("ir").SetIsStemchanger(false).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("ver").SetEnglish("to see").SetStem("v").SetEnding("er").SetIsStemchanger(false).SetIsIrregular(false),
		client.Verb.Create().SetSpanish("poder").SetEnglish("to be able").SetStem("pod").SetEnding("er").SetIsStemchanger(true).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("dar").SetEnglish("to give").SetStem("d").SetEnding("ar").SetIsStemchanger(false).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("saber").SetEnglish("to know").SetStem("sab").SetEnding("er").SetIsStemchanger(false).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("poner").SetEnglish("to put").SetStem("pon").SetEnding("er").SetIsStemchanger(false).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("pedir").SetEnglish("to request").SetStem("ped").SetEnding("ir").SetIsStemchanger(true).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("seguir").SetEnglish("to follow").SetStem("segu").SetEnding("ir").SetIsStemchanger(true).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("servir").SetEnglish("to serve").SetStem("serv").SetEnding("ir").SetIsStemchanger(true).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("elegir").SetEnglish("to choose").SetStem("eleg").SetEnding("ir").SetIsStemchanger(true).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("caer").SetEnglish("to fall").SetStem("ca").SetEnding("er").SetIsStemchanger(false).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("venir").SetEnglish("to come").SetStem("ven").SetEnding("ir").SetIsStemchanger(true).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("traer").SetEnglish("to bring").SetStem("tra").SetEnding("er").SetIsStemchanger(false).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("haber").SetEnglish("to have").SetStem("hab").SetEnding("er").SetIsStemchanger(false).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("hablar").SetEnglish("to speak").SetStem("habl").SetEnding("ar").SetIsStemchanger(false).SetIsIrregular(false),
		client.Verb.Create().SetSpanish("mirar").SetEnglish("to look").SetStem("mir").SetEnding("ar").SetIsStemchanger(false).SetIsIrregular(false),
		client.Verb.Create().SetSpanish("escuchar").SetEnglish("to listen").SetStem("escuch").SetEnding("ar").SetIsStemchanger(false).SetIsIrregular(false),
		client.Verb.Create().SetSpanish("ayudar").SetEnglish("to help").SetStem("ayud").SetEnding("ar").SetIsStemchanger(false).SetIsIrregular(false),
		client.Verb.Create().SetSpanish("comer").SetEnglish("to eat").SetStem("com").SetEnding("er").SetIsStemchanger(false).SetIsIrregular(false),
		client.Verb.Create().SetSpanish("aprender").SetEnglish("to learn").SetStem("aprend").SetEnding("er").SetIsStemchanger(false).SetIsIrregular(false),
		client.Verb.Create().SetSpanish("deber").SetEnglish("must").SetStem("deb").SetEnding("er").SetIsStemchanger(false).SetIsIrregular(false),
		client.Verb.Create().SetSpanish("leer").SetEnglish("to read").SetStem("le").SetEnding("er").SetIsStemchanger(false).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("vivir").SetEnglish("to live").SetStem("viv").SetEnding("ir").SetIsStemchanger(false).SetIsIrregular(false),
		client.Verb.Create().SetSpanish("permitir").SetEnglish("to allow").SetStem("permit").SetEnding("ir").SetIsStemchanger(false).SetIsIrregular(false),
		client.Verb.Create().SetSpanish("recibir").SetEnglish("to receive").SetStem("recib").SetEnding("ir").SetIsStemchanger(false).SetIsIrregular(false),
		client.Verb.Create().SetSpanish("abrir").SetEnglish("to open").SetStem("abr").SetEnding("ir").SetIsStemchanger(false).SetIsIrregular(false),
		client.Verb.Create().SetSpanish("alcanzar").SetEnglish("to reach").SetStem("alcanz").SetEnding("ar").SetIsStemchanger(false).SetIsIrregular(false),
		client.Verb.Create().SetSpanish("llegar").SetEnglish("to arrive").SetStem("lleg").SetEnding("ar").SetIsStemchanger(false).SetIsIrregular(false),
		client.Verb.Create().SetSpanish("buscar").SetEnglish("to search").SetStem("busc").SetEnding("ar").SetIsStemchanger(false).SetIsIrregular(false),
		client.Verb.Create().SetSpanish("perecer").SetEnglish("to perish").SetStem("perec").SetEnding("er").SetIsStemchanger(false).SetIsIrregular(false),
		client.Verb.Create().SetSpanish("conocer").SetEnglish("to know").SetStem("conoc").SetEnding("er").SetIsStemchanger(false).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("querer").SetEnglish("to want").SetStem("quer").SetEnding("er").SetIsStemchanger(true).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("dirigir").SetEnglish("to direct").SetStem("dirig").SetEnding("ir").SetIsStemchanger(false).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("construir").SetEnglish("to build").SetStem("constru").SetEnding("ir").SetIsStemchanger(false).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("exigir").SetEnglish("to demand").SetStem("exig").SetEnding("ir").SetIsStemchanger(false).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("jugar").SetEnglish("to play").SetStem("jug").SetEnding("ar").SetIsStemchanger(true).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("pensar").SetEnglish("to think").SetStem("pens").SetEnding("ar").SetIsStemchanger(true).SetIsIrregular(false),
		client.Verb.Create().SetSpanish("empezar").SetEnglish("to start").SetStem("empez").SetEnding("ar").SetIsStemchanger(true).SetIsIrregular(false),
		client.Verb.Create().SetSpanish("entender").SetEnglish("to understand").SetStem("entend").SetEnding("er").SetIsStemchanger(true).SetIsIrregular(false),
		client.Verb.Create().SetSpanish("volver").SetEnglish("to return").SetStem("volv").SetEnding("er").SetIsStemchanger(true).SetIsIrregular(false),
		client.Verb.Create().SetSpanish("doler").SetEnglish("to hurt").SetStem("dol").SetEnding("er").SetIsStemchanger(true).SetIsIrregular(false),
		client.Verb.Create().SetSpanish("sentir").SetEnglish("to feel").SetStem("sent").SetEnding("ir").SetIsStemchanger(true).SetIsIrregular(false),
		client.Verb.Create().SetSpanish("morir").SetEnglish("to die").SetStem("mor").SetEnding("ir").SetIsStemchanger(true).SetIsIrregular(true),
		client.Verb.Create().SetSpanish("conseguir").SetEnglish("to achieve").SetStem("consig").SetEnding("ir").SetIsStemchanger(true).SetIsIrregular(false),
	).Save(context.Background())

	if err != nil {
		return err
	}

	// Conjugate the seeded verbs
	for _, verb := range verbs {
		c, err := conjugator.NewConjugation(verb, conjugation.TensePresent)
		if err != nil {
			return fmt.Errorf("failed to find conjugation strategy: %w", err)
		}

		verbTable := c.Conjugate(verb)
		for subject, conjugated := range verbTable {
			_ = client.Conjugation.Create().
				SetTense(conjugation.TensePresent).
				SetSubject(subject).
				SetVerbID(verb.ID).
				SetAnswer(conjugated).
				SaveX(context.Background())
		}

	}

	return nil
}
