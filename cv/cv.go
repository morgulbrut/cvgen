package cv

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type CV struct {
	Personalinfo struct {
		Prename  string      `yaml:"prename"`
		Name     string      `yaml:"name"`
		Address  string      `yaml:"address"`
		Postcode string      `yaml:"postcode"`
		City     string      `yaml:"city"`
		Country  string      `yaml:"country"`
		Mobile   string      `yaml:"mobile"`
		Phone    string      `yaml:"phone"`
		Fax      string      `yaml:"fax"`
		Email    string      `yaml:"email"`
		Homepage string      `yaml:"homepage"`
		Linkedin string      `yaml:"linkedin"`
		Twitter  string      `yaml:"twitter"`
		Github   string      `yaml:"github"`
		Gitlab   string      `yaml:"gitlab"`
		Quote    string      `yaml:"quote"`
		Photo    interface{} `yaml:"photo"`
		Title    string      `yaml:"title"`
	} `yaml:"personalinfo"`
	Freetext struct {
		Begin      string `yaml:"begin"`
		Begintitle string `yaml:"begintitle"`
		End        string `yaml:"end"`
		Endtitle   string `yaml:"endtitle"`
	} `yaml:"freetext"`
	Testimonials struct {
		Title        string `yaml:"title"`
		Testimonials []struct {
			Name  string `yaml:"name"`
			Quote string `yaml:"quote"`
		} `yaml:"testimonials"`
	} `yaml:"testimonials"`
	Experience struct {
		Title string `yaml:"title"`
		Jobs  []struct {
			Company     string   `yaml:"company"`
			Role        string   `yaml:"role"`
			City        string   `yaml:"city"`
			Date        string   `yaml:"date"`
			Description string   `yaml:"description"`
			Tasks       []string `yaml:"tasks"`
		} `yaml:"jobs"`
	} `yaml:"experience"`
	Education struct {
		Title   string `yaml:"title"`
		Schools []struct {
			School      string `yaml:"school"`
			Subject     string `yaml:"subject"`
			City        string `yaml:"city"`
			Date        string `yaml:"date"`
			Description string `yaml:"description"`
		} `yaml:"schools"`
	} `yaml:"education"`
	Languages struct {
		Title     string `yaml:"title"`
		Languages []struct {
			Language string `yaml:"language"`
			Level    string `yaml:"level"`
		} `yaml:"languages"`
	} `yaml:"languages"`
	Computerskills struct {
		Title          string `yaml:"title"`
		Computerskills []struct {
			Group string `yaml:"group"`
			Skill string `yaml:"skill"`
		} `yaml:"computerskills"`
	} `yaml:"computerskills"`
	Hackathons struct {
		Title      string `yaml:"title"`
		Hackathons []struct {
			Name        string `yaml:"name"`
			Role        string `yaml:"role"`
			City        string `yaml:"city"`
			Date        string `yaml:"date"`
			Description string `yaml:"description"`
		} `yaml:"hackathons"`
	} `yaml:"hackathons"`
	Projects struct {
		Title    string `yaml:"title"`
		Projects []struct {
			Title       string   `yaml:"title"`
			City        string   `yaml:"city"`
			Context     string   `yaml:"context"`
			Date        string   `yaml:"date"`
			Description string   `yaml:"description"`
			Tasks       []string `yaml:"tasks"`
		} `yaml:"projects"`
	} `yaml:"projects"`
	Talks struct {
		Title string `yaml:"title"`
		Talks []struct {
			Title       string `yaml:"title"`
			Coautors    string `yaml:"coautors"`
			Conference  string `yaml:"conference"`
			Date        string `yaml:"date"`
			Description string `yaml:"description"`
		} `yaml:"talks"`
	} `yaml:"talks"`
	Publications struct {
		Title        string `yaml:"title"`
		Publications []struct {
			Title       string `yaml:"title"`
			Coautors    string `yaml:"coautors"`
			Context     string `yaml:"context"`
			Date        string `yaml:"date"`
			Description string `yaml:"description"`
		} `yaml:"publications"`
	} `yaml:"publications"`
	Certificates struct {
		Title        string `yaml:"title"`
		Certificates []struct {
			Type string `yaml:"Type"`
			Name string `yaml:"Name"`
		} `yaml:"certificates"`
	} `yaml:"certificates"`
	Interests struct {
		Title     string `yaml:"title"`
		Interests []struct {
			Group       string `yaml:"group"`
			Description string `yaml:"description"`
		} `yaml:"interests"`
	} `yaml:"interests"`
	Skills []struct {
		Name  string `yaml:"name"`
		Level int    `yaml:"level"`
	} `yaml:"skills"`
	Hobbies   string `yaml:"hobbies"`
	Nonprofit string `yaml:"nonprofit"`
	Settings  struct {
		Date     string      `yaml:"date"`
		Cvstyle  string      `yaml:"cvstyle"`
		Cvcolor  string      `yaml:"cvcolor"`
		Mainfont interface{} `yaml:"mainfont"`
		Sansfont interface{} `yaml:"sansfont"`
	} `yaml:"settings"`
	Recipient struct {
		Name     string      `yaml:"name"`
		Company  string      `yaml:"company"`
		Prename  interface{} `yaml:"prename"`
		Address  string      `yaml:"address"`
		Postcode int         `yaml:"postcode"`
		City     string      `yaml:"city"`
	} `yaml:"recipient"`
	Letter struct {
		Date    string   `yaml:"date"`
		Opening string   `yaml:"opening"`
		Closing string   `yaml:"closing"`
		Text    []string `yaml:"text"`
	} `yaml:"letter"`
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
