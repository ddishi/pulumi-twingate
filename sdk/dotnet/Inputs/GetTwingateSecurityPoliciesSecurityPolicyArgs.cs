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

    public sealed class GetTwingateSecurityPoliciesSecurityPolicyInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Return a matching Security Policy by its ID. The ID for the Security Policy must be obtained from the Admin API.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// Return a Security Policy that exactly matches this name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetTwingateSecurityPoliciesSecurityPolicyInputArgs()
        {
        }
        public static new GetTwingateSecurityPoliciesSecurityPolicyInputArgs Empty => new GetTwingateSecurityPoliciesSecurityPolicyInputArgs();
    }
}
