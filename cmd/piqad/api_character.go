package main

import (
	"context"
	"strings"
	"sync"

	"github.com/antihax/optional"
	"github.com/azzz/piqad/pkg/stapi"
)

func searchCharacters(name string, apiClient *stapi.APIClient) ([]stapi.CharacterBase, error) {
	response, _, err := apiClient.CharacterApi.CharacterSearchPost(context.TODO(), &stapi.CharacterSearchPostOpts{
		Name: optional.NewString(name),
	})

	if err != nil {
		return nil, err
	}

	return response.Characters, nil
}

func getCharacter(uid string, apiClient *stapi.APIClient) (*stapi.CharacterFull, error) {
	response, _, err := apiClient.CharacterApi.CharacterGet(context.TODO(), uid, nil)

	if err != nil {
		return nil, err
	}

	return response.Character, nil
}

type fetchCharacter struct {
	character *stapi.CharacterFull
	err       error
}

func getCharacters(uids []string, apiClient *stapi.APIClient) []*stapi.CharacterFull {
	var wg sync.WaitGroup
	var fetchChan = make(chan fetchCharacter, len(uids))

	wg.Add(len(uids))

	for _, uid := range uids {
		// fmt.Println(uid)
		go func(uid string) {
			defer wg.Done()
			c, err := getCharacter(uid, apiClient)
			fetchChan <- fetchCharacter{c, err}
		}(uid)
	}

	wg.Wait()
	close(fetchChan)

	var fullCharacters []*stapi.CharacterFull
	for fetch := range fetchChan {
		if fetch.err == nil {
			fullCharacters = append(fullCharacters, fetch.character)
		}
	}

	return fullCharacters
}

func speciesToString(species []stapi.CharacterSpecies) string {
	var names []string

	for _, sp := range species {
		names = append(names, sp.Name)
	}

	return strings.Join(names, ", ")
}
