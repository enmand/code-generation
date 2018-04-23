package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/dave/jennifer/jen"
	"github.com/kr/pretty"
	"gopkg.in/yaml.v2"
)

const generatedQual = "github.com/enmand/code-generation/pets-go/generated"
const outputFolder = "./generated"

// The pets example shows the Pets example, normally covered in Object Oriented. This takes an
// input file defined as:
//
// 		dog:
//		  speaks: bark
//		  waklable: true
// 		cat:
//		  speaks: meow
//        walkable: false
//
// This will generate an Interface for Pets that includes Speak() and Walk() (which will error
// if the pet is not walkable
func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("you must include at least the spec for creating pets")
	}

	f, err := os.Open(args[1])
	if err != nil {
		log.Fatalf("unable to open spec file: %s", err)
	}

	spec, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("unable to read spec file: %s", err)
	}

	m := yaml.MapSlice{}
	yaml.Unmarshal(spec, &m)

	if err := os.Mkdir(outputFolder, 0744); err != nil && err.(*os.PathError).Err != syscall.Errno(0x11) {
		fmt.Printf("%# v", pretty.Formatter(err))
		log.Fatalf("unable to create generated package: %s", err)
	}
	f, err = os.OpenFile(fmt.Sprintf("%s/zz_generated_pets.go", outputFolder), os.O_CREATE|os.O_WRONLY, 0744)
	if err != nil {
		log.Fatalf("unable to create generated data file: %s", err)
	}

	genData := generateData(&m)
	if err := genData.Render(f); err != nil {
		log.Fatalf("unable to render generated data: %s", err)
	}

	genMain := generateMain(&m)
	if err := genMain.Save("active.go"); err != nil {
		log.Fatalf("unable to save main generated file: %s", err)
	}
}

// generate main generates the main file for bin/activity
func generateMain(s *yaml.MapSlice) *jen.File {
	f := jen.NewFile("main")

	f.Func().Id("main").Params().BlockFunc(func(g *jen.Group) {
		g.Var().Id("speaker").String()

		for _, p := range *s {
			sName := strings.Title(p.Key.(string))
			vName := strings.ToLower(p.Key.(string))

			g.Id(vName).Op(":=").Qual(generatedQual, fmt.Sprintf("New%s", sName)).Call()

			g.Id("speaker").Op("=").Id(vName).Dot("Speak").Call()

			g.Qual("fmt", "Printf").Call(jen.Lit(fmt.Sprintf("a %s says: ", sName)+"%s\n"), jen.Id("speaker"))

			g.If(
				jen.Err().Op(":=").Id(vName).Dot("Walk").Call(),
				jen.Err().Op("!=").Nil(),
			).Block(
				jen.Qual("fmt", "Println").Call(jen.Lit(fmt.Sprintf("%s are walkable!", sName))),
			).Else().Block(
				jen.Qual("fmt", "Println").Call(jen.Lit(fmt.Sprintf("%s are not walkable!", sName))),
			)
		}

		for _, p := range *s {
			sName := strings.Title(p.Key.(string))
			vName := strings.ToLower(p.Key.(string))
			petName := fmt.Sprintf("pet%s", sName)

			g.Id(fmt.Sprintf("pet%s", sName)).Op(":=").Id("IdentPet").Call(jen.Id(vName))
			g.Qual("fmt", "Printf").Call(jen.Lit(fmt.Sprintf("pet %s says: ", sName)+"%s\n"), jen.Id(petName).Dot("Speak").Call())
		}
	})

	f.Func().Id("IdentPet").Params(jen.Id("p").Qual(generatedQual, "Pet")).Qual(generatedQual, "Pet").Block(
		jen.Return(jen.Id("p")),
	)

	return f
}

// generate generates the Go code for the provided YAML slice of pet information
func generateData(s *yaml.MapSlice) *jen.File {
	f := jen.NewFile("generated")
	f.Type().Id("Pet").Interface(
		jen.Id("Speak").Params().String(),
		jen.Id("Walk").Params().Error(),
	)

	for _, pet := range *s {
		sName := strings.Title(pet.Key.(string))

		f.Type().Id(sName).Struct(
			jen.Id("speaks").String(),
			jen.Id("walkable").Bool(),
		)

		// Create News for struct objects
		f.Func().Id(fmt.Sprintf("New%s", sName)).Params().Op("*").Id(sName).Block(
			jen.Return(
				jen.Op("&").Id(sName).Values(jen.DictFunc(func(d jen.Dict) {
					for _, v := range pet.Value.(yaml.MapSlice) {
						d[jen.Id(v.Key.(string))] = jen.Lit(v.Value)
					}
				})),
			),
		)

		// Speak()
		f.Func().Params(
			jen.Id("p").Op("*").Id(sName),
		).Id("Speak").Params().String().Block(
			jen.Return(jen.Id("p").Dot("speaks")),
		)

		// Walk
		f.Func().Params(
			jen.Id("p").Op("*").Id(sName),
		).Id("Walk").Params().Error().Block(
			jen.If(
				jen.Id("p").Dot("walkable").Op("==").Lit(true),
			).Block(
				jen.Return(jen.Qual("errors", "New").Call(jen.Lit("unable to walk pet"))),
			),

			jen.Return(jen.Nil()),
		)
	}

	return f
}
