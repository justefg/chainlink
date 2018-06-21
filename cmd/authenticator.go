package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/smartcontractkit/chainlink/logger"
	"github.com/smartcontractkit/chainlink/store"
	"golang.org/x/crypto/ssh/terminal"
)

// Authenticator implements the Authenticate method for the store and
// a password string.
type Authenticator interface {
	Authenticate(*store.Store, string) error
}

// TerminalAuthenticator contains fields for prompting the user and an
// exit code.
type TerminalAuthenticator struct {
	Prompter Prompter
}

// Authenticate checks to see if there are accounts present in
// the KeyStore, and if there are none, a new account will be created
// by prompting for a password. If there are accounts present, the
// account which is unlocked by the given password will be used.
func (auth TerminalAuthenticator) Authenticate(store *store.Store, pwd string) error {
	if len(pwd) != 0 {
		return auth.authenticateWithPwd(store, pwd)
	} else if auth.Prompter.IsTerminal() {
		return auth.authenticationPrompt(store)
	} else {
		return errors.New("No password provided")
	}
}

func (auth TerminalAuthenticator) authenticationPrompt(store *store.Store) error {
	if store.KeyStore.HasAccounts() {
		return auth.promptAndCheckPasswordLoop(store)
	}
	return auth.promptAndCreateAccount(store)
}

func (auth TerminalAuthenticator) authenticateWithPwd(store *store.Store, pwd string) error {
	if !store.KeyStore.HasAccounts() {
		fmt.Println("There are no accounts, creating a new account with the specified password")
		return createAccount(store, pwd)
	}
	return checkPassword(store, pwd)
}

func checkPassword(store *store.Store, phrase string) error {
	if err := store.KeyStore.Unlock(phrase); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (auth TerminalAuthenticator) promptAndCheckPasswordLoop(store *store.Store) error {
	for {
		phrase := auth.Prompter.Prompt("Enter Password:")
		if checkPassword(store, phrase) == nil {
			break
		}
	}

	return nil
}

func (auth TerminalAuthenticator) promptAndCreateAccount(store *store.Store) error {
	for {
		phrase := auth.Prompter.Prompt("New Password: ")
		clearLine()
		phraseConfirmation := auth.Prompter.Prompt("Confirm Password: ")
		clearLine()
		if phrase == phraseConfirmation {
			return createAccount(store, phrase)
		}
		fmt.Printf("Passwords don't match. Please try again... ")
	}
}

func createAccount(store *store.Store, password string) error {
	_, err := store.KeyStore.NewAccount(password)
	return err
}

// Prompter implements the Prompt function to be used to display at
// the console.
type Prompter interface {
	Prompt(string) string
	IsTerminal() bool
}

// PasswordPrompter is used to display and read input from the user.
type PasswordPrompter struct{}

// Prompt displays the prompt for the user to enter the password and
// reads their input.
func (pp PasswordPrompter) Prompt(prompt string) string {
	var rval string
	withTerminalResetter(func() {
		fmt.Print(prompt)
		bytePwd, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			logger.Fatal(err)
		}
		clearLine()
		rval = string(bytePwd)
	})
	return rval
}

// IsTerminal checks if the current process is executing in a terminal, this
// should be used to decide when to use the PasswordPrompter.
func (pp PasswordPrompter) IsTerminal() bool {
	return terminal.IsTerminal(int(os.Stdout.Fd()))
}

// Explicitly reset terminal state in the event of a signal (CTRL+C)
// to ensure typed characters are echoed in terminal:
// https://groups.google.com/forum/#!topic/Golang-nuts/kTVAbtee9UA
func withTerminalResetter(f func()) {
	osSafeStdin := int(os.Stdin.Fd())

	initialTermState, err := terminal.GetState(osSafeStdin)
	if err != nil {
		logger.Fatal(err)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		<-c
		terminal.Restore(osSafeStdin, initialTermState)
		os.Exit(1)
	}()

	f()
	signal.Stop(c)
}

func clearLine() {
	fmt.Printf("\r" + strings.Repeat(" ", 60) + "\r")
}
