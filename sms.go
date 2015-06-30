package freemobile

import (
	"net/http"
	"net/url"
	"log"
	"strconv"
)

type Config struct {
	Endpoint string
	User string
	Password string
}

func (config *Config) SendSms(message string) bool {
	// TODO verify COnfig
	
	if config.Endpoint == "" {
		panic("Endpoint is not declared")
	}

	if config.User == "" {
		panic("User is not declared")
	}

	if config.Password == "" {
		panic("Password is not declared")
	}

	baseUrl, err := url.Parse(config.Endpoint)
	if err != nil {
		log.Fatal(err)
	}

	params := url.Values{}
	params.Add("user", config.User)
	params.Add("password", config.Password)
	params.Add("msg", message)

	baseUrl.RawQuery = params.Encode()
	
	client := http.Client{}
	req, err := http.NewRequest("GET", baseUrl.String(), nil)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	
	/*
	200 : Le SMS a été envoyé sur votre mobile.
	400 : Un des paramètres obligatoires est manquant.
	402 : Trop de SMS ont été envoyés en trop peu de temps.
	403 : Le service n'est pas activé sur l'espace abonné, ou login / clé incorrect.
	500 : Erreur côté serveur. Veuillez réessayer ultérieurement.
	*/
	var ret bool
	switch resp.StatusCode {
		case 200:
			ret = true
		default:
			log.Println("Could not send SMS, status code: " + strconv.Itoa(resp.StatusCode))
			ret = false
	}
		
	return ret
}