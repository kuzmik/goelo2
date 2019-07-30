package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	// "github.com/davecgh/go-spew/spew"
	"github.com/mb-14/gomarkov"
)

func init() {
	// chain := gomarkov.NewChain(1)

	chain, _ := loadModel()

	tokens := []string{gomarkov.StartToken}
	for tokens[len(tokens)-1] != gomarkov.EndToken {
		next, _ := chain.Generate(tokens[(len(tokens) - 1):])
		tokens = append(tokens, next)
	}

	fmt.Println(strings.Join(tokens[1:len(tokens)-1], " "))
}

// Loads the model into the chains
func loadModel() (*gomarkov.Chain, error) {
	var chain gomarkov.Chain
	data, err := ioutil.ReadFile("model.json")
	if err != nil {
		return &chain, err
	}
	err = json.Unmarshal(data, &chain)
	if err != nil {
		return &chain, err
	}
	return &chain, nil
}

// Dumps the chains into a json file
func saveModel(chain *gomarkov.Chain) {
	jsonObj, _ := json.Marshal(chain)
	err := ioutil.WriteFile("model.json", jsonObj, 0600)
	if err != nil {
		fmt.Println(err)
	}
}
