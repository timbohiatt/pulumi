// Package serviceAccount allows to the Creation of Google Cloud Service Account Resource via Pulumi
// In a Google Opinionated best practices manner.
package serviceAccount

// Importing the specific Pulumi Resources required by this module
import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/serviceaccount"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Settings Struct - Contains the optional runtime varibles that are used during execution.
// These variables can be used to instruct the module to execute tests or custom Google Opinionated
// operation on objects created as part of the use of the module itself.
type Settings struct {
	RunTests bool
}

// Args Struct - Contains the Arguments the Pulumi Google Cloud Resource (Service Account).
// Additionally the Args Struct contains a pointer reference to the instanciated Pulumi
// Google Cloud Resource (Service Account)
type Args struct {
	ProjectId   string
	AccountId   string
	DisplayName string
	Description string
	Disabled    bool
}

// Module Struct - Contains the Arguments & Settings used to instanciate and operate the
// Pulumi Google Cloud Resource (Service Account).
// Additionally the Module Struct contains a pointer reference to the instanciated Pulumi
// Google Cloud Resource (Service Account)
type Module struct {
	Args           *Args
	Settings       *Settings
	ServiceAccount *serviceaccount.Account
}

// The Create function is a reference function of type serviceAccount Module.
// When executed on serviceAccount object it consumes the module arguments
// and instanciates a Google Cloud Service account using the Pulumi Classic provider.
// As with all reference functions, it references a pointer to itself and returns only
// an error to be caught by the calling function.
func (module *Module) Create(ctx *pulumi.Context) (err error) {

	// Create Pulumi Google Cloud Service Account Arguments Object
	args := &serviceaccount.AccountArgs{}

	// Confgure Google Cloud Service Account - Pulumi Arguments
	args.Project = pulumi.String(module.Args.ProjectId)
	args.AccountId = pulumi.String(module.Args.AccountId)
	args.DisplayName = pulumi.String(module.Args.DisplayName)
	args.Description = pulumi.String(module.Args.Description)
	args.Disabled = pulumi.Bool(module.Args.Disabled)

	serviceAccount, err := serviceaccount.NewAccount(ctx, "service-account", args)

	ctx.Export("service-account", serviceAccount)
	module.ServiceAccount = serviceAccount

	// Return the MOdule & Error
	return err

}
