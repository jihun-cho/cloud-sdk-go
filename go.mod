module github.com/elastic/cloud-sdk-go

go 1.13

require (
	github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef
	github.com/blang/semver/v4 v4.0.0
	github.com/davecgh/go-spew v1.1.1
	github.com/fatih/color v1.9.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/errors v0.19.7
	github.com/go-openapi/runtime v0.19.23
	github.com/go-openapi/spec v0.19.10
	github.com/go-openapi/strfmt v0.19.6
	github.com/go-openapi/swag v0.19.10
	github.com/go-openapi/validate v0.19.11
	github.com/hashicorp/go-multierror v1.1.0
	github.com/mitchellh/mapstructure v1.3.3
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1
)

replace sourcegraph.com/sourcegraph/go-diff v0.5.1 => github.com/sourcegraph/go-diff v0.5.1
