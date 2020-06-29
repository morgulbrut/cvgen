package cv

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Letter struct {
	Recipient struct {
		Name     string `yaml:"name"`
		Company  string `yaml:"company"`
		Prename  string `yaml:"prename"`
		Address  string `yaml:"address"`
		Postcode string `yaml:"postcode"`
		City     string `yaml:"city"`
		Country  string `yaml:"country"`
		Opening  string `yaml:"opening"`
		Closing  string `yaml:"closing"`
		Letter   string `yaml:"letter"`
	} `yaml:"recipient"`
	Text []string `yaml:"text"`
}

func (l *Letter) Read(f string) *Letter {

	yamlFile, err := ioutil.ReadFile(f)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, l)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return l
}
