package cfg

type ChainCfg struct {
	Name         string
	Version      string
	ChainCode    string
	GenesisBlock string
	Development  bool
	Private      bool
	Public       bool
	HashTable    *HashTableCfg
	Vault        *VaultCfg
	StartState   GenesisAccounts
}

type VaultCfg struct {
	VaultURL         string
	VaultCredsName   string
	VaultSecretsPath string
	VaultTLSKeyFile  string
	VaultTLSCertFile string
}

type HashTableCfg struct {
	MaxHashIDPerTable    int
	MaxChunkSize         int
	MinChunkSize         int
	MaxHashIDSizeInBytes int
	MaxMemoryUsage       int
	GenesisHashID        int
	PrettyPrint          bool
	LoadFromFile         bool
	HashTableFilesPath   string
	EnableMetrics        bool
}

type P2PCfg struct {
	Host            string
	Port            string
	PubKeyFile      string
	PrivKeyFile     string
	LocalIPFS       bool
	RemoteIPFSURL   string
	NodeID          string
	BootstrapNodes  []string
	TrustedNodes    []string
	ProtocolID      string
	ProtocolVersion string
}

type DPOSCfg struct {
	MinValidators    int
	MaxValidators    int
	MinStake         int
	MaxStake         int
	Quorum           int
	MinRewardPercent int
	MaxRewardPercent int
}

type GenesisAccounts struct {
	Addresses             []string
	Balances              []string
	MainCurrencyHashID    string
	StakingCurrencyHashID string
}

type RosettaCfg struct {
}
