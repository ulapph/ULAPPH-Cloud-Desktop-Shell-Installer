
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
	go get github.com/jinzhu/configor
	go get github.com/urfave/cli
	
	export GOBIN=/home/ulapph/gopath/bin
	
	go install ulapphctl.go
	which ulapphctl
	/home/ulapph/gopath/bin/ulapphctl
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
## STEP 4a - Populate required YAML fields
- Populate the below minimum required fields
- Indicate your project ID
```
project:
   - name: ULAPPH Cloud Desktop
     appid: ulapph-demo <-- your project ID
```

- Indicate the path to the ULAPPH Cloud Desktop installer
```
installer:
   - dir: ../ULAPPH-Cloud-Desktop-1
```

- Indicate the application URL
```
configs: 
   - item: APP_URL
     format: Text
     status: Enable
     value: ulapph-demo.appspot.com
```

- Indicate the server name
```
   - item: SYS_SERVER_NAME
     format: Text
     status: Enable
     value: ulapph-demo
```
- Indicate the Google URL Shortener API key
- You can get the API Key from here
	* https://developers.google.com/url-shortener/v1/getting_started
```
   - item: apiKeyUs
     format: Text
     status: Enable
     value: AIzaSyDY93rCNZv_IXLUaz0aRWhX61234567890
```

## STEP 4b - Populate optional YAML fields
- Populate the below optional fields
- Indicate the server description
```
   - item: APP_DESC
     format: Text
     status: Enable
     value: ULAPPH Demo
```

- Indicate if the desktop is public or private
```
   - item: SYS_SITE_PRIVATE
     format: Flag
     status: Enable
     value: true
```

- Indicate if users can register to your desktop or manual only
```
   - item: SYS_REGISTRATION_MANUAL
     format: Flag
     status: Enable
     value: true

   - item: SYS_AUTO_REG_ENABLE
     format: Flag
     status: Enable
     value: false
```

- Indicate the valid email account
```
   - item: ADMIN_ACCOUNT
     format: Text
     status: Enable
     value: demo.ulapph@gmail.com

   - item: EMAIL_ADD_1
     format: Text
     status: Enable
     value: demo.ulapph@gmail.com

   - item: EMAIL_ADD_2
     format: Text
     status: Enable
     value: demo.ulapph@gmail.com
```

- Also indicate the email accounts here
```
   - item: ADMMAIL
     format: Text
     status: Enable
     value: demo.ulapph@gmail.com

   - item: SYSMAIL
     format: Text
     status: Enable
     value: demo.ulapph@gmail.com

   - item: ADSMAIL
     format: Text
     status: Enable
     value: demo.ulapph@gmail.com

   - item: REPMAIL
     format: Text
     status: Enable
     value: demo.ulapph@gmail.com

   - item: FDBKMAIL
     format: Text
     status: Enable
     value: demo.ulapph@gmail.com
```

- Indicate which search engines can search your desktop
```
   - item: var isSearchEngineAllowed
     format: Text
     status: InternalMatched
     value: map[string]bool{"US.?.?":false,"CN.?.?":false,"RU.?.?":false,"UA.?.?":false,"MX.?.?":false,"TM.?.?":false,}
```

- Indicate which countries can access your desktop
```
   - item: var isCountryNotAllowed
     format: Text
     status: InternalMatched
     value: map[string]bool{"RU":false,}

   - item: var isCountryAllowed
     format: Text
     status: InternalMatched
     value: map[string]bool{"PH":true,"US":true,}
```

- Indicate your cloud website meta details
```
   - item: <title>
     format: Text
     status: InternalMatched
     value: ULAPPH - Demo - ULAPPH Cloud Desktop

   - item: <meta description>
     format: Text
     status: InternalMatched
     value: ULAPPH Cloud Desktop of Demo Project

   - item: <meta keywords>
     format: Text
     status: InternalMatched
     value: ULAPPH - Demo Cloud Desktop

   - item: INFO_ABOUT_US
     format: Text
     status: Enable
     value: This is a demo project powered by ULAPPH Cloud Desktop!

   - item: SITE_SLOGAN
     format: Text
     status: Enable
     value: We will only live once in this world. Anything we need to do must be done. So let us do it now!
```

- Other configurations are optional...

## Contacts
- Gmail account: edwin.d.vinas@gmail.com
