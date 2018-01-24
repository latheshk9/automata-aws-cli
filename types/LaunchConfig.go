package types

type LaunchConfig struct {
	AmiImageId       string `yaml:"amiImageId"`
	InstanceType     string `yaml:"instanceType"`
	LaunchConfigName string `yaml:"launchConfigName"`
	SecurityGroups   string `yaml:"securityGroups"`
	UserData         string `yaml:"userData"`
	KeyPairName      string `yaml:"keyPairName"`
}
