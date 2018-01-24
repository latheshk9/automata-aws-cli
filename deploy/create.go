package deploy

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/automata-aws-cli/types"
	yaml "gopkg.in/yaml.v2"
)

func LaunchConfig(config string, env_name string, svc string) {
	automataAwsIp := "localhost"
	yamlFile, err := ioutil.ReadFile(config)
	if err != nil {
		panic(err)
	}
	var launchConfig types.LaunchConfig
	err = yaml.Unmarshal(yamlFile, &launchConfig)
	if err != nil {
		panic(err)
	}
	d, _ := yaml.Marshal(&launchConfig)
	b := bytes.NewBuffer(d)
	fmt.Println(b)

	resp, err := http.Post(automataAwsIp+"v1/createlaunchconfig/", "text/plain", b)
	if err != nil {
		panic(err)
	}
	bs, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		panic(err1)
	}
	tr := string(bs)
	fmt.Println(tr)
}

func AutoScaling(config string, env_name string, svc string) {
	automataAwsIp := "localhost"
	yamlFile, err := ioutil.ReadFile(config)
	if err != nil {
		panic(err)
	}
	var autoScaling types.AutoScaling
	err = yaml.Unmarshal(yamlFile, &autoScaling)
	if err != nil {
		panic(err)
	}
	d, _ := yaml.Marshal(&autoScaling)
	b := bytes.NewBuffer(d)

	resp, err := http.Post(automataAwsIp+"v1/autoscale/", "text/plain", b)
	if err != nil {
		panic(err)
	}
	bs, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		panic(err1)
	}
	tr := string(bs)
	fmt.Println(tr)
}
