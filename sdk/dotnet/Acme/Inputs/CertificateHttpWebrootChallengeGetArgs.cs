// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ByteFoo.PulumiPackage.Acme.Inputs
{

    public sealed class CertificateHttpWebrootChallengeGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("directory", required: true)]
        public Input<string> Directory { get; set; } = null!;

        public CertificateHttpWebrootChallengeGetArgs()
        {
        }
        public static new CertificateHttpWebrootChallengeGetArgs Empty => new CertificateHttpWebrootChallengeGetArgs();
    }
}
