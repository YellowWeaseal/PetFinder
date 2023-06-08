package PetFinder

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	*http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}
	return &Client{
		Client: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

func Access() {
	clientID := "U2eOP2in63EzGFoaOCSzOsjWuOJ4zJby5iPYGH6MatvHpXXPna"
	clientSecret := "LAQC3a1tN64ySgJVnBOuIL5ljVzGX3GY1fX7Sily"

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	resp, err := http.Post(
		"https://api.petfinder.com/v2/oauth2/token",
		"application/x-www-form-urlencoded",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}
	fmt.Println(string(body))
}
func (c Client) GetAnimal(id int) (Animal, error) {
	infoFromId := fmt.Sprintf("https://api.petfinder.com/v2/animals/%d", id)
	resp, err := http.Get(infoFromId)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var message interface{}
	var r Animal
	messageMap := message.(map[string]interface{})
	animalMap := messageMap["animal"].(map[string]interface{})

	err = mapstructure.Decode(animalMap, &r)
	if err != nil {
		return Animal{}, err
	}
	return r, nil
}
