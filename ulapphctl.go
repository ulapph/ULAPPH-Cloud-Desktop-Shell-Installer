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
	
	app := cli.NewApp()

	app.Flags = []cli.Flag {
		cli.StringFlag{
		  Name: "config, c",
		  Usage: "Configuration file for the ulapph cloud destkop",
		  Destination: &config,
		},
	}

	app.Commands = []cli.Command{
	{
	  Name:    "install",
	  Aliases: []string{"i"},
	  Usage:   "install ulapph cloud desktop",
	  Action:  func(c *cli.Context) error {
		if config == "" {
			fmt.Printf("ERROR: Missing configuration file!")
			fmt.Printf("\nTry: ulapphctl install --config your-ulapph-cloud-desktop.yaml")
			
			return nil
		}
		installUlapphCloudDesktop(config)
		return nil
	  },
	},
	{
	  Name:    "upgrade",
	  Aliases: []string{"u"},
	  Usage:   "upgrade an existing ulapph cloud desktop",
	  Action:  func(c *cli.Context) error {
		return nil
	  },
	},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
  
}

//install ulapph cloud desktop
func installUlapphCloudDesktop(CFG_FILE string) (err error) {
	//------------------------------
	//Load the configuration file
	configor.Load(&Config, CFG_FILE)
	//fmt.Printf("config: %#v", Config)
	fmt.Printf("app_id: %#v", Config.Project[0].Appid)
	
	//-----------------------------
	//Backup main.go
	fmt.Printf("\n+ Backup main.go to main2.go...  ")
    	app := "cp"
	currenttime := time.Now().Local()
	TSTMP := currenttime.Format("2006-01-02-15-04-05")	
    	arg1 := Config.Installer[0].Dir+"/main.go"
    	arg2 := Config.Installer[0].Dir+"/main.go"+"."+TSTMP

    	cmd := exec.Command(app, arg1, arg2)
    	stdout, err := cmd.Output()

    	if err != nil {
        println(err.Error())
        return
    	} else {
		stdout = []byte("ok")
	}
	print(string(stdout))
	
	//-----------------------------
	//Configuring installation
	fmt.Printf("\n+ Customizing main.go...  ")
    	file, err := os.Open(Config.Installer[0].Dir+"/main.go")
    	if err != nil {
        log.Fatal(err)
		stdout = []byte(fmt.Sprintf("%v",err))
    	} else {
		stdout = []byte("ok")
	}
	print(string(stdout))
    	defer file.Close()

    	scanner := bufio.NewScanner(file)
	lineCtr := 0
	
	//write buffer
	var buf bytes.Buffer
	//FL_VALID_FILE := false
	FL_START_CUST_CONFIGS := false
	FL_END_CUST_CONFIGS := false
	ERR1CTR := 0
	
    	for scanner.Scan() {
		lineCtr++
		//check if line is a comment
		tLineStr := fmt.Sprintf("%v", scanner.Text())
		tLineStr2 := strings.TrimSpace(tLineStr)
		if len(tLineStr2) > 2 && string(tLineStr2[0]) != "/" && string(tLineStr2[1]) != "/" {

		// internally, it advances token based on sperator
		fmt.Println(fmt.Sprintf("\nLINE: %v", lineCtr))  // token in unicode-char
		fmt.Println(scanner.Text())  // token in unicode-char
        //fmt.Println(scanner.Bytes()) // token in bytes
		FL_WRITTEN_OK := false
		
		if lineCtr == 1 {
			//-----------------------------
			//Configuring installation
			fmt.Printf("\n++ Replacing //GAE_APP_DOM_ID#...  ")
			i := strings.Index(scanner.Text(), "GAE_APP_DOM_ID#")
			if i != -1 {
				print(string("ok"))
				//split in #
				SPL := strings.Split(scanner.Text(), "#")
				configValue := getFromConfig("APP_URL")
				buf.WriteString(fmt.Sprintf("%v#%v\n", SPL[0], configValue))
				FL_WRITTEN_OK = true
				//FL_VALID_FILE = true
				OLD_DOMAIN_NAME = SPL[1]
				NEW_DOMAIN_NAME = configValue
			}
		}
		if lineCtr == 2 {
			//-----------------------------
			//Configuring installation
			fmt.Printf("\n++ Replacing //LAST_UPGRADE#...  ")
			i := strings.Index(scanner.Text(), "LAST_UPGRADE#")
			if i != -1 {
				print(string("ok"))
				//split in #
				SPL := strings.Split(scanner.Text(), "#")
				currenttime := time.Now().Local()
				TSTMP := currenttime.Format("02/01/2006 15:04:05")
				buf.WriteString(fmt.Sprintf("%v#%v\n", SPL[0], TSTMP))
				FL_WRITTEN_OK = true
				//FL_VALID_FILE = true
			}
		}

		i := strings.Index(scanner.Text(), "<title>")
		if i != -1 {
			//-----------------------------
			//Configuring installation
			fmt.Printf("\n++ Processing <title>...  ")
			//replace
			lineText := strings.Replace(scanner.Text(), OLD_DOMAIN_NAME, NEW_DOMAIN_NAME, -1)
			buf.WriteString(fmt.Sprintf("%v\n", lineText))
			print(string("ok"))
			FL_WRITTEN_OK = true
		}
		
		i = strings.Index(scanner.Text(), "ULAPPH_META_DESCRIPTION_CONTENT")
		if i != -1 {
			//-----------------------------
			//Configuring installation
			fmt.Printf("\n++ Processing ULAPPH_META_DESCRIPTION_CONTENT...  ")
			//replace
			configValue := getFromConfig("INFO_ABOUT_US")		
			lineText := strings.Replace(scanner.Text(), "ULAPPH_META_DESCRIPTION_CONTENT", configValue, -1)
			buf.WriteString(fmt.Sprintf("%v\n", lineText))
			print(string("ok"))
			FL_WRITTEN_OK = true
		}
		
		i = strings.Index(scanner.Text(), "ULAPPH_META_KEYWORDS_CONTENT")
		if i != -1 {
			//-----------------------------
			//Configuring installation
			fmt.Printf("\n++ Processing ULAPPH_META_KEYWORDS_CONTENT...  ")
			//replace
			configValue := getFromConfig("<meta keywords>")		
			lineText := strings.Replace(scanner.Text(), "ULAPPH_META_KEYWORDS_CONTENT", configValue, -1)
			buf.WriteString(fmt.Sprintf("%v\n", lineText))
			print(string("ok"))
			FL_WRITTEN_OK = true
		}
		
		i = strings.Index(scanner.Text(), "var isExceptionAccount")
		if i != -1 {
			//-----------------------------
			//Configuring installation
			fmt.Printf("\n++ Processing var isExceptionAccount...  ")
			//replace
			configValue := getFromConfig("var isExceptionAccount")		
			buf.WriteString(fmt.Sprintf("var isExceptionAccount = `%v`\n", configValue))
			print(string("ok"))
			FL_WRITTEN_OK = true
		}

		i = strings.Index(scanner.Text(), "var isCountryAllowed")
		if i != -1 {
			//-----------------------------
			//Configuring installation
			fmt.Printf("\n++ Processing var isCountryAllowed...  ")
			//replace
			configValue := getFromConfig("var isCountryAllowed")		
			buf.WriteString(fmt.Sprintf("var isCountryAllowed = `%v`\n", configValue))
			print(string("ok"))
			FL_WRITTEN_OK = true
		}
		
		i = strings.Index(scanner.Text(), "var isCountryNotAllowed")
		if i != -1 {
			//-----------------------------
			//Configuring installation
			fmt.Printf("\n++ Processing var isCountryNotAllowed...  ")
			//replace
			configValue := getFromConfig("var isCountryNotAllowed")		
			buf.WriteString(fmt.Sprintf("var isCountryNotAllowed = `%v`\n", configValue))
			print(string("ok"))
			FL_WRITTEN_OK = true
		}
		
		i = strings.Index(scanner.Text(), "var isInBoundAppidAllowed")
		if i != -1 {
			//-----------------------------
			//Configuring installation
			fmt.Printf("\n++ Processing var isInBoundAppidAllowed...  ")
			//replace
			configValue := getFromConfig("var isInBoundAppidAllowed")		
			buf.WriteString(fmt.Sprintf("var isInBoundAppidAllowed = `%v`\n", configValue))
			print(string("ok"))
			FL_WRITTEN_OK = true
		}
		
		i = strings.Index(scanner.Text(), "var isSearchEngineAllowed")
		if i != -1 {
			//-----------------------------
			//Configuring installation
			fmt.Printf("\n++ Processing var isSearchEngineAllowed...  ")
			//replace
			configValue := getFromConfig("var isSearchEngineAllowed")		
			buf.WriteString(fmt.Sprintf("var isSearchEngineAllowed = `%v`\n", configValue))
			print(string("ok"))
			FL_WRITTEN_OK = true
		}
		
		i = strings.Index(scanner.Text(), "var freeAccess")
		if i != -1 {
			//-----------------------------
			//Configuring installation
			fmt.Printf("\n++ Processing var freeAccess...  ")
			//replace
			configValue := getFromConfig("var freeAccess")		
			buf.WriteString(fmt.Sprintf("var freeAccess = `%v`\n", configValue))
			print(string("ok"))
			FL_WRITTEN_OK = true
		}
		
		i = strings.Index(scanner.Text(), "var bronzeAccess")
		if i != -1 {
			//-----------------------------
			//Configuring installation
			fmt.Printf("\n++ Processing var bronzeAccesss...  ")
			//replace
			configValue := getFromConfig("var bronzeAccess")		
			buf.WriteString(fmt.Sprintf("var bronzeAccess = `%v`\n", configValue))
			print(string("ok"))
			FL_WRITTEN_OK = true
		}
		
		i = strings.Index(scanner.Text(), "var silverAccess")
		if i != -1 {
			//-----------------------------
			//Configuring installation
			fmt.Printf("\n++ Processing var silverAccess...  ")
			//replace
			configValue := getFromConfig("var silverAccess")		
			buf.WriteString(fmt.Sprintf("var silverAccess = `%v`\n", configValue))
			print(string("ok"))
			FL_WRITTEN_OK = true
		}
		
		i = strings.Index(scanner.Text(), "var goldAccess")
		if i != -1 {
			//-----------------------------
			//Configuring installation
			fmt.Printf("\n++ Processing var goldAccess...  ")
			//replace
			configValue := getFromConfig("var goldAccess")		
			buf.WriteString(fmt.Sprintf("var goldAccess = `%v`\n", configValue))
			print(string("ok"))
			FL_WRITTEN_OK = true
		}
		
		i = strings.Index(scanner.Text(), "UCD_BUILD_STR = ")
		if i != -1 {
			//-----------------------------
			//Configuring installation
			fmt.Printf("\n++ Processing UCD_BUILD_STR = ...  ")
			//replace
			currenttime := time.Now().Local()
			TSTMP := currenttime.Format("2006-01-02-15-04-05")	
			buf.WriteString(fmt.Sprintf("    UCD_BUILD_STR = `BUILD_%v`", TSTMP))
			print(string("ok"))
			FL_WRITTEN_OK = true
		}
		
		i = strings.Index(scanner.Text(), "ULAPPH_ADMIN_EMAIL")
		if i != -1 {
			//-----------------------------
			//Configuring installation
			fmt.Printf("\n++ Processing ULAPPH_ADMIN_EMAIL...  ")
			//replace
			configValue := getFromConfig("ADMIN_ACCOUNT")		
			lineText := strings.Replace(scanner.Text(), "ULAPPH_ADMIN_EMAIL", configValue, -1)
			buf.WriteString(fmt.Sprintf("%v\n", lineText))
			print(string("ok"))
			FL_WRITTEN_OK = true
		}
		
		i = strings.Index(scanner.Text(), "] Read Datastore...")
		if i != -1 {
			//-----------------------------
			//Configuring installation
			fmt.Printf("\n++ Tagging Read Datastore...  ")
			//replace
			//c.Infof("[R001] Read Datastore...")
			SPL := strings.Split(scanner.Text(), "[R")
			//SPL2 := strings.Split(SPL[1], "]")
			//thisNum := SPL2[0]
			ERR1CTR = ERR1CTR + 1
			ERRCODE := padNumberWithZero(4, ERR1CTR)
			buf.WriteString(fmt.Sprintf("%v[R%v] Read Datastore...\n", SPL[0], ERRCODE))
			print(string("ok"))
			FL_WRITTEN_OK = true
		}
		
		i = strings.Index(scanner.Text(), "] Write Datastore...")
		if i != -1 {
			//-----------------------------
			//Configuring installation
			fmt.Printf("\n++ Tagging Write Datastore...  ")
			//replace
			//c.Infof("[R001] Write Datastore...")
			SPL := strings.Split(scanner.Text(), "[W")
			//SPL2 := strings.Split(SPL[1], "]")
			//thisNum := SPL2[0]
			ERR1CTR = ERR1CTR + 1
			ERRCODE := padNumberWithZero(4, ERR1CTR)
			buf.WriteString(fmt.Sprintf("%v[W%v] Write Datastore...\n", SPL[0], ERRCODE))
			print(string("ok"))
			FL_WRITTEN_OK = true
		}
		
		//Process config values
		if FL_START_CUST_CONFIGS == true && FL_END_CUST_CONFIGS == false {
			
			fmt.Printf("\nLINE: %v", scanner.Text())
			//loop from configs
			for _, cfg := range Config.Configs {
				
				//tStr := fmt.Sprintf("%v =", cfg.Item)
				//i = strings.Index(scanner.Text(), tStr)
				SPL := strings.Split(scanner.Text(), "=")
				if strings.TrimSpace(SPL[0]) == cfg.Item {
				//if i != -1 {
					//-----------------------------
					//Configuring installation
					fmt.Printf("\n++ Processing cfg.Item...  ")
					fmt.Printf("\n+++ cfg.Item: %v", cfg.Item)
					fmt.Printf("\n+++ cfg.Format: %v", cfg.Format)
					fmt.Printf("\n+++ cfg.Status: %v", cfg.Status)
					fmt.Printf("\n+++ cfg.Value: %v", cfg.Value)

					switch cfg.Format {
						case "Flag":
							buf.WriteString(fmt.Sprintf("    %v = %v\n", cfg.Item, cfg.Value))
						case "Number":
							num, err := strconv.Atoi(cfg.Value)
							if err != nil {
								fmt.Printf("\nERROR: %v", cfg)
								panic(err)
								break
							}
							buf.WriteString(fmt.Sprintf("    %v = %v\n", cfg.Item, num))
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
							} else {
								buf.WriteString(fmt.Sprintf("    %v = `%v`\n", cfg.Item, cfg.Value))
							}
					}
					print(string("ok"))
					FL_WRITTEN_OK = true
					break
				}	
			}

		
			
		}
		
		//--------------------------------
		i = strings.Index(scanner.Text(), "// !!!CONFIG-STARTS-HERE!!!")
		if i != -1 {
			//-----------------------------
			//Configuring installation
			fmt.Printf("\n++ Setting flag FL_START_CUST_CONFIGS...  ")
			FL_START_CUST_CONFIGS = true
		}
		
		i = strings.Index(scanner.Text(), "// !!!CONFIG-ENDS-HERE!!!")
		if i != -1 {
			//-----------------------------
			//Configuring installation
			fmt.Printf("\n++ Setting flag FL_START_CUST_CONFIGS...  ")
			FL_END_CUST_CONFIGS = true
		}
		
		//--------------------------
		//default
		//for those not edited lines
		if FL_WRITTEN_OK == false {
			buf.WriteString(fmt.Sprintf("%v\n", scanner.Text()))
		}

		}

    	}
	//if FL_VALID_FILE == false {
	//	fmt.Printf("\n+ ERROR: Invalid file!!!")
	//}
	//-----------------------------
	//Writing modified file
	fmt.Printf("\n+ Writing modified file...  ")	
    	err = ioutil.WriteFile(Config.Installer[0].Dir+"/main3.go", buf.Bytes(), 0644)
    	if err != nil {
        log.Fatal(err)
		stdout = []byte(fmt.Sprintf("%v",err))
    	} else {
		stdout = []byte("ok")
	}
	print(string(stdout))
	
	return err
}

//get value of config item
func getFromConfig(key string) (retItem string) {
	//loop from configs
	for _, cfg := range Config.Configs {
		if cfg.Item == key {
			retItem =  cfg.Value
			fmt.Printf("\n+++ cfg.Item: %v", cfg.Item)
			fmt.Printf("\n+++ cfg.Format: %v", cfg.Format)
			fmt.Printf("\n+++ cfg.Status: %v", cfg.Status)
			fmt.Printf("\n+++ cfg.Value: %v", cfg.Value)
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
