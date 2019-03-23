Feature: bank account
  A user's bank account must be able to withdraw and deposit cash

  Scenario: Deposit
    Given I have a bank account with 10$
    When I deposit 10$
    Then it should have a balance of 20$

  Scenario: Withdrawal
    Given I have a bank account with 20$
    When I withdraw 10$
    Then it should have a balance of 10$
  