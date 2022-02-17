package cfg

import (
	"encoding/json"
	goconfig "github.com/iglin/go-config"
	"io/ioutil"
	"os"
)

var ChainConfig *goconfig.Config

func init() {
	var c ChainCfg
	config := c.Init()
	config.SetDefaults()
	err := WriteConfig(ConfigPath, config)
	if err != nil {
		panic(err)
	}
	ChainConfig = goconfig.NewConfig(ConfigPath, goconfig.Json)

}

func (c ChainCfg) Init() *ChainCfg {
	return new(ChainCfg)
}

func (c *ChainCfg) SetDefaults() {
	c.Development = false
	c.GenesisBlock = ""
	c.Name = "QuantosMainNet"
	c.Version = "1.0.0"
	c.Currency = new(CurrencyInfo)
	c.Currency.Pairs = make(map[int]*CurrencyPairs)
	c.ChainCode = ""
	c.Private = false
	c.Public = true
	c.Vault = new(VaultCfg)
	c.Vault.VaultURL = "https://localhost"
	c.Vault.VaultCredsName = "VAULT_TOKEN"
	c.Vault.VaultTLSCertFile = ""
	c.Vault.VaultTLSKeyFile = ""
	c.Vault.VaultSecretsPath = "qsecrets/"
	c.P2P = new(P2PCfg)
}

func WriteConfig(path string, data *ChainCfg) error {

	exists, fileBytes := ConfigFileExists(path, true)

	if !exists {
		d, err := json.Marshal(data)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(path, d, 0600)

		return err
	}
	err := ioutil.WriteFile(path, fileBytes, 0600)
	if err != nil {
		return err
	}
	return nil

}

func ConfigFileExists(path string, truncate bool) (bool, []byte) {
	if truncate == true {
		return false, nil
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		if err == os.ErrNotExist {
			return false, nil
		}
		panic(err)
	}
	return true, b

}
