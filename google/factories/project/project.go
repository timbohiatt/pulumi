package project

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/timbohiatt/pulumi/google/module/project"
)

type Settings struct {
	RunTests bool
}

type Args struct {
	BillingAccount string
	Name           string
	ParentFolder   string
	ParentOrg      string
	Prefix         string
	Services       []string
}

type Factory struct {
	Args     *Args
	Settings *Settings
	Project  *organizations.Project
}

func (factory *Factory) Create(ctx *pulumi.Context) (err error) {

	// Create Project

	// Create Pulumi Google Cloud Project Arguments Object
	project := &project.Module{}

	// Confgure Google Cloud Project - Pulumi Arguments
	project.Args.Name = factory.Args.Name
	project.Args.ProjectId = factory.Args.Name
	project.Args.BillingAccount = factory.Args.Name
	project.Args.AutoCreateNetwork = false
	if factory.Args.ParentFolder != "" {
		project.Args.FolderId = factory.Args.ParentFolder
	} else {
		project.Args.OrgId = factory.Args.ParentOrg
	}

	// Create the Project from Module
	err = project.Create(ctx)
	if err != nil {
		return err
	}
	factory.Project = project.Project

	// Create Service Account

	// Return the Factory & Error
	return err

}
