package tests

import (
	"time"

	"github.com/netm4ul/netm4ul/core/database/models"
)

var (
	NormalProject  models.Project
	EmptyProject   models.Project
	NormalRaw      models.Raw
	NormalProjects []models.Project
	NormalIPs      []models.IP
	NormalPorts    []models.Port
	NormalURIs     []models.URI
	NormalDomains  []models.Domain
	NormalRaws     map[string][]models.Raw
	NormalUser     models.User
)

func init() {
	NormalProject = models.Project{
		Name:        "Test project",
		Description: "Test description",
	}
	NormalIPs = []models.IP{
		{
			Value: "1.1.1.1",
		},
		{
			Value: "2.2.2.2",
		},
	}
	NormalPorts = []models.Port{
		{
			Banner:   "Test banner",
			Number:   80,
			Protocol: "tcp",
			Status:   "open",
			Type:     "web",
		},
		{
			Banner:   "Test banner 2",
			Number:   22,
			Protocol: "tcp",
			Status:   "open",
			Type:     "ssh",
		},
	}
	NormalURIs = []models.URI{
		{
			Code: "200",
			Name: "Test URI",
		},
		{
			Code: "404",
			Name: "Test not found URI",
		},
		{
			Code: "500",
			Name: "Test server error URI",
		},
	}
	NormalDomains = []models.Domain{
		{
			Name: "domain.tld",
		},
		{
			Name: "another.tld",
		},
		{
			Name: "sub1.another.tld",
		},
	}

	NormalRaw = models.Raw{
		ModuleName: "TestModule",
		CreatedAt:  time.Unix(0, 0),
		UpdatedAt:  time.Unix(1, 0),
		Content:    "Test content woa {][@&~#{[[|[`]@^ùm!::!,:;,n=))1234567\nabcder 1231<x§:;>é&",
	}

	NormalProjects = []models.Project{
		NormalProject,
		EmptyProject,
	}

	NormalRaws = map[string][]models.Raw{
		"Test": {
			NormalRaw,
			NormalRaw,
		},
	}

	NormalUser = models.User{
		Name:      "TestUser",
		Password:  "$2y$10$Fu4hg./ZybmFjiPxIpEOROGwQhF3sfwakddzlWFtV.I3rJu6sfy/2", // Test password
		Token:     "testtoken",
		CreatedAt: time.Unix(0, 0),
		UpdatedAt: time.Unix(1, 0),
	}
}
