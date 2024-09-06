package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type GlobalConfig struct {
    LogLevel        string `yaml:"log_level"`
    OutputDirectory string `yaml:"output_directory"`
}

type FleetConfig struct {
	Size       int           `yaml:"size"`
	Timeout    int           `yaml:"timeout"`
	AgentConfig AgentConfig   `yaml:"agent_config"`
	Provider   ProviderConfig `yaml:"provider"`
}

type AgentConfig struct {
	ScanDepth       int `yaml:"scan_depth"`
	Retries         int `yaml:"retries"`
	DelayBetweenScans int `yaml:"delay_between_scans"`
}

type ProviderConfig struct {
	AWS          *AWSConfig   `yaml:"aws,omitempty"`
	DigitalOcean *DOConfig    `yaml:"digitalocean,omitempty"`
}

type AWSConfig struct {
	Enabled      bool   `yaml:"enabled"`
	Region       string `yaml:"region"`
	InstanceType string `yaml:"instance_type"`
}

type DOConfig struct {
	Enabled      bool   `yaml:"enabled"`
	DropletSize  string `yaml:"droplet_size"`
	Datacenter   string `yaml:"datacenter"`
}

type TargetType string

const (
    TargetTypeIP    TargetType = "ipv4"
    TargetTypeDomain TargetType = "domain"
)

type TargetConfig struct {
    Type       TargetType `yaml:"type"`
    Domain     string    `yaml:"domain,omitempty"`
    IPRange    *IPRangeConfig `yaml:"ipv4_space,omitempty"`
}

type IPRangeConfig struct {
    RangeStart string   `yaml:"range_start"`
    RangeEnd   string   `yaml:"range_end"`
}

type BackgroundServiceConfig struct {
    Enabled                bool `yaml:"enabled"`
    Interval               int  `yaml:"interval"`
    MaxConcurrentScans     int  `yaml:"max_concurrent_scans"`
}

type CloudProviderConfig struct {
    AWS struct {
        Enabled         bool `yaml:"enabled"`
        AccessKeyID     string `yaml:"access_key_id"`
        SecretAccessKey string `yaml:"secret_access_key"`
        Region          string `yaml:"region"`
    } `yaml:"aws"`
    DigitalOcean struct {
        Enabled         bool `yaml:"enabled"`
        Token string `yaml:"token"`
    } `yaml:"digitalocean"`
}

type KubernetesConfig struct {
    Enabled      bool   `yaml:"enabled"`
    ConfigPath   string `yaml:"config_path"`
    Namespace    string `yaml:"namespace"`
    APIServer    string `yaml:"api_server"`
    Token        string `yaml:"token"`
}

type ReportingConfig struct {
    Enabled  bool   `yaml:"enabled"`
    Endpoint string `yaml:"endpoint"`
    APIKey   string `yaml:"api_key"`
}


type Config struct {
    Global            GlobalConfig               `yaml:"global"`
    Fleet             FleetConfig                `yaml:"fleet"`
    Targets           []TargetConfig             `yaml:"targets"`
    BackgroundService BackgroundServiceConfig    `yaml:"background_service"`
    CloudProviders    CloudProviderConfig        `yaml:"cloud_providers"`
    Kubernetes        KubernetesConfig           `yaml:"kubernetes"`
    Reporting         ReportingConfig            `yaml:"reporting"`
}

func handleFleet(fleet FleetConfig) {
	if !fleet.Provider.AWS.Enabled && !fleet.Provider.DigitalOcean.Enabled {
		log.Fatalf("No provider is enabled for this fleet")
	}
}

func LoadConfig(filename string) (*Config, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("error reading config file: %v", err)
    }

    var config Config
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        return nil, fmt.Errorf("error parsing config file: %v", err)
    }

    for _, target := range config.Targets {
        switch target.Type {
        case TargetTypeIP:
            if target.IPRange == nil {
                return nil, fmt.Errorf("IP range config missing for target type 'ipv4'")
            }
        case TargetTypeDomain:
            if target.Domain == "" {
                return nil, fmt.Errorf("domain config missing for target type 'domain'")
            }
        default:
            return nil, fmt.Errorf("unknown target type: %s", target.Type)
        }
    }

    handleFleet(config.Fleet)

    // TODO: Add Full checks for the configuration while passing checks for integrations too; 

    return &config, nil
}