Feature: Must function tests

  Scenario: Must function success
    Given a value 1 and no error
    When the Must function is called
    Then the result should be 1
    And it should not panic

  Scenario: Must function failure
    Given a value 0 and an error "some error"
    When the Must function is called
    Then it should panic
