
# ULAPPH-Cloud-Desktop-Shell-Installer
Basic installer for ULAPPH Cloud Desktop using Google Cloud Shell

## For most updated ULAPPH Cloud Desktop installation guide, refer to the link below.
* https://github.com/edwindvinas/ULAPPH-Cloud-Desktop/blob/master/DOCS/Installation%20Guide%20using%20Shell.md

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
	go get github.com/jinzhu/configor
	go get github.com/urfave/cli
	
	export GOBIN=/home/ulapph/gopath/bin
	
	go install ulapphctl.go
	which ulapphctl
	/home/ulapph/gopath/bin/ulapphctl

	** if go install does not work
	go build ulapphctl.go
	which ulapphctl
	cp ulapphctl <location of Go bin>

	ulapphctl help
	
	NAME:
	   ulapphctl - A new cli application

	USAGE:
	   ulapphctl [global options] command [command options] [arguments...]

	VERSION:
	   0.0.0

	COMMANDS:
	     configure, i  configure ulapph cloud desktop
	     deploy, i     deploy ulapph cloud desktop
	     help, h       Shows a list of commands or help for one command

	GLOBAL OPTIONS:
	   --account value, -a value  Google account (email)
	   --config value, -c value   Configuration file for the ulapph cloud destkop
	   --project value, -p value  Target google project ID
	   --yaml value, -y value     YAML source file for Google Appengine
	   --help, -h                 show help
	   --version, -v              print the version
   
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

## STEP 4
- If you have no YAML for your project, you may copy the ulapph-demo.yaml
	* https://github.com/ulapph/ULAPPH-Cloud-Desktop-Shell-Installer/blob/master/ulapph-demo.yaml
```
	wget https://raw.githubusercontent.com/ulapph/ULAPPH-Cloud-Desktop-Shell-Installer/master/ulapph-demo.yaml
	cp ulapph-demo.yaml your-project-id.yaml
```

## What if I need to install/upgrade multiple projects?
- Yes, you can use the quick install script below by indicating the list of all project IDs to be installed.
	https://github.com/ulapph/quick-install-ulapph
## Is there a quick way to code/edit and re-install ULAPPH?
- Yes, you can use the quick install script below. It only takes one command to install a new code.
	https://github.com/ulapph/quick-install-ulapph

## Contacts
- Gmail account: edwin.d.vinas@gmail.com
