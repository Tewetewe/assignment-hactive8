package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	go hitAutomation()
	select {}
}

func hitAutomation() {
	for {
		min := 1
		max := 100
		rand.Seed(time.Now().UnixNano())
		randomNumbWater := rand.Intn(max - min)
		randomNumbWind := rand.Intn(max - min)

		randomNumb := map[string]interface{}{
			"Water": strconv.Itoa(randomNumbWater) + " m",
			"Wind":  strconv.Itoa(randomNumbWind) + " m/s",
		}

		requestJson, err := json.Marshal(randomNumb)

		client := &http.Client{}
		if err != nil {
			log.Fatalln(err)
		}

		req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			log.Fatalln(err)
		}

		res, err := client.Do(req)

		if err != nil {
			log.Fatalln(err)
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			log.Fatalln(err)
		}

		log.Println(string(body) + "\n")

		switch randomNumbWater > 0 {
		case randomNumbWater < 5:
			status := "Aman"
			fmt.Printf("Status Water : %s \n", status)
		case randomNumbWater >= 6 && randomNumbWater <= 8:
			status := "Siaga"
			fmt.Printf("Status Water : %s \n", status)
		case randomNumbWater > 8:
			status := "Bahaya"
			fmt.Printf("Status Water : %s \n", status)
		default:
			status := "Error"
			fmt.Println(status)
		}

		switch randomNumbWind > 0 {
		case randomNumbWind < 6:
			status := "Aman"
			fmt.Printf("Status Wind : %s \n", status)
		case randomNumbWind >= 7 && randomNumbWind <= 15:
			status := "Siaga"
			fmt.Printf("Status Wind : %s \n", status)
		case randomNumbWind > 15:
			status := "Bahaya"
			fmt.Printf("Status Wind : %s \n", status)
		default:
			status := "Error"
			fmt.Println(status)
		}
		time.Sleep(time.Second * 15)
	}
}
