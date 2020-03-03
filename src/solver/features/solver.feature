# file: $GOPATH/src/solver/features/solver.feature

Feature: Compute required fuel
    In order to travel around the world on a rocket ship
    As a pilot I need to be able to compute the rocket fuel required for the trip

    Scenario: Compute the fuel cost
        Given a file containing a list of masses
        When I compute the fuel cost for all modules in the file
        Then the fuel cost should be

