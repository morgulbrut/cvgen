package cv

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type CV struct {
	Settings struct {
		Date     string `yaml:"date"`
		Cvstyle  string `yaml:"cvstyle"`
		Cvcolor  string `yaml:"cvcolor"`
		Mainfont string `yaml:"mainfont"`
		Sansfont string `yaml:"sansfont"`
		Letter   string `yaml:"letter"`
	} `yaml:"settings"`
	Personalinfo struct {
		Name     string `yaml:"name"`
		Prename  string `yaml:"prename"`
		Address  string `yaml:"address"`
		Postcode string `yaml:"postcode"`
		City     string `yaml:"city"`
		Country  string `yaml:"country"`
		Mobile   string `yaml:"mobile"`
		Phone    string `yaml:"phone"`
		Fax      string `yaml:"fax"`
		Title    string `yaml:"title"`
		Email    string `yaml:"email"`
		Homepage string `yaml:"homepage"`
		Linkedin string `yaml:"linkedin"`
		Twitter  string `yaml:"twitter"`
		Github   string `yaml:"github"`
		Gitlab   string `yaml:"gitlab"`
		Quote    string `yaml:"quote"`
		Photo    string `yaml:"photo"`
	} `yaml:"personalinfo"`
	Experience []struct {
		Company     string   `yaml:"company"`
		Role        string   `yaml:"role"`
		City        string   `yaml:"city"`
		Date        string   `yaml:"date"`
		Description string   `yaml:"description"`
		Tasks       []string `yaml:"tasks"`
	} `yaml:"experience"`
	Projects []struct {
		Title       string   `yaml:"title"`
		Context     string   `yaml:"context"`
		City        string   `yaml:"city"`
		Date        string   `yaml:"date"`
		Description string   `yaml:"description"`
		Tasks       []string `yaml:"tasks"`
	} `yaml:"projects"`
	Education []struct {
		School      string `yaml:"school"`
		Subject     string `yaml:"subject"`
		City        string `yaml:"city"`
		Date        string `yaml:"date"`
		Description string `yaml:"description"`
	} `yaml:"education"`
	Computerskills []struct {
		Group string `yaml:"group"`
		Skill string `yaml:"skill"`
	} `yaml:"computerskills"`
	Testimonials []struct {
		Name  string `yaml:"name"`
		Quote string `yaml:"quote"`
	} `yaml:"testimonials"`
	Freetext struct {
		Begin string `yaml:"begin"`
		Enf   string `yaml:"end"`
	} `yaml:"freetext"`
	Hackathons []struct {
		Name        string `yaml:"name"`
		Role        string `yaml:"role"`
		City        string `yaml:"city"`
		Date        string `yaml:"date"`
		Description string `yaml:"description"`
	} `yaml:"hackathons"`
	Talks []struct {
		Title       string `yaml:"title"`
		Coautors    string `yaml:"coautors"`
		Conference  string `yaml:"conference"`
		Date        string `yaml:"date"`
		Description string `yaml:"description"`
	} `yaml:"talks"`
	Publications []struct {
		Title       string `yaml:"title"`
		Coautors    string `yaml:"coautors"`
		Context     string `yaml:"context"`
		Date        string `yaml:"date"`
		Description string `yaml:"description"`
	} `yaml:"publications"`
	Certificates []struct {
		Type string `yaml:"type"`
		Name string `yaml:"name"`
	} `yaml:"certificates"`
	Interests []struct {
		Group       string `yaml:"group"`
		Description string `yaml:"description"`
	} `yaml:"interests"`
	Languages []struct {
		Language string `yaml:"language"`
		Level    string `yaml:"level"`
	} `yaml:"languages"`
}

func (cv *CV) Read(f string) *CV {

	yamlFile, err := ioutil.ReadFile(f)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, cv)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return cv
}
