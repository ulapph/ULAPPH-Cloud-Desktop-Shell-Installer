
# ULAPPH-Cloud-Desktop-Shell-Installer
Basic installer for ULAPPH Cloud Desktop using Google Cloud Shell

## Pre-requisites
- Google Cloud Shell or local Unix/Linux access
- A Google account such as Gmail
- A Google cloud project ID

## STEP 1
- For the ULAPPH cloud desktop source codes repo to your Github account
	- https://github.com/accenture/ULAPPH-Cloud-Desktop#fork-destination-box
- If you have existing repo or dont want to fork, go to step 2

## STEP 2
- Download the shell installer repo to your local computer
	- https://github.com/ulapph/ULAPPH-Cloud-Desktop-Shell-Installer
```
	git clone https://github.com/ulapph/ULAPPH-Cloud-Desktop-Shell-Installer.git
	cd ULAPPH-Cloud-Desktop-Shell-Installer
	go install ulapphctl.go
	ulapphctl help
```

## STEP 3
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
	- If you have cloned the source codes, you may just have "ULAPPH-Cloud-Desktop"
	- You have created a directory "ULAPPH-Cloud-Desktop-Configs" where you will put the yaml files
	- The cloned shell installer are in "ULAPPH-Cloud-Desktop-Shell-Installer"

## Contacts
- Gmail account: edwin.d.vinas@gmail.com
