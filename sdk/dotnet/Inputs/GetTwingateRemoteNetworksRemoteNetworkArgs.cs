// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TwingateLabs.Twingate.Inputs
{

    public sealed class GetTwingateRemoteNetworksRemoteNetworkInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetTwingateRemoteNetworksRemoteNetworkInputArgs()
        {
        }
        public static new GetTwingateRemoteNetworksRemoteNetworkInputArgs Empty => new GetTwingateRemoteNetworksRemoteNetworkInputArgs();
    }
}