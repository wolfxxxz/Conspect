package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

/*
exampl.yml
hits: 55
time: 19581412
*/
type Conf struct {
	Hits int64 `yaml:"hits"`
	Time int64 `yaml:"time"`
}

func GetConf(file string) {
	config := &Conf{}
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		log.Printf("yamFile.Get err: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	fmt.Println(config)
}

func WriteYaml(file string) {
	config := &Conf{Hits: 333, Time: 21212024}
	out, _ := yaml.Marshal(config)

	err := os.WriteFile(file, out, 0666)
	if err != nil {
		log.Printf("err WriteFile %v\n", err)
	}
}

/*
func main(){
	config := &BuildConf{}
	config.ReadBigConf("anchor_example.yml")
	config.WriteBigConf("anchor_example2.yml")
}
*/
/*file1.yml
definitions:
 step:
 - step: &build-test
    name: Build and test
    script:
      - mvn package
    artifacts:
      - target/**

pipelines:
  branches:
    develop:
      - step: *build-test
    main:
      - step: *build-test
*/

type BuildConf struct {
	Definitions map[string]interface{} `yaml:"definitions"`
	Pipelines   map[string]interface{} `yaml:"pipelines"`
}

func (config *BuildConf) ReadBigConf(file string) {

	yamlFile, err := os.ReadFile(file)
	if err != nil {
		log.Printf("ReadFile %v", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(config)
}

func (config *BuildConf) WriteBigConf(file string) {
	out, _ := yaml.Marshal(config)
	err := os.WriteFile(file, out, 0666)
	if err != nil {
		log.Printf("err WriteFile %v\n", err)
	}
}

/*
definitions:
  step:
  - step:
      artifacts:
      - target/**
      name: Build and test
      script:
      - mvn package
pipelines:
  branches:
    develop:
    - step:
        artifacts:
        - target/**
        name: Build and test
        script:
        - mvn package
    main:
    - step:
        artifacts:
        - target/**
        name: Build and test
        script:
        - mvn package
*/
