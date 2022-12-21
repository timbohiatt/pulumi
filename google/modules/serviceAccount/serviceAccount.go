package serviceAccount

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/serviceAccount"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Settings struct {
	RunTests bool
}

type Args struct {
	ProjectId   string
	AccountId   string
	DisplayName string
	Description string
	Disabled    bool
}

type Module struct {
	Args           *Args
	Settings       *Settings
	ServiceAccount *serviceAccount.Account
}

func (module *Module) Create(ctx *pulumi.Context) (err error) {

	// Create Pulumi Google Cloud Service Account Arguments Object
	args := &serviceAccount.AccountArgs{}

	// Confgure Google Cloud Service Account - Pulumi Arguments
	args.Project = pulumi.String(module.Args.ProjectId)
	args.AccountId = pulumi.String(module.Args.AccountId)
	args.DisplayName = pulumi.String(module.Args.DisplayName)
	args.Description = pulumi.String(module.Args.Description)
	args.Disabled = pulumi.Bool(module.Args.Disabled)

	serviceAccount, err := serviceAccount.NewAccount(ctx, "service-account", args)

	ctx.Export("service-account", serviceAccount)
	module.ServiceAccount = serviceAccount

	// Return the MOdule & Error
	return err

}
