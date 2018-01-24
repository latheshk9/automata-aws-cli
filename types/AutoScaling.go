package types

type AutoScaling struct {
	AutoScalingGroupName string `yaml:"autoScalingGroupName"`
	LaunchConfigName     string `yaml:"launchConfigName"`
	MaxSize              string `yaml:"maxSize"`
	MinSize              string `yaml:"minSize"`
	VpcZoneIdentifier    string `yaml:"vpcZoneIdentifier"`
}
