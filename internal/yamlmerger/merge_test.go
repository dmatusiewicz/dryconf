package yamlmerger

import (
	"bytes"
	"embed"
	"strings"
	"testing"

	"github.com/dmatusiewicz/fragmenter"
	"gopkg.in/yaml.v2"
)

//go:embed tests
var tests embed.FS

type testCase struct {
	existingYAML []byte
	sourceYAML   []byte
	resultYAML   []byte
}

func TestMerge(t *testing.T) {

	tc := extractTestCasesFromFiles(tests, t)
	for _, tcc := range tc {
		var unmarshalledDataDst, unmarshalledDataSrc, unmarshalledDataResult map[string]interface{}
		err := yaml.Unmarshal([]byte(tcc.existingYAML), &unmarshalledDataDst)
		if err != nil {
			t.Fatal(err)
		}
		err = yaml.Unmarshal([]byte(tcc.sourceYAML), &unmarshalledDataSrc)
		if err != nil {
			t.Fatal(err)
		}
		err = yaml.Unmarshal([]byte(tcc.resultYAML), &unmarshalledDataResult)
		if err != nil {
			t.Fatal(err)
		}

		out, err := Merge(unmarshalledDataDst, unmarshalledDataSrc)
		if err != nil {
			t.Fatal(err)
		}
		outData, err := yaml.Marshal(out)
		if string(outData) != string(tcc.resultYAML) {
			t.Log("Have:")
			t.Logf("\n%s", outData)
			t.Log("Want:")
			t.Logf("\n%s", tcc.resultYAML)
			t.Logf("Result string dont match what we want.")
			t.Fail()
		}
	}

}

// extractTestCasesFromFiles loads all .yaml files from "tests" directory (subdirectories are ignored). Each file should contain 3 YAML documents:
// - first component,
// - second component,
// - expected merge result
func extractTestCasesFromFiles(f embed.FS, t *testing.T) []testCase {
	var testCases []testCase
	const testDir string = "tests"
	de, err := f.ReadDir(testDir)
	if err != nil {
		t.Fatalf(err.Error())
	}

	for _, j := range de {
		_, err := j.Info()
		if err != nil {
			t.Fatalf(err.Error())
		}

		if j.Type().IsRegular() && strings.HasSuffix(j.Name(), ".yaml") {
			// t.Logf("%s", j.Name())
			fileData, err := f.ReadFile(testDir + "/" + j.Name())
			if err != nil {
				t.Fatal(err)
			}

			searchBytes := []byte("---\n")
			data := fragmenter.Fragment(fileData, searchBytes)

			testCases = append(testCases, testCase{
				existingYAML: bytes.ReplaceAll(data[0], searchBytes, nil),
				sourceYAML:   bytes.ReplaceAll(data[1], searchBytes, nil),
				resultYAML:   bytes.ReplaceAll(data[2], searchBytes, nil),
			})
		}
	}
	return testCases
}
