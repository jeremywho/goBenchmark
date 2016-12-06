package goBenchmark_test

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"testing"
)

type simpleData struct {
	Name string `json:"name"`
}

type complexData struct {
	Id         string   `json:"_id"`
	Index      int      `json:"index"`
	Guid       string   `json:"guid"`
	IsActive   bool     `json:"isActive"`
	Balance    string   `json:"balance"`
	Picture    string   `json:"picture"`
	Age        int      `json:"age"`
	EyeColor   string   `json:"eyeColor"`
	Name       string   `json:"name"`
	Gender     string   `json:"gender"`
	Company    string   `json:"company"`
	Email      string   `json:"email"`
	Phone      string   `json:"phone"`
	Address    string   `json:"address"`
	About      string   `json:"about"`
	Registered string   `json:"registered"`
	Lat        int      `json:"latitude"`
	Long       int      `json:"longitude"`
	Tags       []string `json:"tags"`
	Greeting   string   `json:"greeting"`
	FavFruit   string   `json:"favoriteFruit"`
}

var simpleJSONString string
var inputData []byte

var s = []simpleData{}
var c = []complexData{}

func TestMain(m *testing.M) {
	flag.Parse()

	inputData, _ = ioutil.ReadFile("input.json")

	os.Exit(m.Run())
}

// BenchmarkComplexJSONParse simple benchmark
func BenchmarkComplexJSONParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(inputData, &c)
	}
}

// BenchmarkSimpleJSONParse simple benchmark
func BenchmarkSimpleJSONParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(inputData, &s)
	}
}
