package factory

import (
	"log"

	"github.com/defendops/orca/pkg/config"
	"github.com/defendops/orca/pkg/providers"
	"github.com/defendops/orca/pkg/providers/aws"
	"github.com/defendops/orca/pkg/providers/digitalocean"
)

type CmdFactory struct {
	ServiceName string
	ORCAConfiguration *config.Config
	ProviderClients  *providers.ProviderClient
}

func NewCmdFactory(serviceName string) *CmdFactory {
	return &CmdFactory{
		ServiceName: serviceName,
		ORCAConfiguration: &config.Config{},
		ProviderClients:  &providers.ProviderClient{},
	}
}

func (f *CmdFactory) initializeProviders(cfg *config.Config) {
	if cfg.CloudProviders.DigitalOcean.Enabled {
		f.ProviderClients.DOClient = digitalocean.NewDigitalOceanClient(cfg.CloudProviders.DigitalOcean.Token)
	}

	if cfg.CloudProviders.AWS.Enabled {
		awsClient, err := aws.NewAWSClient(
			cfg.CloudProviders.AWS.AccessKeyID,
			cfg.CloudProviders.AWS.SecretAccessKey,
			cfg.CloudProviders.AWS.Region,
		)

		if err != nil {
			log.Fatalf("Error initializing AWS client: %v", err)
		}

		f.ProviderClients.AWSClient = awsClient
	}
}
func (f *CmdFactory) LoadConfiguration() {
	// log.Println("Loading Configuration:", f.ServiceName)
	
	configFile := "./configuration.yaml"
	cfg, err := config.LoadConfig(configFile)
    if err != nil {
        log.Fatalf("Error loading configuration: %v", err)
    }

	f.ORCAConfiguration = cfg

	f.initializeProviders(cfg)
}
