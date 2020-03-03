/*
Feature: Compute required fuel
  In order to travel around the world on a rocket ship
  As a pilot I need to be able to compute the rocket fuel required for the trip

  Scenario: Compute the fuel cost                            # features\solver.feature:7
    Given a file containing a list of masses
    When I compute the fuel cost for all modules in the file
    Then the fuel cost should be

1 scenarios (1 undefined)
3 steps (3 undefined)
997.3µs

You can implement step definitions for undefined steps with these snippets:
*/

package main

import (
	"fmt"
	//	"testing"

	"github.com/cucumber/godog"
)

//var opt = godog.Options{
//	Output: colors.Colored(os.Stdout),
//	Format: "progress", // can define default values
//}
//
//func init() {
//	godog.BindFlags("godog.", flag.CommandLine, &opt)
//}
//
//func TestMain(m *testing.M) {
//	flag.Parse()
//	opt.Paths = flag.Args()
//
//	status := godog.RunWithOptions("solver", func(s *godog.Suite) {
//		FeatureContext(s)
//	}, opt)
//	if st := m.Run(); st > status {
//		status = st
//
//	}
//	os.Exit(status)
//}

func AFileContainingAListOfMasses() error {
	var a = 1
	var b = 2
	var c = a + b
	fmt.Printf("c is %d\n", c)
	return nil //godog.ErrPending
}

func IComputeTheFuelCostForAllModulesInTheFile() error {
	var ErrPending = fmt.Errorf("This test has not been implemeneted yet")
	return ErrPending
}

func TheFuelCostShouldBe() error {
	return nil //godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^a file containing a list of masses$`, AFileContainingAListOfMasses)
	s.Step(`^I compute the fuel cost for all modules in the file$`, IComputeTheFuelCostForAllModulesInTheFile)
	s.Step(`^the fuel cost should be$`, TheFuelCostShouldBe)
}

//func TestSolver(t *testing.T) {
//t.Run("Acceptance Test 1", FeatureContext)
//}
