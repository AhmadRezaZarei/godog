package godog_test

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/internal/parser"
	"github.com/cucumber/godog/internal/storage"
)

func iEat(arg1 int) error {
	return godog.ErrPending
}

func thereAreGodogs(arg1 int) error {
	return godog.ErrPending
}

func thereShouldBeRemaining(arg1 int) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^there are (\d+) godogs$`, thereAreGodogs)
	ctx.Step(`^I eat (\d+)$`, iEat)
	ctx.Step(`^there should be (\d+) remaining$`, thereShouldBeRemaining)
}

func getFunctionName(x interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(x).Pointer()).Name()
}

func TestInlineFeatures(t *testing.T) {
	mm := make(map[string]interface{})

	mm["a"] = func(x int) {
	}

	mm["b"] = func(x string) {
	}

	tt := reflect.TypeOf(InitializeScenario)

	fmt.Println("#")
	fmt.Println(tt.String())
	fmt.Println(tt.Name())

	// fmt.Println(name)
	fmt.Println("#")
	features, err := parser.ParseFeatures(storage.FS{}, "", "", []string{"features/inline.feature"})
	if err != nil {
		t.Fatal(err)
		return
	}

	feature := features[0]

	fmt.Println(len(feature.Pickles[0].Steps))

	// suite := godog.TestSuite{
	// 	ScenarioInitializer: InitializeScenario,
	// 	Options: &godog.Options{
	// 		Format:   "pretty",
	// 		Paths:    []string{"features/inline.feature"},
	// 		TestingT: t, // Testing instance that will run subtests.
	// 	},
	// }

	// if suite.Run() != 0 {
	// 	t.Fatal("non-zero status returned, failed to run feature tests")
	// }
}
