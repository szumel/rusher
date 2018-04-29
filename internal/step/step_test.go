package step

import (
	"encoding/xml"
	"fmt"
	"github.com/szumel/rusher/internal/platform/schema"
	"log"
	"os"
	"testing"
)

var schemaDefinition = `
<?xml version="1.0"?>

<configPool>
   <config>
       <environment>test</environment>
       <projectPath>/home/projects/releaser/src</projectPath>
       <globals>
           <var name="php" value="/usr/bin/php"/>
       </globals>
       <!-- Execute 'rusher listSteps' for all available steps listing -->
       <sequence>
			<step code="printPwd" />
			<macro source="` + macroPath() + `" v="1.0.0"/>
       </sequence>
   </config>
</configPool>
`

func macroPath() string {
	wd, _ := os.Getwd()
	return wd + "/macro.xml"
}

//@todo mock source? change path?
const (
	macroDefinition = `
	<macro v="1.0.0">
		<step code="printPwd" />
		<step code="changeCwd" dir="/"/>
	</macro>
`
)

func TestExtractSteps(t *testing.T) {
	expected := []stepCtx{
		{step: &PrintPwd{}, ctx: &ContextImpl{}},
		{step: &PrintPwd{}, ctx: &ContextImpl{}},
		{step: &ChangeCwd{}, ctx: &ContextImpl{}},
	}

	cpool, err := schema.NewFromString(schemaDefinition)
	if err != nil {
		t.Fatal(err)
	}

	c, err := schema.GetCurrentConfig(cpool, "test")
	if err != nil {
		log.Fatal(err)
	}

	r := rusher{StepsPool, c}

	scs, err := r.extract()
	if err != nil {
		log.Fatal(err)
	}

	if len(scs) != len(expected) {
		log.Fatalf("rusher.extract failed: expected slice length %d, got %d", len(expected), len(scs))
	}

	for index, sc := range scs {
		fmt.Println(sc.step.Code(), expected[index].step.Code())
		if sc.step.Code() != expected[index].step.Code() {
			log.Fatalf("rusher.extract failed: expected %s in sequence, got %s", expected[index], sc.step.Code())
		}
	}
}

func TestConvertSequenceElemToMacro(t *testing.T) {
	seqElem := schema.SequenceElem{
		XMLName: xml.Name{
			Local: "macro"},
		Params: []xml.Attr{
			{Name: xml.Name{Local: "v"}, Value: "1.0.0"},
			{Name: xml.Name{Local: "source"}, Value: "macro.xml"},
		},
	}

	macroElem := toMacro(seqElem)

	if macroElem.Version != "1.0.0" {
		log.Fatalf("step.toMacro failed: Version expected %s got %s", "1.0.0", macroElem.Version)
	}

	if macroElem.Source != "macro.xml" {
		log.Fatalf("step.toMacro failed: Source expected %s got %s", "macro.xml", macroElem.Source)
	}
}
