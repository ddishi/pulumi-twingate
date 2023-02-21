// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package twingate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Security Policies are defined in the Twingate Admin Console and determine user and device authentication requirements for Resources.
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
//			_, err = twingate.GetTwingateSecurityPolicy(ctx, &GetTwingateSecurityPolicyArgs{
//				Name: pulumi.StringRef("<your security policy name>"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetTwingateSecurityPolicy(ctx *pulumi.Context, args *GetTwingateSecurityPolicyArgs, opts ...pulumi.InvokeOption) (*GetTwingateSecurityPolicyResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetTwingateSecurityPolicyResult
	err := ctx.Invoke("twingate:index/getTwingateSecurityPolicy:getTwingateSecurityPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTwingateSecurityPolicy.
type GetTwingateSecurityPolicyArgs struct {
	// Return a Security Policy by its ID. The ID for the Security Policy must be obtained from the Admin API.
	Id *string `pulumi:"id"`
	// Return a Security Policy that exactly matches this name.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getTwingateSecurityPolicy.
type GetTwingateSecurityPolicyResult struct {
	// Return a Security Policy by its ID. The ID for the Security Policy must be obtained from the Admin API.
	Id *string `pulumi:"id"`
	// Return a Security Policy that exactly matches this name.
	Name *string `pulumi:"name"`
}

func GetTwingateSecurityPolicyOutput(ctx *pulumi.Context, args GetTwingateSecurityPolicyOutputArgs, opts ...pulumi.InvokeOption) GetTwingateSecurityPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTwingateSecurityPolicyResult, error) {
			args := v.(GetTwingateSecurityPolicyArgs)
			r, err := GetTwingateSecurityPolicy(ctx, &args, opts...)
			var s GetTwingateSecurityPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTwingateSecurityPolicyResultOutput)
}

// A collection of arguments for invoking getTwingateSecurityPolicy.
type GetTwingateSecurityPolicyOutputArgs struct {
	// Return a Security Policy by its ID. The ID for the Security Policy must be obtained from the Admin API.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Return a Security Policy that exactly matches this name.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetTwingateSecurityPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTwingateSecurityPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getTwingateSecurityPolicy.
type GetTwingateSecurityPolicyResultOutput struct{ *pulumi.OutputState }

func (GetTwingateSecurityPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTwingateSecurityPolicyResult)(nil)).Elem()
}

func (o GetTwingateSecurityPolicyResultOutput) ToGetTwingateSecurityPolicyResultOutput() GetTwingateSecurityPolicyResultOutput {
	return o
}

func (o GetTwingateSecurityPolicyResultOutput) ToGetTwingateSecurityPolicyResultOutputWithContext(ctx context.Context) GetTwingateSecurityPolicyResultOutput {
	return o
}

// Return a Security Policy by its ID. The ID for the Security Policy must be obtained from the Admin API.
func (o GetTwingateSecurityPolicyResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateSecurityPolicyResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Return a Security Policy that exactly matches this name.
func (o GetTwingateSecurityPolicyResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateSecurityPolicyResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTwingateSecurityPolicyResultOutput{})
}
