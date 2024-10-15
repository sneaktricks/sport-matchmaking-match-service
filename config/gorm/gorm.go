package main

import (
	"github.com/sneaktricks/sport-matchmaking-match-service/model"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "dal",
		Mode:          gen.WithDefaultQuery,
		FieldNullable: true,
	})

	g.ApplyBasic(model.Match{}, model.Participation{})

	g.Execute()
}
