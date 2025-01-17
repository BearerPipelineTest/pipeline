// Copyright © 2020 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eks

// AMIs taken form https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html
// nolint: gochecknoglobals
var defaultImages = ImageSelectors{
	KubernetesVersionImageSelector{ // Kubernetes Version 1.19
		Constraint: mustConstraint("1.19"),
		ImageSelector: RegionMapImageSelector{
			"ap-northeast-1": "ami-071e81675987911dd",
			"ap-northeast-2": "ami-060ea45269718d267",
			"ap-northeast-3": "ami-0cd051b3e38aa5f1a",
			"ap-southeast-1": "ami-077640972b17421de",
			"ap-southeast-2": "ami-0bbd3c34240eb62f3",
			"ap-south-1":     "ami-041e3306226b8c8ab",
			"ca-central-1":   "ami-0e24377c244ee200b",
			"eu-central-1":   "ami-0d8faf43ff567c10f",
			"eu-north-1":     "ami-09e9309e94932dace",
			"eu-west-1":      "ami-0a8119c391d4c0679",
			"eu-west-2":      "ami-0cbed66ed952f31b3",
			"eu-west-3":      "ami-083a9af0576036a79",
			"sa-east-1":      "ami-085e140454b969e36",
			"us-east-1":      "ami-0babee45de9606b3f",
			"us-east-2":      "ami-0dfd6f52226bc33ea",
			"us-west-1":      "ami-057ac7ec608243d73",
			"us-west-2":      "ami-0f89b397be013a2d5",
		},
	},
	KubernetesVersionImageSelector{ // Kubernetes Version 1.20
		Constraint: mustConstraint("1.20"),
		ImageSelector: RegionMapImageSelector{
			"ap-northeast-1": "ami-0c4b454e21d5656cd",
			"ap-northeast-2": "ami-002b0026c5dd8dfa0",
			"ap-northeast-3": "ami-02da313aed06a5fac",
			"ap-southeast-1": "ami-0452d3cd8c67295ad",
			"ap-southeast-2": "ami-05a87a8bbb7ceaeee",
			"ap-south-1":     "ami-03b6dce2bfbdb3331",
			"ca-central-1":   "ami-0c07e1f2e78d30f21",
			"eu-central-1":   "ami-00dc179746934b861",
			"eu-north-1":     "ami-0e65a1b32e2257fb5",
			"eu-west-1":      "ami-08ce9adc1160e9c1d",
			"eu-west-2":      "ami-0add2c2a9497bf7d7",
			"eu-west-3":      "ami-07dfee37e3909a334",
			"sa-east-1":      "ami-030d972bbbb77d42d",
			"us-east-1":      "ami-0e1ce97e46e7a027f",
			"us-east-2":      "ami-0cd50b331314a4415",
			"us-west-1":      "ami-071badcf59d49d36b",
			"us-west-2":      "ami-082adc11c87e4a462",
		},
	},
	KubernetesVersionImageSelector{ // Kubernetes Version 1.21
		Constraint: mustConstraint("1.21"),
		ImageSelector: RegionMapImageSelector{
			"ap-northeast-1": "ami-06dc12568363f25b4",
			"ap-northeast-2": "ami-0d4bbf30d9907be82",
			"ap-northeast-3": "ami-083e2fa499e503354",
			"ap-southeast-1": "ami-03d9d62e920be1193",
			"ap-southeast-2": "ami-0a42ddf675da067c3",
			"ap-south-1":     "ami-08896ae019d7c2834",
			"ca-central-1":   "ami-0453fd0b36fbc1c1f",
			"eu-central-1":   "ami-0b4700a0adca90ae7",
			"eu-north-1":     "ami-0b71ebb38c4685c93",
			"eu-west-1":      "ami-0b8e12a0df2cb151e",
			"eu-west-2":      "ami-0b1aba45378946235",
			"eu-west-3":      "ami-089506e08f2707d6a",
			"sa-east-1":      "ami-072258e16c04c4418",
			"us-east-1":      "ami-0ff8d483010b138ac",
			"us-east-2":      "ami-07544b0fe2b0e4d9c",
			"us-west-1":      "ami-076a563eaac6ccd4f",
			"us-west-2":      "ami-05074c40f29040248",
		},
	},
	KubernetesVersionImageSelector{ // Kubernetes Version 1.22
		Constraint: mustConstraint("1.22"),
		ImageSelector: RegionMapImageSelector{
			"ap-northeast-1": "ami-0f54c217ddeb4226e",
			"ap-northeast-2": "ami-028f26a5a0798bf1e",
			"ap-northeast-3": "ami-0748d097b6cee692f",
			"ap-southeast-1": "ami-084775e47d5a5b85a",
			"ap-southeast-2": "ami-07e0a785ca42c0fe2",
			"ap-south-1":     "ami-05c200ff7ef5b2583",
			"ca-central-1":   "ami-06815037d723ac630",
			"eu-central-1":   "ami-03215683d5de485a1",
			"eu-north-1":     "ami-0802a34a268e26e5e",
			"eu-west-1":      "ami-06d695da292f6164b",
			"eu-west-2":      "ami-0270e03ea3eee9efe",
			"eu-west-3":      "ami-046238128ef2aee7c",
			"sa-east-1":      "ami-0e8885b0803dd261e",
			"us-east-1":      "ami-0cca3d8b6715c9f74",
			"us-east-2":      "ami-0d7d38acf70b3c067",
			"us-west-1":      "ami-0dac13a60aa388979",
			"us-west-2":      "ami-09f0868e933b1494d",
		},
	},
	KubernetesVersionImageSelector{ // Kubernetes Version 1.23
		Constraint: mustConstraint("1.23"),
		ImageSelector: RegionMapImageSelector{
			"ap-northeast-1": "ami-0955f5be044203d0f",
			"ap-northeast-2": "ami-0f8ec68dd2f634010",
			"ap-northeast-3": "ami-0a3f8ba15abacd307",
			"ap-southeast-1": "ami-00dece352acd1283b",
			"ap-southeast-2": "ami-0f161aae5b9a624ea",
			"ap-south-1":     "ami-00c0507b59bd0ad85",
			"ca-central-1":   "ami-0971752fc69d93e60",
			"eu-central-1":   "ami-0ba3a46c3614adf25",
			"eu-north-1":     "ami-0b93e4144abb2aea3",
			"eu-west-1":      "ami-0eba9c07a41b01c3b",
			"eu-west-2":      "ami-07193cdcbbe8a5bc5",
			"eu-west-3":      "ami-0b565ab15ac4d8f52",
			"sa-east-1":      "ami-050ebe789faaf75c0",
			"us-east-1":      "ami-028df9f7b798ba67e",
			"us-east-2":      "ami-01a1852b9546bb86d",
			"us-west-1":      "ami-02ab603f9007d425c",
			"us-west-2":      "ami-0943a3de8e3f14ef8",
		},
	},
}

