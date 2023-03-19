package main

import (
	"os"
    "github.com/the-singularity-labs/holocron/pkg/forge"
    
	"github.com/hoenirvili/skapt"
	"github.com/hoenirvili/skapt/argument"
	"github.com/hoenirvili/skapt/flag"
)

func main() {
	app := skapt.Application{
		Name:        "Forge",
		Description: "Forge a Holocron",
		Version:     "1.0.0",
		Handler: func(ctx *skapt.Context) error {
			name := ctx.String("name")
			prompt := ctx.String("gatekeeper")
			ascertainment := ctx.String("ascertainment")
			treasure := ctx.String("treasure")
			salt := ctx.String("salt")
			outdir := ctx.String("outdir")

			h := forge.NewHolocron(name, prompt, ascertainment, treasure, salt)
			err := h.Forge(outdir)
			return err
		},
		Flags: flag.Flags{{
			Short: "n", Long: "name",
			Description: "Name of the horcrux",
			Type:        argument.String,
			Required:	 true,
		}, {
			Short: "g", Long: "gatekeeper",
			Description: "Prompt that gatekeeps the treasure",
			Type:        argument.String,
			Required:	 true,
		}, {
			Short: "a", Long: "ascertainment",
			Description: "Answer to the prompt that proves worthiness to spoils of treasure",
			Type:        argument.String,
			Required:	 true,
		}, {
			Short: "t", Long: "treasure",
			Description: "Treasure guarded the gatekeeper",
			Type:        argument.String,
			Required:	 true,
		}, {
			Short: "s", Long: "salt",
			Description: "Salt for encryption",
			Type:        argument.String,
			Required:	 true,
		}, {
			Short: "o", Long: "outdir",
			Description: "Directory to output horcrux",
			Type:        argument.String,
			Required:	 true,
		}},
	}
	app.Exec(os.Args)

}
