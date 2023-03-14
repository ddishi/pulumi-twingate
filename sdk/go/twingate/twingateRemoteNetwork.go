// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package twingate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TwingateRemoteNetwork struct {
	pulumi.CustomResourceState

	// The location of the Remote Network. Must be one of the following: AWS, AZURE, GOOGLE_CLOUD, ON_PREMISE, OTHER.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the Remote Network
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewTwingateRemoteNetwork registers a new resource with the given unique name, arguments, and options.
func NewTwingateRemoteNetwork(ctx *pulumi.Context,
	name string, args *TwingateRemoteNetworkArgs, opts ...pulumi.ResourceOption) (*TwingateRemoteNetwork, error) {
	if args == nil {
		args = &TwingateRemoteNetworkArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource TwingateRemoteNetwork
	err := ctx.RegisterResource("twingate:index/twingateRemoteNetwork:TwingateRemoteNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTwingateRemoteNetwork gets an existing TwingateRemoteNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTwingateRemoteNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TwingateRemoteNetworkState, opts ...pulumi.ResourceOption) (*TwingateRemoteNetwork, error) {
	var resource TwingateRemoteNetwork
	err := ctx.ReadResource("twingate:index/twingateRemoteNetwork:TwingateRemoteNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TwingateRemoteNetwork resources.
type twingateRemoteNetworkState struct {
	// The location of the Remote Network. Must be one of the following: AWS, AZURE, GOOGLE_CLOUD, ON_PREMISE, OTHER.
	Location *string `pulumi:"location"`
	// The name of the Remote Network
	Name *string `pulumi:"name"`
}

type TwingateRemoteNetworkState struct {
	// The location of the Remote Network. Must be one of the following: AWS, AZURE, GOOGLE_CLOUD, ON_PREMISE, OTHER.
	Location pulumi.StringPtrInput
	// The name of the Remote Network
	Name pulumi.StringPtrInput
}

func (TwingateRemoteNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*twingateRemoteNetworkState)(nil)).Elem()
}

type twingateRemoteNetworkArgs struct {
	// The location of the Remote Network. Must be one of the following: AWS, AZURE, GOOGLE_CLOUD, ON_PREMISE, OTHER.
	Location *string `pulumi:"location"`
	// The name of the Remote Network
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a TwingateRemoteNetwork resource.
type TwingateRemoteNetworkArgs struct {
	// The location of the Remote Network. Must be one of the following: AWS, AZURE, GOOGLE_CLOUD, ON_PREMISE, OTHER.
	Location pulumi.StringPtrInput
	// The name of the Remote Network
	Name pulumi.StringPtrInput
}

func (TwingateRemoteNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*twingateRemoteNetworkArgs)(nil)).Elem()
}

type TwingateRemoteNetworkInput interface {
	pulumi.Input

	ToTwingateRemoteNetworkOutput() TwingateRemoteNetworkOutput
	ToTwingateRemoteNetworkOutputWithContext(ctx context.Context) TwingateRemoteNetworkOutput
}

func (*TwingateRemoteNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**TwingateRemoteNetwork)(nil)).Elem()
}

func (i *TwingateRemoteNetwork) ToTwingateRemoteNetworkOutput() TwingateRemoteNetworkOutput {
	return i.ToTwingateRemoteNetworkOutputWithContext(context.Background())
}

func (i *TwingateRemoteNetwork) ToTwingateRemoteNetworkOutputWithContext(ctx context.Context) TwingateRemoteNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwingateRemoteNetworkOutput)
}

// TwingateRemoteNetworkArrayInput is an input type that accepts TwingateRemoteNetworkArray and TwingateRemoteNetworkArrayOutput values.
// You can construct a concrete instance of `TwingateRemoteNetworkArrayInput` via:
//
//	TwingateRemoteNetworkArray{ TwingateRemoteNetworkArgs{...} }
type TwingateRemoteNetworkArrayInput interface {
	pulumi.Input

	ToTwingateRemoteNetworkArrayOutput() TwingateRemoteNetworkArrayOutput
	ToTwingateRemoteNetworkArrayOutputWithContext(context.Context) TwingateRemoteNetworkArrayOutput
}

type TwingateRemoteNetworkArray []TwingateRemoteNetworkInput

func (TwingateRemoteNetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TwingateRemoteNetwork)(nil)).Elem()
}

func (i TwingateRemoteNetworkArray) ToTwingateRemoteNetworkArrayOutput() TwingateRemoteNetworkArrayOutput {
	return i.ToTwingateRemoteNetworkArrayOutputWithContext(context.Background())
}

