package utils

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/catalogfi/wbtc-garden/model"
	"github.com/fatih/color"
	"github.com/tyler-smith/go-bip39"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var HomeDir string

var ErrMnemonicFileMissing = errors.New("mnemonic file missing")

func init() {
	var err error
	HomeDir, err = os.UserHomeDir()
	if err != nil {
		log.Fatal("failed to get $HOME value")
	}
}

func DefaultCobiDirectory() string {
	return filepath.Join(HomeDir, ".cobi")
}

func DefaultMnemonicPath() string {
	return filepath.Join(HomeDir, ".cobi", "MNEMONIC")
}

func DefaultConfigPath() string {
	return filepath.Join(HomeDir, ".cobi", "config.json")
}

func DefaultInstantWalletDBDialector() gorm.Dialector {
	return sqlite.Open(filepath.Join(HomeDir, ".cobi", "btciw.db"))
}

func GetIWConfig(isIW bool) model.InstantWalletConfig {
	if isIW {
		return model.InstantWalletConfig{
			Dialector: DefaultInstantWalletDBDialector(),
		}
	}
	return model.InstantWalletConfig{}
}

func LoadMnemonic(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, ErrMnemonicFileMissing
		}
		return nil, err
	}
	return bip39.EntropyFromMnemonic(string(data))
}

func NewMnemonic(path string) ([]byte, error) {
	entropy := make([]byte, 32)
	if _, err := rand.Read(entropy[:]); err != nil {
		return nil, err
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}
	color.Green("Generating new mnemonic:\n[ %v ]", mnemonic)

	// Create the `.cobi` folder if not exist
	if _, err := os.Stat(filepath.Dir(path)); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(filepath.Dir(path), 0700)
		if err != nil {
			return nil, err
		}
	}

	// Create the mnemonic file
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = file.WriteString(mnemonic)
	if err != nil {
		return nil, err
	}
	return entropy[:], nil
}

func LoadConfigFromFile(file string) model.Network {
	config := model.Network{}
	configFile, err := os.ReadFile(file)
	if err != nil {
		return config
	}
	json.Unmarshal(configFile, &config)
	return config
}
