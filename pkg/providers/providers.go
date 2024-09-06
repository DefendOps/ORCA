package providers

import (
	"github.com/defendops/orca/pkg/providers/aws"
	"github.com/defendops/orca/pkg/providers/digitalocean"
)

type Network struct {
    IPAddress string `json:"ip_address"`
    Type      string `json:"type"`
}

type Agent struct {
    ID           int      `json:"id"`
    Name         string   `json:"name"`
    Status       string   `json:"status"`
    Memory       int      `json:"memory"`
    VCPUs        int      `json:"vcpus"`
    Disk         int      `json:"disk"`
    Region       string   `json:"region"`
    IPv4         []Network `json:"v4"`
    IPv6         []Network `json:"v6"`
    Tags         []string `json:"tags"`
    CreatedAt    string   `json:"created_at"`
}

type Provider interface {
	ListAgents() ([]Agent, error)
}

type ProviderClient struct {
	DOClient  *digitalocean.DigitalOceanClient
	AWSClient *aws.AWSClient
}

func NewProviderClient(doToken, awsAccessKey, awsSecretKey, awsRegion string, doEnabled, awsEnabled bool) (*ProviderClient, error) {
	client := &ProviderClient{}

	if doEnabled {
		client.DOClient = digitalocean.NewDigitalOceanClient(doToken)
	}

	if awsEnabled {
		awsClient, err := aws.NewAWSClient(awsAccessKey, awsSecretKey, awsRegion)
		if err != nil {
			return nil, err
		}
		client.AWSClient = awsClient
	}

	return client, nil
}
