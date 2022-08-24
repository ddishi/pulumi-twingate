// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package twingate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the twingate package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// The access key for API operations. You can retrieve this from the Twingate Admin Console
	// ([documentation](https://docs.twingate.com/docs/api-overview)). Alternatively, this can be specified using the
	// TWINGATE_API_TOKEN environment variable.
	ApiToken pulumi.StringPtrOutput `pulumi:"apiToken"`
	// Your Twingate network ID for API operations. You can find it in the Admin Console URL, for example:
	// `autoco.twingate.com`, where `autoco` is your network ID Alternatively, this can be specified using the TWINGATE_NETWORK
	// environment variable.
	Network pulumi.StringPtrOutput `pulumi:"network"`
	// The default is 'twingate.com' This is optional and shouldn't be changed under normal circumstances.
	Url pulumi.StringPtrOutput `pulumi:"url"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if isZero(args.HttpMaxRetry) {
		args.HttpMaxRetry = pulumi.IntPtr(5)
	}
	if isZero(args.HttpTimeout) {
		args.HttpTimeout = pulumi.IntPtr(10)
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:twingate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// The access key for API operations. You can retrieve this from the Twingate Admin Console
	// ([documentation](https://docs.twingate.com/docs/api-overview)). Alternatively, this can be specified using the
	// TWINGATE_API_TOKEN environment variable.
	ApiToken *string `pulumi:"apiToken"`
	// Specifies a retry limit for the http requests made. This default value is 10. Alternatively, this can be specified using
	// the TWINGATE_HTTP_MAX_RETRY environment variable
	HttpMaxRetry *int `pulumi:"httpMaxRetry"`
	// Specifies a time limit in seconds for the http requests made. The default value is 10 seconds. Alternatively, this can
	// be specified using the TWINGATE_HTTP_TIMEOUT environment variable
	HttpTimeout *int `pulumi:"httpTimeout"`
	// Your Twingate network ID for API operations. You can find it in the Admin Console URL, for example:
	// `autoco.twingate.com`, where `autoco` is your network ID Alternatively, this can be specified using the TWINGATE_NETWORK
	// environment variable.
	Network *string `pulumi:"network"`
	// The default is 'twingate.com' This is optional and shouldn't be changed under normal circumstances.
	Url *string `pulumi:"url"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// The access key for API operations. You can retrieve this from the Twingate Admin Console
	// ([documentation](https://docs.twingate.com/docs/api-overview)). Alternatively, this can be specified using the
	// TWINGATE_API_TOKEN environment variable.
	ApiToken pulumi.StringPtrInput
	// Specifies a retry limit for the http requests made. This default value is 10. Alternatively, this can be specified using
	// the TWINGATE_HTTP_MAX_RETRY environment variable
	HttpMaxRetry pulumi.IntPtrInput
	// Specifies a time limit in seconds for the http requests made. The default value is 10 seconds. Alternatively, this can
	// be specified using the TWINGATE_HTTP_TIMEOUT environment variable
	HttpTimeout pulumi.IntPtrInput
	// Your Twingate network ID for API operations. You can find it in the Admin Console URL, for example:
	// `autoco.twingate.com`, where `autoco` is your network ID Alternatively, this can be specified using the TWINGATE_NETWORK
	// environment variable.
	Network pulumi.StringPtrInput
	// The default is 'twingate.com' This is optional and shouldn't be changed under normal circumstances.
	Url pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

// The access key for API operations. You can retrieve this from the Twingate Admin Console
// ([documentation](https://docs.twingate.com/docs/api-overview)). Alternatively, this can be specified using the
// TWINGATE_API_TOKEN environment variable.
func (o ProviderOutput) ApiToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ApiToken }).(pulumi.StringPtrOutput)
}

// Your Twingate network ID for API operations. You can find it in the Admin Console URL, for example:
// `autoco.twingate.com`, where `autoco` is your network ID Alternatively, this can be specified using the TWINGATE_NETWORK
// environment variable.
func (o ProviderOutput) Network() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Network }).(pulumi.StringPtrOutput)
}

// The default is 'twingate.com' This is optional and shouldn't be changed under normal circumstances.
func (o ProviderOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
