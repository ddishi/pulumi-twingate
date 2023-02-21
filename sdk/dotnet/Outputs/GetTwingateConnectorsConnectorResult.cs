// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TwingateLabs.Twingate.Outputs
{

    [OutputType]
    public sealed class GetTwingateConnectorsConnectorResult
    {
        /// <summary>
        /// The ID of the Connector
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Name of the Connector
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The ID of the Remote Network attached to the Connector
        /// </summary>
        public readonly string RemoteNetworkId;

        [OutputConstructor]
        private GetTwingateConnectorsConnectorResult(
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
