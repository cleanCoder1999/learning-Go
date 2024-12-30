package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	// this is a (function) type conversion from "func LogOutput(msg string)" to "type LoggerAdapter func (msg string)"
	// why does this work?
	// because FUNCTIONS ARE also considered VALUES in Go,
	// and their TYPE IS built out of their SIGNATURE (keyword "func" + types of their params + types of return values)
	l := LoggerAdapter(LogOutput)

	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)

	http.HandleFunc("/hello", c.SayHello)
	http.ListenAndServe(":8080", nil)
}

// ############# Controller
type Controller struct {
	l     Logger
	logic Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")

	userId := r.URL.Query().Get("user_id")

	msg, err := c.logic.SayHello(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(msg))
}

// NOTE: accept interfaces and return structs
func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

// ############# BusinessLogic
type Logic interface {
	SayHello(userId string) (string, error)
}

// ############# Logic
type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userId string) (string, error) {
	sl.l.Log("in SayHello for " + userId)

	name, ok := sl.ds.UserNameForId(userId)
	if !ok {
		return "", errors.New("unknown user")
	}

	return "Hello " + name, nil
}

// nobody cares about SayGoodbye (so there is no need to provide an interface to the web controller)
func (sl SimpleLogic) SayGoodbye(userId string) (string, error) {
	sl.l.Log("in SayGoodbye for " + userId)

	name, ok := sl.ds.UserNameForId(userId)
	if !ok {
		return "", errors.New("unknown user")
	}

	return "Goodbye " + name, nil
}

// NOTE: accept interfaces and return structs
func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

// ############# Logger
type Logger interface {
	Log(msg string)
}

// to make func LogOutput match the Logger interface,
// we define a FUNCTION TYPE (LoggerAdaoter) to build a BRIDGE to the INTERFACE
// (1)
type LoggerAdapter func(msg string)

// (2)
func (la LoggerAdapter) Log(msg string) {
	la(msg)
}

func LogOutput(msg string) {
	fmt.Println(msg)
}

// ############# DataStore
type DataStore interface {
	UserNameForId(userId string) (string, bool)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForId(userId string) (string, bool) {
	name, ok := sds.userData[userId]
	return name, ok
}

// NOTE: accept interfaces and return structs
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Bob",
			"3": "Dan",
		},
	}
}