package cfg

var ConfigPath = "./chaincfg.json"

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
	Currency     *CurrencyInfo
	P2P          *P2PCfg
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
	Stakes                []string
	MainCurrencyHashID    string
	StakingCurrencyHashID string
}

type CurrencyInfo struct {
	DenomName            string
	DenomCoinbaseAddress string
	MaxAvailable         int64
	AvailableAtStart     int64
	BaseValuePairID      string
	Pairs                map[int]*CurrencyPairs
	CanStake             bool
	IsDefault            bool
}

type CurrencyPairs struct {
	ID          string
	PairName    string
	Pair2IsFiat bool
	Pair1       struct {
		Name          string
		Value         int64
		Pair2ForPair1 int64
		Address       string
	}
	Pair2 struct {
		Name          string
		Value         int64
		Pair1ForPair1 int64
		Address       string
	}
}

type RosettaCfg struct {
}