// GPU accelerated AMIs taken form https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html
// nolint: gochecknoglobals
var defaultAcceleratedImages = ImageSelectors{
	KubernetesVersionImageSelector{ // Kubernetes Version 1.19
		Constraint: mustConstraint("1.19"),
		ImageSelector: RegionMapImageSelector{
			"ap-northeast-1": "ami-0c7c4193d6535c155",
			"ap-northeast-2": "ami-04b2f445a9ca9a173",
			"ap-northeast-3": "ami-028c58e773d9eca05",
			"ap-southeast-1": "ami-0ddbb9f51f86b24db",
			"ap-southeast-2": "ami-0b3969c3ba37e8aea",
			"ap-south-1":     "ami-026873a01e80f776d",
			"ca-central-1":   "ami-03e0c0f3c7c166678",
			"eu-central-1":   "ami-0ebdfc250970004d5",
			"eu-north-1":     "ami-033ddd0db17fa9978",
			"eu-west-1":      "ami-04e55cfd07a662844",
			"eu-west-2":      "ami-0223cb5de8b305ffb",
			"eu-west-3":      "ami-0994368073157041b",
			"sa-east-1":      "ami-0c3a70643d7ce04ca",
			"us-east-1":      "ami-001ae1aedf36858e2",
			"us-east-2":      "ami-09fb1d569f456c581",
			"us-west-1":      "ami-04c4fbd280a01b55b",
			"us-west-2":      "ami-0355445c252399726",
		},
	},
	KubernetesVersionImageSelector{ // Kubernetes Version 1.20
		Constraint: mustConstraint("1.20"),
		ImageSelector: RegionMapImageSelector{
			"ap-northeast-1": "ami-0d7e268b316d40d93",
			"ap-northeast-2": "ami-03335432a0b416070",
			"ap-northeast-3": "ami-0505750d08e4e943b",
			"ap-southeast-1": "ami-073edabb5382447de",
			"ap-southeast-2": "ami-0499fd8aeef67ae3b",
			"ap-south-1":     "ami-0faf73f4745e97714",
			"ca-central-1":   "ami-045ab9b897cd28957",
			"eu-central-1":   "ami-09b484756c56bb4d3",
			"eu-north-1":     "ami-056eca8ece163c78b",
			"eu-west-1":      "ami-043ab58a8182d611a",
			"eu-west-2":      "ami-05ba7dee6fe281234",
			"eu-west-3":      "ami-0252e7ed3ddc36496",
			"sa-east-1":      "ami-08415152aec3986f3",
			"us-east-1":      "ami-0da16979da16dc6dd",
			"us-east-2":      "ami-08d4fc083db40a96b",
			"us-west-1":      "ami-0e0c57d867270477a",
			"us-west-2":      "ami-089c7cfada4d9ecff",
		},
	},
	KubernetesVersionImageSelector{ // Kubernetes Version 1.21
		Constraint: mustConstraint("1.21"),
		ImageSelector: RegionMapImageSelector{
			"ap-northeast-1": "ami-0261319e810169ab2",
			"ap-northeast-2": "ami-0238f4e739e2f9ae5",
			"ap-northeast-3": "ami-03416bef648d47756",
			"ap-southeast-1": "ami-08502c7d9eb851607",
			"ap-southeast-2": "ami-002a9d26e43f80ea0",
			"ap-south-1":     "ami-0d91a8a963d6d23c8",
			"ca-central-1":   "ami-03ad584d5c61b90f6",
			"eu-central-1":   "ami-02c4cedf3335bf746",
			"eu-north-1":     "ami-097939547c4549c9e",
			"eu-west-1":      "ami-0aab8f5c9e79fb9d2",
			"eu-west-2":      "ami-085aeff26b7db1312",
			"eu-west-3":      "ami-0b16ff6ec60cca4d1",
			"sa-east-1":      "ami-016bdb254b981c062",
			"us-east-1":      "ami-002d7fc3119a9ee52",
			"us-east-2":      "ami-0d2d23f11e2d24616",
			"us-west-1":      "ami-0337e5043c29df5c4",
			"us-west-2":      "ami-09b89c11c40bde883",
		},
	},
	KubernetesVersionImageSelector{ // Kubernetes Version 1.22
		Constraint: mustConstraint("1.22"),
		ImageSelector: RegionMapImageSelector{
			"ap-northeast-1": "ami-05c74365f50b94028",
			"ap-northeast-2": "ami-0f835b76791fd6b39",
			"ap-northeast-3": "ami-097e83f383e808c39",
			"ap-southeast-1": "ami-0fa213b6146184add",
			"ap-southeast-2": "ami-0f5fdf1bdfc2274ac",
			"ap-south-1":     "ami-065f66734e4a6e06f",
			"ca-central-1":   "ami-0198635613fa39af1",
			"eu-central-1":   "ami-0bc396c4093a0bd1a",
			"eu-north-1":     "ami-0eec8b23bb385d833",
			"eu-west-1":      "ami-0a8beade9b259828d",
			"eu-west-2":      "ami-0cd114b5e65e4b7a1",
			"eu-west-3":      "ami-03ee748a112c6aa49",
			"sa-east-1":      "ami-02b977a8583e4c11f",
			"us-east-1":      "ami-04ed79e2602553b42",
			"us-east-2":      "ami-096913c42aefe2a7d",
			"us-west-1":      "ami-0b129acc5d1ddb1cc",
			"us-west-2":      "ami-0f8863a8011c5518a",
		},
	},
	KubernetesVersionImageSelector{ // Kubernetes Version 1.23
		Constraint: mustConstraint("1.23"),
		ImageSelector: RegionMapImageSelector{
			"ap-northeast-1": "ami-0e7c91b684b1f4c24",
			"ap-northeast-2": "ami-0da0aae3699001a81",
			"ap-northeast-3": "ami-0b875a20c5b841e2c",
			"ap-southeast-1": "ami-006d7fbf2a8b3250b",
			"ap-southeast-2": "ami-061766b022a142919",
			"ap-south-1":     "ami-0b05e1f637e7dd813",
			"ca-central-1":   "ami-00026971e3ca9c37c",
			"eu-central-1":   "ami-08cc993934c9c378c",
			"eu-north-1":     "ami-0ffb2b3cb45b7bf30",
			"eu-west-1":      "ami-091e7510354ceba93",
			"eu-west-2":      "ami-0ce37be2c8d4394ff",
			"eu-west-3":      "ami-0f836782b7318414c",
			"sa-east-1":      "ami-0f3afd73a2cc50c84",
			"us-east-1":      "ami-02c3ec3a581b5c9d0",
			"us-east-2":      "ami-05c95125ab29ba015",
			"us-west-1":      "ami-09f80355c7c1fb159",
			"us-west-2":      "ami-001c64ca38434e509",
		},
	},
}

