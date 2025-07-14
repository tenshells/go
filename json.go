package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"log"
)

type User struct {
    ID       int     `json:"id"`
    Name     string  `json:"name"`
    Username string  `json:"username"`
    Email    string  `json:"email"`
    Address  Address `json:"address"`
    Phone    string  `json:"phone"`
    Website  string  `json:"website"`
    Company  Company `json:"company"`
}

type Address struct {
    Street  string `json:"street"`
    Suite   string `json:"suite"`
    City    string `json:"city"`
    Zipcode string `json:"zipcode"`
    Geo     Geo    `json:"geo"`
}

type Geo struct {
    Lat string `json:"lat"`
    Lng string `json:"lng"`
}

type Company struct {
    Name        string `json:"name"`
    CatchPhrase string `json:"catchPhrase"`
    Bs          string `json:"bs"`
}

func (u User) String() string {
    return fmt.Sprintf(
        "User:\n"+
            "\tID: %d\n"+
            "\tName: %s\n"+
            "\tUsername: %s\n"+
            "\tEmail: %s\n"+
            "\tPhone: %s\n"+
            "\tWebsite: %s\n"+
            "\tAddress:\n"+
            "\t\tStreet: %s\n"+
            "\t\tSuite: %s\n"+
            "\t\tCity: %s\n"+
            "\t\tZipcode: %s\n"+
            "\t\tGeo:\n"+
            "\t\t\tLat: %s\n"+
            "\t\t\tLng: %s\n"+
            "\tCompany:\n"+
            "\t\tName: %s\n"+
            "\t\tCatchPhrase: %s\n"+
            "\t\tBs: %s\n",
        u.ID, u.Name, u.Username, u.Email, u.Phone, u.Website,
        u.Address.Street, u.Address.Suite, u.Address.City, u.Address.Zipcode,
        u.Address.Geo.Lat, u.Address.Geo.Lng,
        u.Company.Name, u.Company.CatchPhrase, u.Company.Bs,
    )
}
	
func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
    	if err != nil {
        	log.Fatal(err)
    	}
    	defer resp.Body.Close()

    	body, err := ioutil.ReadAll(resp.Body)
    	if err != nil {
        	log.Fatal(err)
    	}

    	var user User
    	if err := json.Unmarshal(body, &user); err != nil {
        	log.Fatal(err)
    	}

    	fmt.Printf("%+v\n", user)
}
