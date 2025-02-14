Feature: Must3 function tests

  Scenario: Must3 function success
    Given values 1, 2, 3, and no error
    When the Must3 function is called
    Then the results should be 1, 2, and 3
    And it should not panic

  Scenario: Must3 function failure
    Given values 1, 2, 3, and an error "some error"
    When the Must3 function is called
    Then it should panic
