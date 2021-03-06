# RancherOS

The smallest, easiest way to run Docker in production at scale.  Everything in RancherOS is a container managed by Docker.  This includes system services such as udev and rsyslog.  RancherOS includes only the bare minimum amount of software needed to run Docker.  This keeps the binary download of RancherOS very small.  Everything else can be pulled in dynamically through Docker.

## How this works

Everything in RancherOS is a Docker container.  We accomplish this by launching two instances of
Docker.  One is what we call the system Docker which runs as the first process.  System Docker then launches
a container that runs the user Docker.  The user Docker is then the instance that gets primarily
used to create containers.  We created this separation because it seemed logical and also
it would really be bad if somebody did `docker rm -f $(docker ps -qa)` and deleted the entire OS.

![How it works](docs/rancheros.png "How it works")

## Latest Release

**v1.0.1 - Docker 17.03.1-ce - Linux 4.9.24**

### ISO

- https://releases.rancher.com/os/latest/rancheros.iso
- https://releases.rancher.com/os/v1.0.1/rancheros.iso

### Additional Downloads

#### Latest Links

* https://releases.rancher.com/os/latest/initrd
* https://releases.rancher.com/os/latest/initrd-v1.0.1
* https://releases.rancher.com/os/latest/iso-checksums.txt
* https://releases.rancher.com/os/latest/rancheros-openstack.img
* https://releases.rancher.com/os/latest/rancheros.ipxe
* https://releases.rancher.com/os/latest/rancheros.iso
* https://releases.rancher.com/os/latest/rancheros-v1.0.1.tar.gz
* https://releases.rancher.com/os/latest/rootfs.tar.gz
* https://releases.rancher.com/os/latest/vmlinuz
* https://releases.rancher.com/os/latest/vmlinuz-4.9.24-rancher

#### v1.0.1 Links

* https://releases.rancher.com/os/v1.0.1/initrd
* https://releases.rancher.com/os/v1.0.1/initrd-v1.0.1
* https://releases.rancher.com/os/v1.0.1/iso-checksums.txt
* https://releases.rancher.com/os/v1.0.1/rancheros-openstack.img
* https://releases.rancher.com/os/v1.0.1/rancheros.ipxe
* https://releases.rancher.com/os/v1.0.1/rancheros.iso
* https://releases.rancher.com/os/v1.0.1/rancheros-v1.0.1.tar.gz
* https://releases.rancher.com/os/v1.0.1/rootfs.tar.gz
* https://releases.rancher.com/os/v1.0.1/vmlinuz
* https://releases.rancher.com/os/v1.0.1/vmlinuz-4.9.24-rancher

#### v1.0.0 Links - updates coming soon

* https://releases.rancher.com/os/latest/rootfs_arm.tar.gz
* https://releases.rancher.com/os/latest/rootfs_arm64.tar.gz
* https://releases.rancher.com/os/latest/rancheros-raspberry-pi.zip

* https://releases.rancher.com/os/v1.0.0/rootfs_arm.tar.gz
* https://releases.rancher.com/os/v1.0.0/rootfs_arm64.tar.gz
* https://releases.rancher.com/os/v1.0.0/rancheros-raspberry-pi.zip

**Note**: you can use `http` instead of `https` in the above URLs, e.g. for iPXE.

### Amazon

SSH keys are added to the **`rancher`** user, so you must log in using the **rancher** user.

**HVM**

Region | Type | AMI |
-------|------|------
ap-south-1 | HVM | [ami-9b4e3cf4](https://ap-south-1.console.aws.amazon.com/ec2/home?region=ap-south-1#launchInstanceWizard:ami=ami-9b4e3cf4)
eu-west-2 | HVM | [ami-d8eafebc](https://eu-west-2.console.aws.amazon.com/ec2/home?region=eu-west-2#launchInstanceWizard:ami=ami-d8eafebc)
eu-west-1 | HVM | [ami-75cbcb13](https://eu-west-1.console.aws.amazon.com/ec2/home?region=eu-west-1#launchInstanceWizard:ami=ami-75cbcb13)
ap-northeast-2 | HVM | [ami-797cae17](https://ap-northeast-2.console.aws.amazon.com/ec2/home?region=ap-northeast-2#launchInstanceWizard:ami=ami-797cae17)
ap-northeast-1 | HVM | [ami-e499b383](https://ap-northeast-1.console.aws.amazon.com/ec2/home?region=ap-northeast-1#launchInstanceWizard:ami=ami-e499b383)
sa-east-1 | HVM | [ami-f4cca198](https://sa-east-1.console.aws.amazon.com/ec2/home?region=sa-east-1#launchInstanceWizard:ami=ami-f4cca198)
ca-central-1 | HVM | [ami-59b60a3d](https://ca-central-1.console.aws.amazon.com/ec2/home?region=ca-central-1#launchInstanceWizard:ami=ami-59b60a3d)
ap-southeast-1 | HVM | [ami-be853edd](https://ap-southeast-1.console.aws.amazon.com/ec2/home?region=ap-southeast-1#launchInstanceWizard:ami=ami-be853edd)
ap-southeast-2 | HVM | [ami-e2dcd481](https://ap-southeast-2.console.aws.amazon.com/ec2/home?region=ap-southeast-2#launchInstanceWizard:ami=ami-e2dcd481)
eu-central-1 | HVM | [ami-ef0bd780](https://eu-central-1.console.aws.amazon.com/ec2/home?region=eu-central-1#launchInstanceWizard:ami=ami-ef0bd780)
us-east-1 | HVM | [ami-5c5a3f4a](https://us-east-1.console.aws.amazon.com/ec2/home?region=us-east-1#launchInstanceWizard:ami=ami-5c5a3f4a)
us-east-2 | HVM | [ami-902304f5](https://us-east-2.console.aws.amazon.com/ec2/home?region=us-east-2#launchInstanceWizard:ami=ami-902304f5)
us-west-1 | HVM | [ami-4a73542a](https://us-west-1.console.aws.amazon.com/ec2/home?region=us-west-1#launchInstanceWizard:ami=ami-4a73542a)
us-west-2 | HVM | [ami-7caa341c](https://us-west-2.console.aws.amazon.com/ec2/home?region=us-west-2#launchInstanceWizard:ami=ami-7caa341c)


### Google Compute Engine

We are providing a disk image that users can download and import for use in Google Compute Engine. The image can be obtained from the release artifacts for RancherOS.

[Download Image](https://github.com/rancher/os/releases/download/v1.0.0/rancheros-v1.0.0.tar.gz)

Please follow the directions at our [docs to launch in GCE](http://docs.rancher.com/os/running-rancheros/cloud/gce/).

## Documentation for RancherOS

Please refer to our [RancherOS Documentation](http://docs.rancher.com/os/) website to read all about RancherOS. It has detailed information on how RancherOS works, getting-started and other details.

## Support, Discussion, and Community
If you need any help with RancherOS or Rancher, please join us at either our [Rancher forums](http://forums.rancher.com) or [#rancher IRC channel](http://webchat.freenode.net/?channels=rancher) where most of our team hangs out at.

For security issues, please email security@rancher.com instead of posting a public issue in GitHub.  You may (but are not required to) use the GPG key located on [Keybase](https://keybase.io/rancher).


Please submit any **RancherOS** bugs, issues, and feature requests to [rancher/os](//github.com/rancher/os/issues).

Please submit any **Rancher** bugs, issues, and feature requests to [rancher/rancher](//github.com/rancher/rancher/issues).

#License
Copyright (c) 2014-2017 [Rancher Labs, Inc.](http://rancher.com)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
