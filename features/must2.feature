Feature: Must2 function tests

  Scenario: Must2 function success
    Given values 1 and 2 and no error
    When the Must2 function is called
    Then the results should be 1 and 2
    And it should not panic

  Scenario: Must2 function failure
    Given values 1 and 2 and an error "some error"
    When the Must2 function is called
    Then it should panic
