package step

import (
	"encoding/xml"
	"fmt"
	"github.com/szumel/rusher/internal/platform/schema"
	"github.com/szumel/rusher/internal/step/macro"
	"log"
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
			<macro source="macro.xml" v="1.0.0"/>
       </sequence>
   </config>
</configPool>
`

var currentConfig *schema.Config

func init() {
	cpool, err := schema.NewFromString(schemaDefinition)
	if err != nil {
		log.Fatal(err)
	}

	c, err := schema.GetCurrentConfig(cpool, "test")
	if err != nil {
		log.Fatal(err)
	}

	currentConfig = c
}

func TestExtractSteps(t *testing.T) {
	expected := []stepCtx{
		{step: &PrintPwd{}, ctx: &ContextImpl{}},
		{step: &PrintPwd{}, ctx: &ContextImpl{}},
		{step: &ChangeCwd{}, ctx: &ContextImpl{}},
	}

	r := rusher{StepsPool, currentConfig, &loaderMock{version: "1.0.0"}}

	scs, err := r.extract()
	if err != nil {
		t.Fatal(err)
	}

	if len(scs) != len(expected) {
		t.Fatalf("rusher.extract failed: expected slice length %d, got %d", len(expected), len(scs))
	}

	for index, sc := range scs {
		fmt.Println(sc.step.Code(), expected[index].step.Code())
		if sc.step.Code() != expected[index].step.Code() {
			t.Fatalf("rusher.extract failed: expected %s in sequence, got %s", expected[index], sc.step.Code())
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

func TestWrongVersion(t *testing.T) {
	r := rusher{StepsPool, currentConfig, &loaderMock{version: "1.1.0"}}
	_, err := r.extract()
	_, ok := err.(*VersionNotFoundError)
	if !ok {
		t.Fatal("step.extrac faild: expected VersionNotFoundError")
	}
}

type loaderMock struct {
	version string
}

func (l *loaderMock) Load(source string) (macro.Schema, error) {
	ms := "<?xml version=\"1.0\"?><macro v=\"" + l.version + "\"><step code=\"printPwd\"/><step code=\"changeCwd\"/></macro>"

	return macro.Schema(ms), nil
}
