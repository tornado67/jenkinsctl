package api

import (
	"errors"

	"github.com/emman27/jenkinsctl/pkg/builds"
)

const parameterActionClass = "hudson.model.ParametersAction"

// GetParameters the parameters for a particular job and build
func (c *JenkinsClient) GetParameters(jobName string, buildID int) (*builds.BuildParameters, error) {
	build, err := c.GetBuild(jobName, buildID)
	if err != nil {
		return nil, err
	}
	return getParams(build)
}

func getParams(build *builds.Build) (*builds.BuildParameters, error) {
	for _, action := range build.Actions {
		if action.Class == parameterActionClass {
			params := builds.BuildParameters(*action.Parameters)
			return &params, nil
		}
	}
	return nil, errors.New("this build is not parameterized")
}
