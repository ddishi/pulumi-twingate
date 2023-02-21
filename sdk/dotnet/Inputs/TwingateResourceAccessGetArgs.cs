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

    public sealed class TwingateResourceAccessGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("groupIds")]
        private InputList<string>? _groupIds;

        /// <summary>
        /// List of Group IDs that will have permission to access the Resource.
        /// </summary>
        public InputList<string> GroupIds
        {
            get => _groupIds ?? (_groupIds = new InputList<string>());
            set => _groupIds = value;
        }

        [Input("serviceAccountIds")]
        private InputList<string>? _serviceAccountIds;

        /// <summary>
        /// List of Service Account IDs that will have permission to access the Resource.
        /// </summary>
        public InputList<string> ServiceAccountIds
        {
            get => _serviceAccountIds ?? (_serviceAccountIds = new InputList<string>());
            set => _serviceAccountIds = value;
        }

        public TwingateResourceAccessGetArgs()
        {
        }
        public static new TwingateResourceAccessGetArgs Empty => new TwingateResourceAccessGetArgs();
    }
}
