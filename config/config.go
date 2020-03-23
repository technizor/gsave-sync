package config

import (
	"encoding/json"
	"log"
	"os"
	"os/user"
	"regexp"

	"github.com/gobuffalo/packr"
)

// SaveSettings A SaveSettings
type SaveSettings struct {
	Timestamp bool
	History   int64
	Path      string
	Compress  bool
	Prefix    string
}

// DefaultSaveSettings The default SaveSettings
func DefaultSaveSettings() SaveSettings {
	return SaveSettings{
		Timestamp: true,
		History:   3,
		Path:      ".",
		Compress:  false,
		Prefix:    "save",
	}
}

// Game config
type Game struct {
	Name      string
	ShortName string
	Path      map[string]string
}

// Config A config for multiple game saves
type Config struct {
	Games        []Game
	SaveSettings SaveSettings
}

func getUsrHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

var gamesBox = packr.NewBox("../games")
var usrHomeDir = getUsrHomeDir()
var usrHomeRe = regexp.MustCompile(`(^~)(.*$)`)

func pathReplace(fpath string) string {
	s := usrHomeRe.ReplaceAllString(fpath, usrHomeDir+`$2`)

	return s
}

// ReadConfig Read a config JSON
func ReadConfig(fpath string) {
	f, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	enc := json.NewEncoder(os.Stdout)

	var v Config
	if err := dec.Decode(&v); err != nil {
		log.Println(err)
		return
	}
	if err := enc.Encode(&v); err != nil {
		log.Println(err)
	}
}

// WriteConfig Write a config JSON
func WriteConfig(fpath string, v Config) {
	f, err := os.Create(fpath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	enc := json.NewEncoder(f)
	enc2 := json.NewEncoder(os.Stdout)

	if err := enc.Encode(&v); err != nil {
		log.Println(err)
	}
	if err := enc2.Encode(&v); err != nil {
		log.Println(err)
	}
}
