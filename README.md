# cloud-director-cli

Simple CLI tool for interfacing with Cloud Director.

## Installation

### Brew

```bash
brew tap giantswarm/cd-cli
brew install cd-cli
```

Update to the latest version: 

```bash
brew update && brew upgrade cd-cli
```

Or alternatively download the latest binary for your platform at https://github.com/giantswarm/cloud-director-cli/releases
and put it to your `$PATH`.

### Config

```bash
cat ~/.cd-cli/config.yaml
refreshToken: ***
site: https://vmware.ikoula.com
org: giantswarm
ovdc: vDC 73640
```

HINT: when using w/ CAPVCD provider, you can get the `refreshToken` from k8s secret 

```bash
kubectl get secret refresh-token-secret -n org-multi-project -o jsonpath='{.data .refreshToken}{"\n"}' | base64 --decode
```

## Examples

In general it is `cd-cli ${verb} ${resource} ${params}`.
where `${verb}` can be `list` or `clean|delete`
`${resource}` can be `vm(s), vapp(s)|virtualapp(s), disk(s), vs(s)|virtualservice(s)|virtualsvc(s)|vsvc(s)|vsvcs`

### VMs

List all the VMs:

```bash
cd-cli list vms
guppy-8fb68
guppy-w4chm
guppy-worker-79fbbb5b7c-9mvpm
guppy-worker-79fbbb5b7c-9g7gx
guppy-s8k46
guppy-worker-79fbbb5b7c-sxzkq
squid-proxy
```

Verbose version:

```bash
cd-cli list vms -v
NAME                               	VAPP            	STATUS    	DEPLOYED
guppy-8fb68                        	guppy           	POWERED_ON	true
guppy-w4chm                        	guppy           	POWERED_ON	true
guppy-worker-79fbbb5b7c-9mvpm      	guppy           	POWERED_ON	true
guppy-worker-79fbbb5b7c-9g7gx      	guppy           	POWERED_ON	true
guppy-s8k46                        	guppy           	POWERED_ON	true
guppy-worker-79fbbb5b7c-sxzkq      	guppy           	POWERED_ON	true
squid-proxy                        	installation-proxy	POWERED_ON	true
```

Get VMs (names only) of given vApp only:

```bash
cd-cli list vms -a installation-proxy
squid-proxy
```

Delete individual VMs:

```bash
cd-cli clean vms --vapp=jiri3 jiri3-worker-7b4d46494-8rj59 jiri3-worker-7b4d46494-p6vhp
Are you sure you want to delete following VMs: [jiri3-worker-7b4d46494-8rj59, jiri3-worker-7b4d46494-p6vhp] [y/n]?
y
```

### vApps

List vApps:
```bash
cd-cli list vapp -v
NAME                               	ID
guppy                              	urn:vcloud:vapp:afe1a37f-4b7d-4c0f-a5f3-14f19bf5f073
installation-proxy                 	urn:vcloud:vapp:8994a22f-4870-43d4-8897-6945f2e96d9b
gs-eric-vcd                        	urn:vcloud:vapp:26f79f84-908b-4ee8-88a9-36d5066175f8
```

Delete whole vApp called `jiri3` and its associated VMs:

```bash
cd-cli clean vapp jiri3 --asumeyes
```

### Disks

List disks:
```bash
cd-cli list disks -v
NAME                                         	SIZE(Mb)  	STATUS    	VMs	TYPE
pvc-69969a35-b9df-4605-b052-d60beabf0d20     	5120      	RESOLVED  	0	Paravirtual (SCSI)
pvc-37eef8f3-8708-40fb-b4c3-6d6cc3e0a760     	1024      	RESOLVED  	0	Paravirtual (SCSI)
pvc-5add9939-513c-4017-a76b-927221881ac1     	1024      	RESOLVED  	0	Paravirtual (SCSI)
pvc-f197529a-2e79-43ea-a910-338658d217d1     	102400    	RESOLVED  	0	Paravirtual (SCSI)
pvc-eb6062e3-8c1d-4cf1-8406-f8463dd4a1dd     	102400    	RESOLVED  	1	Paravirtual (SCSI)
pvc-522bdd65-fc59-4769-8606-5d328af48eb1     	5120      	RESOLVED  	1	Paravirtual (SCSI)
pvc-c6984359-97ac-4e72-b4f0-b7ef9531b3e1     	1024      	RESOLVED  	1	Paravirtual (SCSI)
pvc-7da07a37-c4a1-4e8e-8de6-7cf24278cfc0     	5120      	RESOLVED  	0	Paravirtual (SCSI)
pvc-60b68772-bdb9-49bb-95f9-2f49b6972c90     	102400    	RESOLVED  	0	Paravirtual (SCSI)
```

delete:

```bash
cd-cli delete disks sdf1 sdf2 -y
```

### Virtual Services

Verbose listing:

```bash
cd-cli list vs -v
NAME                                                                                      	IP               	HEALTH
gs-eric-vcd-NO_RDE_b03a4df5-585f-48a9-8916-d378c44b7c16-tcp                               	178.170.32.55    	UP
ingress-vs-nginx-ingress-controller-app-NO_RDE_b03a4df5-585f-48a9-8916-d378c44b7c16-http  	192.168.8.6      	UP
ingress-vs-nginx-ingress-controller-app-NO_RDE_b03a4df5-585f-48a9-8916-d378c44b7c16-https 	192.168.8.7      	UP
ingress-vs-nginx-ingress-controller-app--http                                             	192.168.8.4      	UP
ingress-vs-nginx-ingress-controller-app--https                                            	192.168.8.5      	UP
guppy-NO_RDE_ca501275-f986-4d50-a6ec-e084341d15d2-tcp                                     	178.170.32.23    	UP
ingress-vs-nginx-ingress-controller-app-NO_RDE_ca501275-f986-4d50-a6ec-e084341d15d2-http  	192.168.8.2      	UP
ingress-vs-nginx-ingress-controller-app-NO_RDE_ca501275-f986-4d50-a6ec-e084341d15d2-https 	192.168.8.3      	UP
```

deleting:

```bash
cd-cli delete vs sdf --failifabsent
2022/12/08 11:53:23 virtual Service [sdf] does not exist
exit status 1
```

```bash
cd-cli delete vs guppy-NO_RDE_ca501275-f986-4d50-a6ec-e084341d15d2-tcp  --assumeyes
```

### LB Pools

Verbose listing:

```bash
cd-cli list lbps -v
NAME                                                                                      	ALGOTITHM        	MEMBERS
ingress-pool-nginx-ingress-controller-app--http                                           	LEAST_CONNECTIONS	6
ingress-pool-nginx-ingress-controller-app--https                                          	LEAST_CONNECTIONS	6
gs-eric-vcd-NO_RDE_b03a4df5-585f-48a9-8916-d378c44b7c16-tcp                               	ROUND_ROBIN      	1
guppy-NO_RDE_ca501275-f986-4d50-a6ec-e084341d15d2-tcp                                     	ROUND_ROBIN      	3
```

deleting:

```bash
cd-cli delete lbp sdf1 sdf2 sdf3 -y
```
