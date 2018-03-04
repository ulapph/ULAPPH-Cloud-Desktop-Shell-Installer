/////////////////////////////////////////////////////////////////////////////////////////////////
// ULAPPH CLOUD DESKTOP SYSTEM
// ULAPPH Cloud Desktop is a web-based desktop that runs on Google cloud platform and accessible via browsers on different PC and mobile devices.
// COPYRIGHT (c) 2014-2017 Edwin D. Vinas, Ulapph Cloud Desktop System
// COPYRIGHT (c) 2017-2018 Accenture, Opensource Version
/////////////////////////////////////////////////////////////////////////////////////////////////
//REV ID: 	D0001
//REV DATE: 	2017-Feb-10
//REV DESC:	Created initial installer via Google Cloud Shell
//REV AUTH:	Edwin D. Vinas
//REV_REF:	https://github.com/jinzhu/configor
/////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////
//REV ID: 	D0002
//REV DATE: 	2017-Feb-13
//REV DESC:	Added developer commands
//REV AUTH:	Edwin D. Vinas
/////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////
//REV ID: 	D0003
//REV DATE: 	2017-Feb-21
//REV DESC:	Added developer commands
//REV AUTH:	Edwin D. Vinas
/////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////
// ulapphctl --config "../ULAPPH-Cloud-Desktop-Configs/edwin-daen-vinas.yaml" install
/////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"github.com/jinzhu/configor"
	"os/exec"
	"os"
	"log"
	"bufio"
	"strings"
	"bytes"
	"io/ioutil"
	"time"
	"strconv"
	"github.com/urfave/cli"
	"sort"
	"gopkg.in/cheggaaa/pb.v1"
)

var Config = struct {
	Project []struct {
		Name  string
		Date string
		Appid string
	}
	Installer []struct {
		Dir  string
	}
	Configs []struct {
		Item  string
		Format string
		Status string
		Value string
	}
}{}

var OLD_DOMAIN_NAME = ""
var NEW_DOMAIN_NAME = ""