// ARM architecture AMIs taken form https://docs.aws.amazon.com/eks/latest/userguide/eks-optimized-ami.html
// nolint: gochecknoglobals
var defaultARMImages = ImageSelectors{
	KubernetesVersionImageSelector{ // Kubernetes Version 1.19
		Constraint: mustConstraint("1.19"),
		ImageSelector: RegionMapImageSelector{
			"ap-northeast-1": "ami-08815d41753b3192b",
			"ap-northeast-2": "ami-0b11acb21540c4b57",
			"ap-northeast-3": "ami-0d35b5c947eaf1f6b",
			"ap-southeast-1": "ami-040c36e6158040aaf",
			"ap-southeast-2": "ami-07a2b627c668ffc23",
			"ap-south-1":     "ami-0a0c8ee754369b277",
			"ca-central-1":   "ami-08466e6a632d524fc",
			"eu-central-1":   "ami-0b9b4757d6ac27269",
			"eu-north-1":     "ami-051e85c7e3695a552",
			"eu-west-1":      "ami-0f250597cacfb554a",
			"eu-west-2":      "ami-0b44fe0c3cf25484f",
			"eu-west-3":      "ami-0607dadee6a5949e3",
			"sa-east-1":      "ami-07af803449294c372",
			"us-east-1":      "ami-06802ec97e993de1b",
			"us-east-2":      "ami-09c7e196efda02a38",
			"us-west-1":      "ami-0e1a1dc0adf06f293",
			"us-west-2":      "ami-0181d2c7386d906b0",
		},
	},
	KubernetesVersionImageSelector{ // Kubernetes Version 1.20
		Constraint: mustConstraint("1.20"),
		ImageSelector: RegionMapImageSelector{
			"ap-northeast-1": "ami-040d8e5747466e2eb",
			"ap-northeast-2": "ami-07d8ce28fed142034",
			"ap-northeast-3": "ami-0af5a56f11ef8d48a",
			"ap-southeast-1": "ami-02d8679b68b421c87",
			"ap-southeast-2": "ami-0348fcb76751f93d1",
			"ap-south-1":     "ami-07ae4595f0dac1aea",
			"ca-central-1":   "ami-0a03a30d310e62c45",
			"eu-central-1":   "ami-0abf8e51df7115461",
			"eu-north-1":     "ami-05b4473e8551fb9e5",
			"eu-west-1":      "ami-04b178d07a05fa3ac",
			"eu-west-2":      "ami-045098d08517abc36",
			"eu-west-3":      "ami-060111f7513470a8b",
			"sa-east-1":      "ami-0ca1128c720435ef1",
			"us-east-1":      "ami-0faf1bf5541c35791",
			"us-east-2":      "ami-043d05adbf4348fc8",
			"us-west-1":      "ami-05c027af24864c097",
			"us-west-2":      "ami-03ff439180f85ca2c",
		},
	},
	KubernetesVersionImageSelector{ // Kubernetes Version 1.21
		Constraint: mustConstraint("1.21"),
		ImageSelector: RegionMapImageSelector{
			"ap-northeast-1": "ami-0707dee5f3b4b2efc",
			"ap-northeast-2": "ami-03535232fdd65155f",
			"ap-northeast-3": "ami-0a6e4dadf20e3c998",
			"ap-southeast-1": "ami-0d72386000b1f5b92",
			"ap-southeast-2": "ami-0daa5193f2eafea94",
			"ap-south-1":     "ami-09db6d3cd89c5702c",
			"ca-central-1":   "ami-0845c6a5d4a8b8848",
			"eu-central-1":   "ami-0295954fe8e7f45e3",
			"eu-north-1":     "ami-024074f073e195067",
			"eu-west-1":      "ami-0cc1c1a4be567835e",
			"eu-west-2":      "ami-0514fc4bdf872a483",
			"eu-west-3":      "ami-09b9c55795a048220",
			"sa-east-1":      "ami-01f95cc635f935c04",
			"us-east-1":      "ami-0342f19a86783c50a",
			"us-east-2":      "ami-0c49e2aec47b65200",
			"us-west-1":      "ami-0b42e2d501434fd30",
			"us-west-2":      "ami-01874f45972c94558",
		},
	},
	KubernetesVersionImageSelector{ // Kubernetes Version 1.22
		Constraint: mustConstraint("1.22"),
		ImageSelector: RegionMapImageSelector{
			"ap-northeast-1": "ami-0868e6e662fa1f9c2",
			"ap-northeast-2": "ami-0b8afdd95b3e24414",
			"ap-northeast-3": "ami-0608530c5f9880d16",
			"ap-southeast-1": "ami-08ac8e8a1592e736c",
			"ap-southeast-2": "ami-08126366b8eab6720",
			"ap-south-1":     "ami-01f350da999a60c47",
			"ca-central-1":   "ami-022565bad39545d2b",
			"eu-central-1":   "ami-0817589dc9fcc47f9",
			"eu-north-1":     "ami-08a9a4b84db171ea3",
			"eu-west-1":      "ami-00844a41a91fbcc20",
			"eu-west-2":      "ami-04f4a0dfe1a17c879",
			"eu-west-3":      "ami-08adcf914bd823be3",
			"sa-east-1":      "ami-0563b6eef71c104c2",
			"us-east-1":      "ami-06879c0ace931603e",
			"us-east-2":      "ami-09837e0a5bfd6e8dd",
			"us-west-1":      "ami-0752d1dff0d69cbe3",
			"us-west-2":      "ami-0614df54c8693bfdc",
		},
	},
	KubernetesVersionImageSelector{ // Kubernetes Version 1.23
		Constraint: mustConstraint("1.23"),
		ImageSelector: RegionMapImageSelector{
			"ap-northeast-1": "ami-0225977b23e4bb228",
			"ap-northeast-2": "ami-04424b15127b2941a",
			"ap-northeast-3": "ami-083e8f5ff72a48160",
			"ap-southeast-1": "ami-04a491f635aa7eb73",
			"ap-southeast-2": "ami-05e19337f11e8f72a",
			"ap-south-1":     "ami-026caaff4b4d7b7ef",
			"ca-central-1":   "ami-0bed3e5e2aef3692f",
			"eu-central-1":   "ami-0038ffad1d638ccf6",
			"eu-north-1":     "ami-05b5c77bab81f86c9",
			"eu-west-1":      "ami-01308f2619f1d52a4",
			"eu-west-2":      "ami-03073d32f232bb16e",
			"eu-west-3":      "ami-02d13fa5963ed4642",
			"sa-east-1":      "ami-0bf8b040bcea204a8",
			"us-east-1":      "ami-0b9602b6b3859d6e7",
			"us-east-2":      "ami-08263de230c8341a1",
			"us-west-1":      "ami-00cd3b0e7f2c017e1",
			"us-west-2":      "ami-0be69580f9644ddf8",
		},
	},
}
