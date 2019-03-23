package bank

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/godog"
)

var opt = godog.Options{
	Format: "progress",
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opt.Paths = flag.Args()

	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, opt)

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

var testAccount *account

func iHaveABankAccountWith(balance int) error {
	testAccount = &account{balance: balance}
	return nil
}

func iDeposit(amount int) error {
	testAccount.deposit(amount)
	return nil
}

func iWithdraw(amount int) error {
	testAccount.withdraw(amount)
	return nil
}

func itShouldHaveABalanceOf(balance int) error {
	if testAccount.balance == balance {
		return nil
	}
	return fmt.Errorf("Incorrect account balance")
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I have a bank account with (\d+)\$$`, iHaveABankAccountWith)
	s.Step(`^I deposit (\d+)\$$`, iDeposit)
	s.Step(`^I withdraw (\d+)\$$`, iWithdraw)
	s.Step(`^it should have a balance of (\d+)\$$`, itShouldHaveABalanceOf)

	s.BeforeScenario(func(interface{}) {
		testAccount = nil
	})
}
