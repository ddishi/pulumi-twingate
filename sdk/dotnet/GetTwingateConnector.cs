// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TwingateLabs.Twingate
{
    public static class GetTwingateConnector
    {
        public static Task<GetTwingateConnectorResult> InvokeAsync(GetTwingateConnectorArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTwingateConnectorResult>("twingate:index/getTwingateConnector:getTwingateConnector", args ?? new GetTwingateConnectorArgs(), options.WithDefaults());

        public static Output<GetTwingateConnectorResult> Invoke(GetTwingateConnectorInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTwingateConnectorResult>("twingate:index/getTwingateConnector:getTwingateConnector", args ?? new GetTwingateConnectorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTwingateConnectorArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetTwingateConnectorArgs()
        {
        }
        public static new GetTwingateConnectorArgs Empty => new GetTwingateConnectorArgs();
    }

    public sealed class GetTwingateConnectorInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetTwingateConnectorInvokeArgs()
        {
        }
        public static new GetTwingateConnectorInvokeArgs Empty => new GetTwingateConnectorInvokeArgs();
    }


    [OutputType]
    public sealed class GetTwingateConnectorResult
    {
        public readonly string Id;
        public readonly string Name;
        public readonly string RemoteNetworkId;

        [OutputConstructor]
        private GetTwingateConnectorResult(
            string id,

            string name,

            string remoteNetworkId)
        {
            Id = id;
            Name = name;
            RemoteNetworkId = remoteNetworkId;
        }
    }
}