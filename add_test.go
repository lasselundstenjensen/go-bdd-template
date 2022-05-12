package main

import (
	"context"

	"github.com/cucumber/godog"
)

func iAddTheNumber(b int) error {
	return godog.ErrPending
}

func theNumber(a int) error {
	return godog.ErrPending
}

func theNumberShouldBe(arg1 int) error {
	return godog.ErrPending
}

func InitializeTestSuite(sc *godog.TestSuiteContext) {
	sc.BeforeSuite(func() { /* Do setup logic here, if any */ })
}

func InitializeAddScenario(sc *godog.ScenarioContext) {
	sc.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		// Do things before each scenario is run, e.g. clean the state

		return ctx, nil
	})

	sc.Step(`^I add the number (\d+)$`, iAddTheNumber)
	sc.Step(`^the number (\d+)$`, theNumber)
	sc.Step(`^the number should be (\d+)$`, theNumberShouldBe)
}
