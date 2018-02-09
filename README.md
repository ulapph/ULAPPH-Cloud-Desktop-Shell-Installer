# ULAPPH-Cloud-Desktop-Shell-Installer
Basic installer for ULAPPH Cloud Desktop using Google Cloud Shell

## Pre-requisites
- Google Cloud Shell or local Unix/Linux access

## STEP 1
- Download the repo to your local computer
	https://github.com/ulapph/ULAPPH-Cloud-Desktop-Shell-Installer
```
	git clone https://github.com/ulapph/ULAPPH-Cloud-Desktop-Shell-Installer.git
	cd ULAPPH-Cloud-Desktop-Shell-Installer
	go install ulapphctl.go
	ulapphctl help
```

## STEP 2
- Once you have installed ulapphctl, run it by pointing to the configuration yaml file
```
	ulapphctl --config "../ULAPPH-Cloud-Desktop-Configs/edwin-daen-vinas.yaml" install
```
- Note that the recommended directory below
```
cd /c/Development/golang/ulapph

$ ls -la
total 894
drwxr-xr-x 1 edwin.d.vinas 1049089      0 Feb 10 07:10  ULAPPH-Cloud-Desktop-1/
drwxr-xr-x 1 edwin.d.vinas 1049089      0 Feb 10 07:01  ULAPPH-Cloud-Desktop-Configs/
drwxr-xr-x 1 edwin.d.vinas 1049089      0 Feb 10 06:59  ULAPPH-Cloud-Desktop-Shell-Installer/
```
- This assumes that:
	- The cloned ULAPPH Cloud Desktop source codes are in "ULAPPH-Cloud-Desktop-1"
	- The configuration files are in "ULAPPH-Cloud-Desktop-Configs"
	- The cloned shell installer are in "ULAPPH-Cloud-Desktop-Shell-Installer"

## Contact
- Gmail account: edwin.d.vinas@gmail.com
