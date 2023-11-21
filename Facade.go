package main

import "fmt"

//facade

type walletFacade struct {
	wallet       *wallet
	account      *account
	ledger       *ledger
	notification *notification
	securitycode *securitycode
}

func newWalletFacade(accountid string, code int) *walletFacade {
	fmt.Println("creating your account")
	WF := &walletFacade{
		account:      newAcc(accountid),
		securitycode: newSecurity(code),
		wallet:       newWallet(),
		notification: &notification{},
		ledger:       &ledger{},
	}
	fmt.Println("Accoutn created")
	return WF
}

func (wf *walletFacade) addMoneyToWallet(account string, code int, amount int) error {
	err := wf.account.checkAcc(account)
	if err != nil {
		return err
	}
	err = wf.securitycode.checkSecurity(code)
	if err != nil {
		return err
	}
	wf.wallet.addMoney(amount)
	wf.notification.sendingWalletCreditNotification()
	wf.ledger.makeEntry(account, "credit", amount)
	return nil
}

func (wf *walletFacade) takeMoneyFrom(account string, code int, amount int) error {
	err := wf.account.checkAcc(account)
	if err != nil {
		return err
	}
	err = wf.securitycode.checkSecurity(code)
	if err != nil {
		return err
	}
	err = wf.wallet.takeMoney(amount)
	if err != nil {
		return err
	}
	wf.notification.sendingdebitcardNotification()
	wf.ledger.makeEntry(account, "debit", amount)
	return nil
}

//account
type account struct {
	name string
}

func newAcc(name string) *account {
	return &account{name: name}
}

func (a *account) checkAcc(name string) error {
	if a.name == name {
		fmt.Println("Verified")
		return nil
	}
	return fmt.Errorf("shite")
}

//security code
type securitycode struct {
	code int
}

func newSecurity(code int) *securitycode {
	return &securitycode{code: code}
}

func (s *securitycode) checkSecurity(code int) error {
	if s.code == code {
		fmt.Println("you are in")
		return nil
	}
	return fmt.Errorf("bitch ass")
}

//wallet

type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{balance: 0}
}

func (w *wallet) addMoney(amount int) {
	w.balance += amount
	fmt.Printf("%d dollars added sucessfuly", amount)
	return
}

func (w *wallet) takeMoney(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("you can'y take that shit")
	}
	w.balance -= amount
	fmt.Printf("%d of dollars are left", w.balance)
	return nil
}

//ledger

type ledger struct{}

func (l *ledger) makeEntry(accountId, txntype string, amount int) {
	fmt.Printf("make entry fot them %s, %s, %d", accountId, txntype, amount)
	return
}

//notification
type notification struct {
}

func (n *notification) sendingWalletCreditNotification() {
	fmt.Println("sending Wallet Credit Notification")
}
func (n *notification) sendingdebitcardNotification() {
	fmt.Println("sendin gdebit card Notification")
}
