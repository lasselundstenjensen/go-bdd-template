package main

import "github.com/cucumber/godog"

func iSubtractTheNumber(arg1 int) error {
	return godog.ErrPending
}

func theNumberShouldBeMinus(arg1 int) error {
	return godog.ErrPending
}

func InitializeSubtractScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I subtract the number (\d+)$`, iSubtractTheNumber)
	ctx.Step(`^the number should be minus (\d+)$`, theNumberShouldBeMinus)
}