func (i TwingateRemoteNetworkArray) ToTwingateRemoteNetworkArrayOutputWithContext(ctx context.Context) TwingateRemoteNetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwingateRemoteNetworkArrayOutput)
}

// TwingateRemoteNetworkMapInput is an input type that accepts TwingateRemoteNetworkMap and TwingateRemoteNetworkMapOutput values.
// You can construct a concrete instance of `TwingateRemoteNetworkMapInput` via:
//
//	TwingateRemoteNetworkMap{ "key": TwingateRemoteNetworkArgs{...} }
type TwingateRemoteNetworkMapInput interface {
	pulumi.Input

	ToTwingateRemoteNetworkMapOutput() TwingateRemoteNetworkMapOutput
	ToTwingateRemoteNetworkMapOutputWithContext(context.Context) TwingateRemoteNetworkMapOutput
}

type TwingateRemoteNetworkMap map[string]TwingateRemoteNetworkInput

func (TwingateRemoteNetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TwingateRemoteNetwork)(nil)).Elem()
}

func (i TwingateRemoteNetworkMap) ToTwingateRemoteNetworkMapOutput() TwingateRemoteNetworkMapOutput {
	return i.ToTwingateRemoteNetworkMapOutputWithContext(context.Background())
}

func (i TwingateRemoteNetworkMap) ToTwingateRemoteNetworkMapOutputWithContext(ctx context.Context) TwingateRemoteNetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwingateRemoteNetworkMapOutput)
}

type TwingateRemoteNetworkOutput struct{ *pulumi.OutputState }

func (TwingateRemoteNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TwingateRemoteNetwork)(nil)).Elem()
}

func (o TwingateRemoteNetworkOutput) ToTwingateRemoteNetworkOutput() TwingateRemoteNetworkOutput {
	return o
}

func (o TwingateRemoteNetworkOutput) ToTwingateRemoteNetworkOutputWithContext(ctx context.Context) TwingateRemoteNetworkOutput {
	return o
}

// The location of the Remote Network. Must be one of the following: AWS, AZURE, GOOGLE_CLOUD, ON_PREMISE, OTHER.
func (o TwingateRemoteNetworkOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TwingateRemoteNetwork) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the Remote Network
func (o TwingateRemoteNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TwingateRemoteNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type TwingateRemoteNetworkArrayOutput struct{ *pulumi.OutputState }

func (TwingateRemoteNetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TwingateRemoteNetwork)(nil)).Elem()
}

func (o TwingateRemoteNetworkArrayOutput) ToTwingateRemoteNetworkArrayOutput() TwingateRemoteNetworkArrayOutput {
	return o
}

func (o TwingateRemoteNetworkArrayOutput) ToTwingateRemoteNetworkArrayOutputWithContext(ctx context.Context) TwingateRemoteNetworkArrayOutput {
	return o
}

func (o TwingateRemoteNetworkArrayOutput) Index(i pulumi.IntInput) TwingateRemoteNetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TwingateRemoteNetwork {
		return vs[0].([]*TwingateRemoteNetwork)[vs[1].(int)]
	}).(TwingateRemoteNetworkOutput)
}

type TwingateRemoteNetworkMapOutput struct{ *pulumi.OutputState }

func (TwingateRemoteNetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TwingateRemoteNetwork)(nil)).Elem()
}

func (o TwingateRemoteNetworkMapOutput) ToTwingateRemoteNetworkMapOutput() TwingateRemoteNetworkMapOutput {
	return o
}

func (o TwingateRemoteNetworkMapOutput) ToTwingateRemoteNetworkMapOutputWithContext(ctx context.Context) TwingateRemoteNetworkMapOutput {
	return o
}

func (o TwingateRemoteNetworkMapOutput) MapIndex(k pulumi.StringInput) TwingateRemoteNetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TwingateRemoteNetwork {
		return vs[0].(map[string]*TwingateRemoteNetwork)[vs[1].(string)]
	}).(TwingateRemoteNetworkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TwingateRemoteNetworkInput)(nil)).Elem(), &TwingateRemoteNetwork{})
	pulumi.RegisterInputType(reflect.TypeOf((*TwingateRemoteNetworkArrayInput)(nil)).Elem(), TwingateRemoteNetworkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TwingateRemoteNetworkMapInput)(nil)).Elem(), TwingateRemoteNetworkMap{})
	pulumi.RegisterOutputType(TwingateRemoteNetworkOutput{})
	pulumi.RegisterOutputType(TwingateRemoteNetworkArrayOutput{})
	pulumi.RegisterOutputType(TwingateRemoteNetworkMapOutput{})
}
