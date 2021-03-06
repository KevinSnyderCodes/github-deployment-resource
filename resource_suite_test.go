package resource_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shipt/go-github/v32/github"

	"testing"
)

func TestGithubDeploymentResource(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GithubDeploymentResource Suite")
}

func newDeployment(id int64) *github.Deployment {
	return &github.Deployment{
		ID: github.Int64(id),
	}
}

func newDeploymentWithEnvironment(id int64, env string) *github.Deployment {
	return &github.Deployment{
		ID:          github.Int64(id),
		Environment: &env,
	}
}
