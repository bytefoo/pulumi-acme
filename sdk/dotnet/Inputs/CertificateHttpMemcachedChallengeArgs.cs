// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ByteFoo.Acme.Inputs
{

    public sealed class CertificateHttpMemcachedChallengeArgs : global::Pulumi.ResourceArgs
    {
        [Input("hosts", required: true)]
        private InputList<string>? _hosts;
        public InputList<string> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<string>());
            set => _hosts = value;
        }

        public CertificateHttpMemcachedChallengeArgs()
        {
        }
        public static new CertificateHttpMemcachedChallengeArgs Empty => new CertificateHttpMemcachedChallengeArgs();
    }
}