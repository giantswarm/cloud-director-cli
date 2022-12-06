# cloud-director-cli

Simple CLI tool for interfacing with Cloud Director.

## Installation

## Examples

List all the VMs:

```bash
λ cd-cli list vms                                                                                       5s
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
λ cd-cli list vms -v
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
λ cd-cli list vms -a installation-proxy
squid-proxy
```

List vApps:
```bash
λ cd-cli list vapp -v
NAME                               	ID
guppy                              	urn:vcloud:vapp:afe1a37f-4b7d-4c0f-a5f3-14f19bf5f073
installation-proxy                 	urn:vcloud:vapp:8994a22f-4870-43d4-8897-6945f2e96d9b
gs-eric-vcd                        	urn:vcloud:vapp:26f79f84-908b-4ee8-88a9-36d5066175f8
```

List disks:
```bash
λ cd-cli list disks -v
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

Delete whole vApp and its associated VMs:

```bash
λ cd-cli clean vapp jiri3
```

or delete individual VMs:

```bash
λ cd-cli clean vms --vapp=jiri3 jiri3-worker-7b4d46494-8rj59 jiri3-worker-7b4d46494-p6vhp
Are you sure you want to delete following VMs: [jiri3-worker-7b4d46494-8rj59, jiri3-worker-7b4d46494-p6vhp] [y/n]?
y
```