// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package twingate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources in Twingate represent servers on the private network that clients can connect to. Resources can be defined by IP, CIDR range, FQDN, or DNS zone. For more information, see the Twingate [documentation](https://docs.twingate.com/docs/resources-and-access-nodes).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/Twingate-Labs/pulumi-twingate/sdk/go/twingate"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err = twingate.GetTwingateResources(ctx, &GetTwingateResourcesArgs{
//				Name: "<your resource's name>",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetTwingateResources(ctx *pulumi.Context, args *GetTwingateResourcesArgs, opts ...pulumi.InvokeOption) (*GetTwingateResourcesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetTwingateResourcesResult
	err := ctx.Invoke("twingate:index/getTwingateResources:getTwingateResources", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTwingateResources.
type GetTwingateResourcesArgs struct {
	// The name of the Resource
	Name string `pulumi:"name"`
	// List of Resources
	Resources []GetTwingateResourcesResource `pulumi:"resources"`
}

// A collection of values returned by getTwingateResources.
type GetTwingateResourcesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the Resource
	Name string `pulumi:"name"`
	// List of Resources
	Resources []GetTwingateResourcesResource `pulumi:"resources"`
}

func GetTwingateResourcesOutput(ctx *pulumi.Context, args GetTwingateResourcesOutputArgs, opts ...pulumi.InvokeOption) GetTwingateResourcesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTwingateResourcesResult, error) {
			args := v.(GetTwingateResourcesArgs)
			r, err := GetTwingateResources(ctx, &args, opts...)
			var s GetTwingateResourcesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTwingateResourcesResultOutput)
}

// A collection of arguments for invoking getTwingateResources.
type GetTwingateResourcesOutputArgs struct {
	// The name of the Resource
	Name pulumi.StringInput `pulumi:"name"`
	// List of Resources
	Resources GetTwingateResourcesResourceArrayInput `pulumi:"resources"`
}

func (GetTwingateResourcesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTwingateResourcesArgs)(nil)).Elem()
}

// A collection of values returned by getTwingateResources.
type GetTwingateResourcesResultOutput struct{ *pulumi.OutputState }

func (GetTwingateResourcesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTwingateResourcesResult)(nil)).Elem()
}

func (o GetTwingateResourcesResultOutput) ToGetTwingateResourcesResultOutput() GetTwingateResourcesResultOutput {
	return o
}

func (o GetTwingateResourcesResultOutput) ToGetTwingateResourcesResultOutputWithContext(ctx context.Context) GetTwingateResourcesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetTwingateResourcesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTwingateResourcesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the Resource
func (o GetTwingateResourcesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetTwingateResourcesResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of Resources
func (o GetTwingateResourcesResultOutput) Resources() GetTwingateResourcesResourceArrayOutput {
	return o.ApplyT(func(v GetTwingateResourcesResult) []GetTwingateResourcesResource { return v.Resources }).(GetTwingateResourcesResourceArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTwingateResourcesResultOutput{})
}