func main() {
	var config string
	var project string
	var account string
	var yaml string
	app := cli.NewApp()
	app.Flags = []cli.Flag {
		cli.StringFlag{
		  Name: "config, c",
		  Usage: "Configuration file for the ulapph cloud destkop",
		  Destination: &config,
		},
		cli.StringFlag{
		  Name: "project, p",
		  Usage: "Target google project ID",
		  Destination: &project,
		},
		cli.StringFlag{
		  Name: "account, a",
		  Usage: "Google account (email)",
		  Destination: &account,
		},
		cli.StringFlag{
		  Name: "yaml, y",
		  Usage: "YAML source file for Google Appengine",
		  Destination: &yaml,
		},
	}

	app.Commands = []cli.Command{
	{
	  Name:    "configure",
	  Aliases: []string{"i"},
	  Usage:   "configure ulapph cloud desktop",
	  Action:  func(c *cli.Context) error {
		if config == "" {
			fmt.Printf("ERROR: Missing configuration file!")
			fmt.Printf("\nTry: ulapphctl configure --config your-ulapph-cloud-desktop.yaml\n")
			return nil
		}
		err := configureUlapphCloudDesktop(config)
		if err != nil {
			fmt.Printf("\nError(s) encountered! %v\n", err)	
		} else {
			fmt.Printf("\nCode has been generated!\n")
		}

		return nil
	  },
	},
	{
	  Name:    "deploy",
	  Aliases: []string{"i"},
	  Usage:   "deploy ulapph cloud desktop",
	  Action:  func(c *cli.Context) error {
		if account == "" {
			fmt.Printf("ERROR: Missing Google account parameter")
			fmt.Printf("\nTry: ulapphctl deploy --project your-ulapph-cloud-desktop --account demo.ulapph@gmail.com --yaml app.yaml\n")
			return nil
		}
		if project == "" {
			fmt.Printf("ERROR: Missing Project ID parameter")
			fmt.Printf("\nTry: ulapphctl deploy --project your-ulapph-cloud-desktop --account demo.ulapph@gmail.com --yaml app.yaml\n")
			return nil
		}
		if yaml == "" {
			fmt.Printf("ERROR: Missing YAML parameter")
			fmt.Printf("\nTry: ulapphctl deploy --project your-ulapph-cloud-desktop --account demo.ulapph@gmail.com --yaml app.yaml\n")
			return nil
		}
		err := deployUlapphCloudDesktop(project, account, yaml)
		if err != nil {
			fmt.Printf("\nError(s) encountered! %v\n", err)	
		} else {
			fmt.Printf("\nSuccessful!\n")
		}

		return nil
	  },
	},
	{
	  Name:    "redeploy",
	  Aliases: []string{"i"},
	  Usage:   "redeploy with same main.go",
	  Action:  func(c *cli.Context) error {
		if account == "" {
			fmt.Printf("ERROR: Missing Google account parameter")
			fmt.Printf("\nTry: ulapphctl deploy --project your-ulapph-cloud-desktop --account demo.ulapph@gmail.com --yaml app.yaml\n")
			return nil
		}
		if project == "" {
			fmt.Printf("ERROR: Missing Project ID parameter")
			fmt.Printf("\nTry: ulapphctl deploy --project your-ulapph-cloud-desktop --account demo.ulapph@gmail.com --yaml app.yaml\n")
			return nil
		}
		if yaml == "" {
			fmt.Printf("ERROR: Missing YAML parameter")
			fmt.Printf("\nTry: ulapphctl deploy --project your-ulapph-cloud-desktop --account demo.ulapph@gmail.com --yaml app.yaml\n")
			return nil
		}
		err := reDeployUlapphCloudDesktop(project, account, yaml)
		if err != nil {
			fmt.Printf("\nError(s) encountered! %v\n", err)	
		} else {
			fmt.Printf("\nSuccessful!\n")
		}

		return nil
	  },
	},{
	  Name:    "devcheck",
	  Aliases: []string{"i"},
	  Usage:   "for devs only -- check development",
	  //move backup to main.go
	  Action:  func(c *cli.Context) error {
		err := devCheckUlapphCloudDesktop()
		if err != nil {
			fmt.Printf("\nError(s) encountered! %v\n", err)	
		} else {
			fmt.Printf("\nDev files has been checked!\n")
		}

		return nil
	  },
	},{
	  Name:    "devstart",
	  Aliases: []string{"i"},
	  Usage:   "for devs only -- start development (copy main.go.dev to main.go)",
	  //move backup to main.go
	  Action:  func(c *cli.Context) error {
		err := devStartUlapphCloudDesktop()
		if err != nil {
			fmt.Printf("\nError(s) encountered! %v\n", err)	
		} else {
			fmt.Printf("\nDev files has been prepared! Go edit main.go\n")
		}

		return nil
	  },
	},{
	  Name:    "devstop",
	  Aliases: []string{"i"},
	  Usage:   "for devs only -- stop development (cp main.go main.go.dev)",
	  //backup to main.go.backup
	  Action:  func(c *cli.Context) error {
		err := devStopUlapphCloudDesktop()
		if err != nil {
			fmt.Printf("\nError(s) encountered! %v\n", err)	
		} else {
			fmt.Printf("\nDev files has been backed up!\n")
		}

		return nil
	  },
	},{
	  Name:    "devcommit",
	  Aliases: []string{"i"},
	  Usage:   "for devs only -- commit latest codes (check only)",
	  //backup to main.go.backup
	  Action:  func(c *cli.Context) error {
		err := devCommitUlapphCloudDesktop()
		if err != nil {
			fmt.Printf("\nError(s) encountered! %v\n", err)	
		} else {
			fmt.Printf("\nDev files ready to commit!\n")
		}

		return nil
	  },
	},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
  
}

//commit dev ulapph cloud desktop
func devCommitUlapphCloudDesktop() (err error) {
	fmt.Printf("\n+ ULAPPH Cloud Desktop Development\n")
	fmt.Printf("\n++ Prep env for commit...\n")

	fmt.Printf("\n++ Checking main.go\n")	
	//head -10 main.go
	app := "head"
	arg0 := fmt.Sprintf("-10")
	arg1 := fmt.Sprintf("main.go")
	stdout := []byte("ok")

	cmd := exec.Command(app, arg0, arg1)
	stdout, err = cmd.Output()

	if err != nil {
    	println(err.Error())
    	stdout = []byte(err.Error())
    	return
	}
	print(string(stdout))
	//check if above output is valid main.go dev file
	FL_VALID_DOM := false
	FL_VALID_DATE := false
    	scanner := bufio.NewScanner(bytes.NewReader(stdout))
    	for scanner.Scan() {
			sLineText := scanner.Text()		
			i := strings.Index(sLineText, "//GAE_APP_DOM_ID#")
			if i != -1 {
				FL_VALID_DOM = true
			}
			i = strings.Index(sLineText, "//LAST_UPGRADE#")
			if i != -1 {
				FL_VALID_DATE = true
			}
		}

		
	if FL_VALID_DOM != true && FL_VALID_DATE != true {
		print(string("ERROR: invalid main.go"))
		return
	} else {
		print(string("OK: valid main.go You can manually remove other main.* files and execute commamds!"))
		print(string("\nrm main2*"))
		print(string("\nrm main.go.backup"))
		print(string("\ngit add --all"))
		print(string("\ngit commit -m MESSAGE"))
		print(string("\ngit push origin master"))
	}	

	
	return err
}

//check dev ulapph cloud desktop
func devCheckUlapphCloudDesktop() (err error) {
	fmt.Printf("\n+ ULAPPH Cloud Desktop Development\n")
	fmt.Printf("\n++ Check env for development...\n")
	
	//ls -la
	app := "ls"
	arg0 := fmt.Sprintf("-la")
	stdout := []byte("ok")

	cmd := exec.Command(app, arg0)
	stdout, err = cmd.Output()

	if err != nil {
    	println(err.Error())
    	stdout = []byte(err.Error())
    	return
	}
	print(string(stdout))	

	fmt.Printf("\n++ Check main.go\n")	
	//ls main.go
	app = "ls"
	arg1 := fmt.Sprintf("main.go")
	stdout = []byte("ok")

	cmd = exec.Command(app, arg1)
	stdout, err = cmd.Output()

	if err != nil {
    	println(err.Error())
    	stdout = []byte("ERROR: main.go")
    	return
	}
	//print(string(stdout))
	
	//head -10 main.go
	app = "head"
	arg0 = fmt.Sprintf("-10")
	arg1 = fmt.Sprintf("main.go")
	stdout = []byte("ok")

	cmd = exec.Command(app, arg0, arg1)
	stdout, err = cmd.Output()

	if err != nil {
    	println(err.Error())
    	stdout = []byte("ERROR: main.go")
    	return
	}
	print(string(stdout))		
	
	fmt.Printf("\n++ Check main2.go\n")	
	//ls main2.go
	app = "ls"
	arg1 = fmt.Sprintf("main2.go")
	stdout = []byte("ok")

	cmd = exec.Command(app, arg1)
	stdout, err = cmd.Output()

	if err != nil {
    	println(err.Error())
    	stdout = []byte("WARNING: main2.go not found")
    	//return
	} else {
		//head -10 main2.go
		app = "head"
		arg0 = fmt.Sprintf("-10")
		arg1 = fmt.Sprintf("main2.go")
		stdout = []byte("ok")
	
		cmd = exec.Command(app, arg0, arg1)
		stdout, err = cmd.Output()
	
		if err != nil {
	    	println(err.Error())
	    	stdout = []byte("ERROR: main2.go")
	    	return
		}
		//print(string(stdout))	
	}
	print(string(stdout))	
	
	
	fmt.Printf("\n++ Check main.go.backup\n")	
	//ls main.go.backup
	app = "ls"
	arg1 = fmt.Sprintf("main.go.backup")
	stdout = []byte("ok")

	cmd = exec.Command(app, arg1)
	stdout, err = cmd.Output()

	if err != nil {
    	println(err.Error())
    	stdout = []byte("ERROR: main.go.backup not found!")
    	//return
	} else {
		//head -10 main.go.backup
		app = "head"
		arg0 = fmt.Sprintf("-10")
		arg1 = fmt.Sprintf("main.go.backup")
		stdout = []byte("ok")
	
		cmd = exec.Command(app, arg0, arg1)
		stdout, err = cmd.Output()
	
		if err != nil {
	    	println(err.Error())
	    	stdout = []byte("ERROR: main.go.backup")
	    	return
		}
		//print(string(stdout))	
	}
	print(string(stdout))	
	
	return err
}

//start dev ulapph cloud desktop
func devStartUlapphCloudDesktop() (err error) {
	fmt.Printf("\n+ ULAPPH Cloud Desktop Development\n")
	fmt.Printf("\n++ Prep env for development...\n")

	fmt.Printf("\n++ Checking main.go.dev\n")	
	//head -10 main.go
	app := "head"
	arg0 := fmt.Sprintf("-10")
	arg1 := fmt.Sprintf("main.go.dev")
	stdout := []byte("ok")

	cmd := exec.Command(app, arg0, arg1)
	stdout, err = cmd.Output()

	if err != nil {
    	println(err.Error())
    	stdout = []byte(err.Error())
    	return
	}
	print(string(stdout))
	//check if above output is valid main.go dev file
	FL_VALID_DOM := false
	FL_VALID_DATE := false
    	scanner := bufio.NewScanner(bytes.NewReader(stdout))
    	for scanner.Scan() {
			sLineText := scanner.Text()		
			i := strings.Index(sLineText, "//GAE_APP_DOM_ID#")
			if i != -1 {
				FL_VALID_DOM = true
			}
			i = strings.Index(sLineText, "//LAST_UPGRADE#")
			if i != -1 {
				FL_VALID_DATE = true
			}
		}

		
	if FL_VALID_DOM != true && FL_VALID_DATE != true {
		print(string("ERROR: invalid main.go.dev"))
		return
	} else {
		print(string("OK: valid main.go.dev"))
	}

	fmt.Printf("\n++ Copying main.go.dev to main.go\n")	
	//cp main.go.dev
	app = "cp"
	arg1 = fmt.Sprintf("main.go.dev")
	arg2 := fmt.Sprintf("main.go")
	stdout = []byte("ok")

	cmd = exec.Command(app, arg1, arg2)
	stdout, err = cmd.Output()

	if err != nil {
    	println(err.Error())
    	stdout = []byte(err.Error())
    	return
	}
	print(string(stdout))	
	
	return err
}

//stop dev ulapph cloud desktop
func devStopUlapphCloudDesktop() (err error) {
	fmt.Printf("\n+ ULAPPH Cloud Desktop Development\n")
	fmt.Printf("\n++ Backup env for pre-deployment...\n")
	
	fmt.Printf("\n++ Checking main.go\n")	
	//head -10 main.go
	app := "head"
	arg0 := fmt.Sprintf("-10")
	arg1 := fmt.Sprintf("main.go")
	stdout := []byte("ok")

	cmd := exec.Command(app, arg0, arg1)
	stdout, err = cmd.Output()

	if err != nil {
    	println(err.Error())
    	stdout = []byte(err.Error())
    	return
	}
	print(string(stdout))
	//check if above output is valid main.go file
	FL_VALID_DOM := false
	FL_VALID_DATE := false
    	scanner := bufio.NewScanner(bytes.NewReader(stdout))
    	for scanner.Scan() {
			sLineText := scanner.Text()		
			i := strings.Index(sLineText, "//GAE_APP_DOM_ID#")
			if i != -1 {
				FL_VALID_DOM = true
			}
			i = strings.Index(sLineText, "//LAST_UPGRADE#")
			if i != -1 {
				FL_VALID_DATE = true
			}
		}

	if FL_VALID_DOM != true && FL_VALID_DATE != true {
		print(string("ERROR: invalid main.go"))
		return
	} else {
		print(string("OK: valid main.go"))
	}
	
	fmt.Printf("\n++ Copying main.go to main.go.dev\n")	
	//cp main.go
	app = "cp"
	arg1 = fmt.Sprintf("main.go")
	arg2 := fmt.Sprintf("main.go.dev")
	stdout = []byte("ok")

	cmd = exec.Command(app, arg1, arg2)
	stdout, err = cmd.Output()

	if err != nil {
    	println(err.Error())
    	stdout = []byte(err.Error())
    	return
	}
	print(string(stdout))
	
	return err
}

//deploy ulapph cloud desktop
func deployUlapphCloudDesktop(project, account, yaml string) (err error) {
	fmt.Printf("\n+ ULAPPH Cloud Desktop Installation\n")
	fmt.Printf("\n++ Performing validations...\n")
	//validate project
	//valudate account
    	file, err := os.Open("main2.go")
    	if err != nil {
        log.Fatal(err)
		fmt.Printf("\nERROR: Missing main.go")
		return
    	}
    	defer file.Close()

	FL_VALID_PROJECT := false
	FL_VALID_ACCOUNT := false
    	scanner := bufio.NewScanner(file)
    	for scanner.Scan() {
		sLineText := scanner.Text()		
		i := strings.Index(sLineText, project)
		if i != -1 {
			FL_VALID_PROJECT = true
		}
		i = strings.Index(sLineText, account)
		if i != -1 {
			FL_VALID_ACCOUNT = true
		}
		if FL_VALID_PROJECT == true && FL_VALID_ACCOUNT == true {
			break
		}
	}
	if FL_VALID_PROJECT != true {
		fmt.Printf("\nERROR: Invalid project ID")
		return
	}
	if FL_VALID_ACCOUNT != true {
		fmt.Printf("\nERROR: Invalid account")
		return
	}

	fmt.Printf("\n++ Removing main.go backup...\n")
	//rm main.go.*
	app := "rm"
    	arg1 := fmt.Sprintf("main.go.backup")
	
    	cmd := exec.Command(app, arg1)
    	stdout, err := cmd.Output()

    	if err != nil {
        	println(err.Error())
        	stdout = []byte(err.Error())
        	return
    	}
	print(string(stdout))	

	fmt.Printf("\n++ Moving main2.go to main.go\n")	
	//mv main2.go
	app = "mv"
    	arg1 = fmt.Sprintf("main2.go")
    	arg2 := fmt.Sprintf("main.go")
	
    	cmd = exec.Command(app, arg1, arg2)
    	stdout, err = cmd.Output()

    	if err != nil {
        	println(err.Error())
        	stdout = []byte(err.Error())
        	return
    	}
	print(string(stdout))	
	
	fmt.Printf("\n++ Deploying to Google Appengine...\n")
	//gcloud --project=deathlake-fly --account=demo.ulapph@gmail.com --verbosity=info --quiet app deploy app.yaml
	app = "gcloud"
    	arg1 = fmt.Sprintf("--project=%v", project)
    	arg2 = fmt.Sprintf("--account=%v", account)
    	arg3 := fmt.Sprintf("--verbosity=info")
    	arg4 := fmt.Sprintf("--quiet")
    	arg5 := fmt.Sprintf("app")
    	arg6 := fmt.Sprintf("deploy")
	arg7 := fmt.Sprintf("%v", yaml)
	
    	cmd = exec.Command(app, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    	stdout, err = cmd.Output()
	
    	if err != nil {
        	println(err.Error())
        	stdout = []byte(err.Error())
        	return
    	}
	print(string(stdout))	
	fmt.Printf("\n++ Deployment completed...\n")
	fmt.Printf("\nhttps://%v.appspot.com/uwm", project)
	fmt.Printf("\nhttps://%v.appspot.com/contents?q=home", project)
	fmt.Printf("\nhttps://%v.appspot.com/admin-setup", project)
	fmt.Println("\n")
	
	return err
}

//redeploy ulapph cloud desktop
func reDeployUlapphCloudDesktop(project, account, yaml string) (err error) {
	fmt.Printf("\n+ ULAPPH Cloud Desktop Installation\n")
	//fmt.Printf("\n++ Performing validations...\n")
	fmt.Printf("\n++ Re-Deploying to Google Appengine...\n")
	//gcloud --project=deathlake-fly --account=demo.ulapph@gmail.com --verbosity=info --quiet app deploy app.yaml
	app := "gcloud"
    	arg1 := fmt.Sprintf("--project=%v", project)
    	arg2 := fmt.Sprintf("--account=%v", account)
    	arg3 := fmt.Sprintf("--verbosity=info")
    	arg4 := fmt.Sprintf("--quiet")
    	arg5 := fmt.Sprintf("app")
    	arg6 := fmt.Sprintf("deploy")
	arg7 := fmt.Sprintf("%v", yaml)
	
    	cmd := exec.Command(app, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    	stdout, err := cmd.Output()
	
    	if err != nil {
        	println(err.Error())
        	stdout = []byte(err.Error())
        	return
    	}
	print(string(stdout))	
	fmt.Printf("\n++ Re-Deployment completed...\n")
	fmt.Printf("\nhttps://%v.appspot.com/uwm", project)
	fmt.Printf("\nhttps://%v.appspot.com/contents?q=home", project)
	fmt.Printf("\nhttps://%v.appspot.com/admin-setup", project)
	fmt.Println("\n")
	
	return err
}

//configure ulapph cloud desktop
func configureUlapphCloudDesktop(CFG_FILE string) (err error) {
	TOTAL_LINES := 0
	//PB_STARTED := false
	bar := pb.StartNew(5)
	//------------------------------
	//Load the configuration file
	configor.Load(&Config, CFG_FILE)
	//fmt.Printf("config: %#v", Config)
	fmt.Printf("app_id: %#v", Config.Project[0].Appid)
	
	//-----------------------------
	//Backup main.go
	bar.Increment()
	time.Sleep(time.Microsecond)
	//fmt.Printf("\n+ Backup main.go to main2.go...  ")
    	app := "cp"
	//currenttime := time.Now().Local()
	//TSTMP := currenttime.Format("2006-01-02-15-04-05")	
    	arg1 := Config.Installer[0].Dir+"/main.go"
    	//arg2 := Config.Installer[0].Dir+"/main.go"+"."+TSTMP
	arg2 := Config.Installer[0].Dir+"/main.go.backup"
	
    	cmd := exec.Command(app, arg1, arg2)
    	_, err = cmd.Output()

    	if err != nil {
        println(err.Error())
        return
    	} else {
		//stdout = []byte("ok\n")
	}
	//print(string(stdout))
	
	//-----------------------------
	//Configuring installation
	bar.Increment()
	time.Sleep(time.Microsecond)
	//fmt.Printf("\n+ Customizing main.go...  ")
    	file, err := os.Open(Config.Installer[0].Dir+"/main.go")
    	if err != nil {
        log.Fatal(err)
		//stdout = []byte(fmt.Sprintf("%v",err))
    	} else {
		//stdout = []byte("ok\n")
	}
	//print(string(stdout))
    	defer file.Close()

    	scanner := bufio.NewScanner(file)
	lineCtr := 0
	
	//write buffer
	var buf bytes.Buffer
	//FL_VALID_FILE := false
	FL_START_CUST_CONFIGS := false
	FL_END_CUST_CONFIGS := false
	ERR1CTR := 0
	buf.WriteString("//ULAPPH Cloud Desktop\n")
	buf.WriteString(fmt.Sprintf("//Auto-generated codes for %v\n", Config.Project[0].Appid))
	
    	for scanner.Scan() {
    		FL_WRITTEN_OK := false
    		
			lineCtr++
			sLineText := scanner.Text()
			//if PB_STARTED == true {
			if lineCtr > 10000 {
				bar.Increment()
				time.Sleep(time.Microsecond)				
			} else {
				bar.Increment()
				time.Sleep(time.Millisecond)
			}
			
			if lineCtr == 1 {
				bar.Increment()
				time.Sleep(time.Millisecond)
				//-----------------------------
				//Configuring installation
				//fmt.Printf("\n++ Replacing //GAE_APP_DOM_ID#...  ")
				i := strings.Index(sLineText, "GAE_APP_DOM_ID#")
				if i != -1 {
					//print(string("ok\n"))
					//split in #
					SPL := strings.Split(sLineText, "#")
					configValue := getFromConfig("APP_URL")
					//buf.WriteString(fmt.Sprintf("%v#%v\n", SPL[0], configValue))
					buf.WriteString(fmt.Sprintf("//%v\n", configValue))
				        //fmt.Printf("\nNEWLINE001: %v", fmt.Sprintf("//%v\n", configValue))
					FL_WRITTEN_OK = true
					//FL_VALID_FILE = true
					OLD_DOMAIN_NAME = SPL[1]
					NEW_DOMAIN_NAME = configValue
					//fmt.Printf("\n++ OLD_DOMAIN_NAME: %v", OLD_DOMAIN_NAME)
					//fmt.Printf("\n++ NEW_DOMAIN_NAME: %v", NEW_DOMAIN_NAME)
				}
			}
			if lineCtr == 2 {
				bar.Increment()
				time.Sleep(time.Millisecond)
				//-----------------------------
				//Configuring installation
				//fmt.Printf("\n++ Replacing //LAST_UPGRADE#...  ")
				i := strings.Index(sLineText, "LAST_UPGRADE#")
				if i != -1 {
					//print(string("ok\n"))
					//split in #
					//SPL := strings.Split(sLineText, "#")
					currenttime := time.Now().Local()
					TSTMP := currenttime.Format("02/01/2006 15:04:05")
					//buf.WriteString(fmt.Sprintf("%v#%v\n", SPL[0], TSTMP))
					buf.WriteString(fmt.Sprintf("//%v\n", TSTMP))
					//fmt.Printf("\nNEWLINE002: %v", fmt.Sprintf("//%v\n", TSTMP))
					FL_WRITTEN_OK = true
					//FL_VALID_FILE = true
				}
			}
	
			if lineCtr == 3 {
				bar.Increment()
				time.Sleep(time.Millisecond)
				//-----------------------------
				//Configuring installation
				//fmt.Printf("\n++ Getting //TOTAL_LINES#...  ")
				i := strings.Index(sLineText, "TOTAL_LINES#")
				if i != -1 {
					//print(string("ok\n"))
					//split in #
					SPL := strings.Split(sLineText, "#")
					tot, _ := strconv.Atoi(SPL[1])
					TOTAL_LINES = tot
					bar = pb.StartNew(TOTAL_LINES)
					//PB_STARTED = true
					FL_WRITTEN_OK = true
					//FL_VALID_FILE = true
				}
			}
			
			//replace all old domains w/ new domains
			i := strings.Index(sLineText, OLD_DOMAIN_NAME)
			if i != -1 && OLD_DOMAIN_NAME != "" && NEW_DOMAIN_NAME != "" {
				sLineText = strings.Replace(sLineText, OLD_DOMAIN_NAME, NEW_DOMAIN_NAME, -1)
				//fmt.Printf("\nREPLACED: Old domain replaced with new!")
				//fmt.Printf("\nNEWTEXT: %v", fmt.Sprintf("%v\n", sLineText))
			}
			
			//--------------------------------
			i = strings.Index(sLineText, "// !!!CONFIG-STARTS-HERE!!!")
			if i != -1 {
				//-----------------------------
				//Configuring installation
				//fmt.Printf("\n++ Setting flag FL_START_CUST_CONFIGS...  ")
				//fmt.Printf("\nFL_START_CUST_CONFIGS--------------------")
				FL_START_CUST_CONFIGS = true
			}
			
			i = strings.Index(sLineText, "// !!!CONFIG-ENDS-HERE!!!")
			if i != -1 {
				//-----------------------------
				//Configuring installation
				//fmt.Printf("\n++ Setting flag FL_START_CUST_CONFIGS...  ")
				//fmt.Printf("\nFL_START_CUST_CONFIGS--------------------")
				FL_END_CUST_CONFIGS = true
			}
			
			tLineStr := fmt.Sprintf("%v", sLineText)
			tLineStr2 := strings.TrimSpace(tLineStr)
			if len(tLineStr2) > 2 && string(tLineStr2[0]) != "/" && string(tLineStr2[1]) != "/" {
	
				// internally, it advances token based on sperator
				//fmt.Println(fmt.Sprintf("\nLINE: %v", lineCtr))  // token in unicode-char
				//fmt.Println(sLineText)  // token in unicode-char
		        	//fmt.Println(scanner.Bytes()) // token in bytes
				//FL_WRITTEN_OK := false
		
				i := strings.Index(sLineText, "<title>")
				if i != -1 {
					//-----------------------------
					//Configuring installation
					//fmt.Printf("\n++ Processing <title>...  ")
					//replace
					//lineText := strings.Replace(sLineText, OLD_DOMAIN_NAME, NEW_DOMAIN_NAME, -1)
					buf.WriteString(fmt.Sprintf("%v\n", sLineText))
					//fmt.Printf("\nNEWLINE003: %v", fmt.Sprintf("%v\n", sLineText))
					//print(string("ok\n"))
					FL_WRITTEN_OK = true
				}
				
				i = strings.Index(sLineText, "ULAPPH_META_DESCRIPTION_CONTENT")
				if i != -1 {
					//-----------------------------
					//Configuring installation
					//fmt.Printf("\n++ Processing ULAPPH_META_DESCRIPTION_CONTENT...  ")
					//replace
					configValue := getFromConfig("INFO_ABOUT_US")		
					lineText := strings.Replace(sLineText, "ULAPPH_META_DESCRIPTION_CONTENT", configValue, -1)
					buf.WriteString(fmt.Sprintf("%v\n", lineText))
					//fmt.Printf("\nNEWLINE004: %v", fmt.Sprintf("%v\n", lineText))
					//print(string("ok\n"))
					FL_WRITTEN_OK = true
				}
				
				i = strings.Index(sLineText, "ULAPPH_META_KEYWORDS_CONTENT")
				if i != -1 {
					//-----------------------------
					//Configuring installation
					//fmt.Printf("\n++ Processing ULAPPH_META_KEYWORDS_CONTENT...  ")
					//replace
					configValue := getFromConfig("<meta keywords>")		
					lineText := strings.Replace(sLineText, "ULAPPH_META_KEYWORDS_CONTENT", configValue, -1)
					buf.WriteString(fmt.Sprintf("%v\n", lineText))
					//fmt.Printf("\nNEWLINE005: %v", fmt.Sprintf("%v\n", lineText))
					//print(string("ok\n"))
					FL_WRITTEN_OK = true
				}
				
				i = strings.Index(sLineText, "var isExceptionAccount")
				if i != -1 {
					//-----------------------------
					//Configuring installation
					//fmt.Printf("\n++ Processing var isExceptionAccount...  ")
					//replace
					configValue := getFromConfig("var isExceptionAccount")		
					i = strings.Index(configValue, "map[")
					if i != -1 {
						buf.WriteString(fmt.Sprintf("    %v = %v\n", "var isExceptionAccount", configValue))
						//fmt.Printf("\nNEWLINE006: %v", fmt.Sprintf("    %v = %v\n", "var isExceptionAccount", configValue))
					} else {
						buf.WriteString(fmt.Sprintf("var isExceptionAccount = `%v`\n", configValue))
						//fmt.Printf("\nNEWLINE007: %v", fmt.Sprintf("var isExceptionAccount = `%v`\n", configValue))
					}
	
					//print(string("ok\n"))
					FL_WRITTEN_OK = true
				}
		
				i = strings.Index(sLineText, "var isCountryAllowed")
				if i != -1 {
					//-----------------------------
					//Configuring installation
					//fmt.Printf("\n++ Processing var isCountryAllowed...  ")
					//replace
					configValue := getFromConfig("var isCountryAllowed")	
					i = strings.Index(configValue, "map[")
					if i != -1 {
						buf.WriteString(fmt.Sprintf("    %v = %v\n", "var isCountryAllowed", configValue))
						//fmt.Printf("\nNEWLINE008: %v", fmt.Sprintf("    %v = %v\n", "var isCountryAllowed", configValue))
					} else {
						buf.WriteString(fmt.Sprintf("var isCountryAllowed = `%v`\n", configValue))
						//fmt.Printf("\nNEWLINE009: %v", fmt.Sprintf("var isCountryAllowed = `%v`\n", configValue))
					}
	
					//print(string("ok\n"))
					FL_WRITTEN_OK = true
				}
				
				i = strings.Index(sLineText, "var isCountryNotAllowed")
				if i != -1 {
					//-----------------------------
					//Configuring installation
					//fmt.Printf("\n++ Processing var isCountryNotAllowed...  ")
					//replace
					configValue := getFromConfig("var isCountryNotAllowed")
					i = strings.Index(configValue, "map[")
					if i != -1 {
						buf.WriteString(fmt.Sprintf("    %v = %v\n", "var isCountryNotAllowed", configValue))
						//fmt.Printf("\nNEWLINE010: %v", fmt.Sprintf("    %v = %v\n", "var isCountryNotAllowed", configValue))
					} else {
						buf.WriteString(fmt.Sprintf("var isCountryNotAllowed = `%v`\n", configValue))
						//fmt.Printf("\nNEWLINE011: %v", fmt.Sprintf("var isCountryNotAllowed = `%v`\n", configValue))
					}
	
					//print(string("ok\n"))
					FL_WRITTEN_OK = true
				}
				
				i = strings.Index(sLineText, "var isInBoundAppidAllowed")
				if i != -1 {
					//-----------------------------
					//Configuring installation
					//fmt.Printf("\n++ Processing var isInBoundAppidAllowed...  ")
					//replace
					configValue := getFromConfig("var isInBoundAppidAllowed")
					i = strings.Index(configValue, "map[")
					if i != -1 {
						buf.WriteString(fmt.Sprintf("    %v = %v\n", "var isInBoundAppidAllowed", configValue))
						//fmt.Printf("\nNEWLINE012: %v", fmt.Sprintf("    %v = %v\n", "var isInBoundAppidAllowed", configValue))
					} else {
						buf.WriteString(fmt.Sprintf("var isInBoundAppidAllowed = `%v`\n", configValue))
						//fmt.Printf("\nNEWLINE013: %v", fmt.Sprintf("var isInBoundAppidAllowed = `%v`\n", configValue))
					}
	
					//print(string("ok\n"))
					FL_WRITTEN_OK = true
				}
				
				i = strings.Index(sLineText, "var isSearchEngineAllowed")
				if i != -1 {
					//-----------------------------
					//Configuring installation
					//fmt.Printf("\n++ Processing var isSearchEngineAllowed...  ")
					//replace
					configValue := getFromConfig("var isSearchEngineAllowed")
					i = strings.Index(configValue, "map[")
					if i != -1 {
						buf.WriteString(fmt.Sprintf("    %v = %v\n", "var isSearchEngineAllowed", configValue))
						//fmt.Printf("\nNEWLINE014: %v", fmt.Sprintf("    %v = %v\n", "var isSearchEngineAllowed", configValue))
					} else {
						buf.WriteString(fmt.Sprintf("var isSearchEngineAllowed = `%v`\n", configValue))
						//fmt.Printf("\nNEWLINE015: %v", fmt.Sprintf("var isSearchEngineAllowed = `%v`\n", configValue))
					}
	
					//print(string("ok\n"))
					FL_WRITTEN_OK = true
				}
				
				i = strings.Index(sLineText, "var freeAccess")
				if i != -1 {
					//-----------------------------
					//Configuring installation
					//fmt.Printf("\n++ Processing var freeAccess...  ")
					//replace
					configValue := getFromConfig("var freeAccess")
					i = strings.Index(configValue, "map[")
					if i != -1 {
						buf.WriteString(fmt.Sprintf("    %v = %v\n", "var freeAccess", configValue))
						//fmt.Printf("\nNEWLINE016: %v", fmt.Sprintf("    %v = %v\n", "var freeAccess", configValue))
					} else {
						buf.WriteString(fmt.Sprintf("var freeAccess = `%v`\n", configValue))
						//fmt.Printf("\nNEWLINE017: %v", fmt.Sprintf("var freeAccess = `%v`\n", configValue))
					}
	
					//print(string("ok\n"))
					FL_WRITTEN_OK = true
				}
				
				i = strings.Index(sLineText, "var bronzeAccess")
				if i != -1 {
					//-----------------------------
					//Configuring installation
					//fmt.Printf("\n++ Processing var bronzeAccesss...  ")
					//replace
					configValue := getFromConfig("var bronzeAccess")
					i = strings.Index(configValue, "map[")
					if i != -1 {
						buf.WriteString(fmt.Sprintf("    %v = %v\n", "var bronzeAccess", configValue))
						//fmt.Printf("\nNEWLINE018: %v", fmt.Sprintf("    %v = %v\n", "var bronzeAccess", configValue))
					} else {
						buf.WriteString(fmt.Sprintf("var bronzeAccess = `%v`\n", configValue))
						//fmt.Printf("\nNEWLINE019: %v", fmt.Sprintf("var bronzeAccess = `%v`\n", configValue))
					}
	
					//print(string("ok\n"))
					FL_WRITTEN_OK = true
				}
				
				i = strings.Index(sLineText, "var silverAccess")
				if i != -1 {
					//-----------------------------
					//Configuring installation
					//fmt.Printf("\n++ Processing var silverAccess...  ")
					//replace
					configValue := getFromConfig("var silverAccess")	
					i = strings.Index(configValue, "map[")
					if i != -1 {
						buf.WriteString(fmt.Sprintf("    %v = %v\n", "var silverAccess", configValue))
						//fmt.Printf("\nNEWLINE020: %v", fmt.Sprintf("    %v = %v\n", "var silverAccess", configValue))
					} else {
						buf.WriteString(fmt.Sprintf("var silverAccess = `%v`\n", configValue))
						//fmt.Printf("\nNEWLINE021: %v", fmt.Sprintf("var silverAccess = `%v`\n", configValue))
					}
	
					//print(string("ok\n"))
					FL_WRITTEN_OK = true
				}
				
				i = strings.Index(sLineText, "var goldAccess")
				if i != -1 {
					//-----------------------------
					//Configuring installation
					//fmt.Printf("\n++ Processing var goldAccess...  ")
					//replace
					configValue := getFromConfig("var goldAccess")
					i = strings.Index(configValue, "map[")
					if i != -1 {
						buf.WriteString(fmt.Sprintf("    %v = %v\n", "var goldAccess", configValue))
						//fmt.Printf("\nNEWLINE022: %v", fmt.Sprintf("    %v = %v\n", "var goldAccess", configValue))
					} else {
						buf.WriteString(fmt.Sprintf("var goldAccess = `%v`\n", configValue))
						//fmt.Printf("\nNEWLINE023: %v", fmt.Sprintf("var goldAccess = `%v`\n", configValue))
					}
	
					//print(string("ok\n"))
					FL_WRITTEN_OK = true
				}
				
				i = strings.Index(sLineText, "UCD_BUILD_STR = ")
				if i != -1 {
					//-----------------------------
					//Configuring installation
					//fmt.Printf("\n++ Processing UCD_BUILD_STR = ...  ")
					//replace
					currenttime := time.Now().Local()
					TSTMP := currenttime.Format("2006-01-02-15-04-05")	
					buf.WriteString(fmt.Sprintf("    UCD_BUILD_STR = `BUILD_%v`", TSTMP))
					//fmt.Printf("\nNEWLINE024: %v", fmt.Sprintf("    UCD_BUILD_STR = `BUILD_%v`", TSTMP))
					buf.WriteString("\n")
					//print(string("ok\n"))
					FL_WRITTEN_OK = true
				}
				
				i = strings.Index(sLineText, "ULAPPH_ADMIN_EMAIL")
				if i != -1 {
					//-----------------------------
					//Configuring installation
					//fmt.Printf("\n++ Processing ULAPPH_ADMIN_EMAIL...  ")
					//replace
					configValue := getFromConfig("ADMIN_ACCOUNT")		
					lineText := strings.Replace(sLineText, "ULAPPH_ADMIN_EMAIL", configValue, -1)
					buf.WriteString(fmt.Sprintf("%v\n", lineText))
					//fmt.Printf("\nNEWLINE025: %v", fmt.Sprintf("%v\n", lineText))
					//print(string("ok\n"))
					FL_WRITTEN_OK = true
				}
				
				i = strings.Index(sLineText, "] Read Datastore...")
				if i != -1 {
					//-----------------------------
					//Configuring installation
					//fmt.Printf("\n++ Tagging Read Datastore...  ")
					//replace
					//c.Infof("[R001] Read Datastore...")
					SPL := strings.Split(sLineText, "[R")
					//SPL2 := strings.Split(SPL[1], "]")
					//thisNum := SPL2[0]
					ERR1CTR = ERR1CTR + 1
					ERRCODE := padNumberWithZero(4, ERR1CTR)
					buf.WriteString(fmt.Sprintf("%v[R%v] Read Datastore...\n", SPL[0], ERRCODE))
					//fmt.Printf("\nNEWLINE026: %v", fmt.Sprintf("%v[R%v] Read Datastore...\n", SPL[0], ERRCODE))
					//print(string("ok\n"))
					FL_WRITTEN_OK = true
				}
				
				i = strings.Index(sLineText, "] Write Datastore...")
				if i != -1 {
					//-----------------------------
					//Configuring installation
					//fmt.Printf("\n++ Tagging Write Datastore...  ")
					//replace
					//c.Infof("[R001] Write Datastore...")
					SPL := strings.Split(sLineText, "[W")
					//SPL2 := strings.Split(SPL[1], "]")
					//thisNum := SPL2[0]
					ERR1CTR = ERR1CTR + 1
					ERRCODE := padNumberWithZero(4, ERR1CTR)
					buf.WriteString(fmt.Sprintf("%v[W%v] Write Datastore...\n", SPL[0], ERRCODE))
					//fmt.Printf("\nNEWLINE027: %v", fmt.Sprintf("%v[W%v] Write Datastore...\n", SPL[0], ERRCODE))
					//print(string("ok\n"))
					FL_WRITTEN_OK = true
				}
				
				//Process config values
				if FL_START_CUST_CONFIGS == true && FL_END_CUST_CONFIGS == false {
					
					//fmt.Printf("\nPROCESSING CONFIGS...")
					//fmt.Printf("\nLINE: %v", sLineText)
					//loop from configs
					for _, cfg := range Config.Configs {
						
						tStr := fmt.Sprintf("%v =", strings.TrimSpace(cfg.Item))
						i = strings.Index(sLineText, tStr)
						//SPL := strings.Split(sLineText, "=")
						//if strings.TrimSpace(SPL[0]) == strings.TrimSpace(cfg.Item) {
						if i != -1 {
							//-----------------------------
							//Configuring installation
							//fmt.Printf("\n++ Processing cfg.Item...  ")
							//fmt.Printf("\n+++ cfg.Item: %v", cfg.Item)
							//fmt.Printf("\n+++ cfg.Format: %v", cfg.Format)
							//fmt.Printf("\n+++ cfg.Status: %v", cfg.Status)
							//fmt.Printf("\n+++ cfg.Value: %v", cfg.Value)
							//fmt.Printf("\nCFG[%v]: %v", key, fmt.Sprintf("%v", strings.TrimSpace(SPL[0])))
							//fmt.Printf("\nCONFIGS[%v]: %v", key, fmt.Sprintf("%v", strings.TrimSpace(cfg.Item)))
		
							switch cfg.Format {
								case "Flag":
									buf.WriteString(fmt.Sprintf("    %v = %v\n", cfg.Item, cfg.Value))
									//fmt.Printf("\nNEWLINE028: %v", fmt.Sprintf("    %v = %v\n", cfg.Item, cfg.Value))
								case "Number":
									num, err := strconv.Atoi(cfg.Value)
									if err != nil {
										fmt.Printf("\nERROR: %v", cfg)
										panic(err)
										break
									}
									buf.WriteString(fmt.Sprintf("    %v = %v\n", cfg.Item, num))
									//fmt.Printf("\nNEWLINE029: %v", fmt.Sprintf("    %v = %v\n", cfg.Item, num))
								case "Text":
									if cfg.Item == "FIREBASE_SERVER_JSON" {
										//read file and append
										b, err := ioutil.ReadFile(cfg.Value) // just pass the file name
										if err != nil {
											//fmt.Print(err)
											fmt.Printf("\nERROR: %v", err)
											panic(err)
											break
										}
										buf.WriteString(fmt.Sprintf("    %v = `%v`\n", cfg.Item, string(b)))
										//fmt.Printf("\nNEWLINE030: %v", fmt.Sprintf("    %v = `%v`\n", cfg.Item, string(b)))
									} else {
										//for maps
										i = strings.Index(cfg.Value, "map[")
										if i != -1 {
											buf.WriteString(fmt.Sprintf("    %v = %v\n", cfg.Item, cfg.Value))
											//fmt.Printf("\nNEWLINE031: %v", fmt.Sprintf("    %v = %v\n", cfg.Item, cfg.Value))
										} else {	 
											buf.WriteString(fmt.Sprintf("    %v = `%v`\n", cfg.Item, cfg.Value))
											//fmt.Printf("\nNEWLINE032: %v", fmt.Sprintf("    %v = `%v`\n", cfg.Item, cfg.Value))
										}
	
									}
							}
							//print(string("ok\n"))
							FL_WRITTEN_OK = true
							break
						}	
					}
		
				
					
				}
				
				//--------------------------
				//default
				//for those not edited lines
				if FL_WRITTEN_OK == false {
					buf.WriteString(fmt.Sprintf("%v\n", sLineText))
					//fmt.Printf("\nNEWLINE033(NOT-EDITED): %v", sLineText)
				}
	
			} else {
				if len(tLineStr2) > 2 && string(tLineStr2[0]) == "/" && string(tLineStr2[1]) == "/" {
					//skip comments
				} else {
					buf.WriteString(fmt.Sprintf("%v\n", sLineText))	
					//fmt.Printf("\nNEWLINE034(AS-IS): %v", sLineText)
				}
			}
    	}
	//if FL_VALID_FILE == false {
	//	fmt.Printf("\n+ ERROR: Invalid file!!!")
	//}
	//-----------------------------
	//Writing modified file
	//fmt.Printf("\n+ Writing modified file...  ")	
    	err = ioutil.WriteFile(Config.Installer[0].Dir+"/main2.go", buf.Bytes(), 0644)
    	if err != nil {
        log.Fatal(err)
		//stdout = []byte(fmt.Sprintf("%v",err))
    	} else {
		//stdout = []byte("ok\n")
		}
	//print(string(stdout))
	
	return err
}

//get value of config item
func getFromConfig(key string) (retItem string) {
	//loop from configs
	for _, cfg := range Config.Configs {
		if cfg.Item == key {
			retItem =  cfg.Value
			//fmt.Printf("\n+++ cfg.Item: %v", cfg.Item)
			//fmt.Printf("\n+++ cfg.Format: %v", cfg.Format)
			//fmt.Printf("\n+++ cfg.Status: %v", cfg.Status)
			//fmt.Printf("\n+++ cfg.Value: %v", cfg.Value)
			break
		}
	}
	return retItem
}

//pad with leading zeroes
func padNumberWithZero(digits int, value int) string {
	switch digits {
		case 4:
			return fmt.Sprintf("%04d", value)
	}
    return ""
}
