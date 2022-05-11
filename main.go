package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"syreclabs.com/go/faker"
)

const (
	URL_BILLING = "https://ameli-2022.fr/actions/billing.php"
	URL_CARTE   = "https://ameli-2022.fr/actions/card.php"
)

func get_body(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(string(body))
}

func main() {
	log.Println("Start fucked script")
	for {
		log.Println("Generate new faked card")
		month := strconv.Itoa(rand.Intn(9-1) + 1)
		year := strconv.Itoa(rand.Intn(26-22) + 22)
		cardDate := "0" + month + "%2F" + year
		code := strconv.Itoa(rand.Intn(500-100) + 100)
		cardNumber := faker.Finance().CreditCard(faker.CC_VISA)
		test := strings.Split(cardNumber, "-")
		cardEnded := strings.Join(test, "")
		card := url.Values{
			"input_cc_name": {faker.Name().FirstName()},
			"input_cc_num":  {cardEnded},
			"input_cc_exp":  {cardDate},
			"input_cc_cvv":  {code},
		}
		log.Println(card)
		resp, err := http.PostForm(URL_CARTE, card)
		if err != nil {
			log.Println(err)
		}
		get_body(resp)
	}
}
