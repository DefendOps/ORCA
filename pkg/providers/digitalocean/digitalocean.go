package digitalocean

import (
	"context"
	"fmt"
	"log"

	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

type DigitalOceanClient struct {
	APIToken string
	Client   *godo.Client
}

func NewDigitalOceanClient(apiToken string) *DigitalOceanClient {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: apiToken,
	})

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)

	client := godo.NewClient(oauthClient)

	return &DigitalOceanClient{
		APIToken: apiToken,
		Client:   client,
	}
}

func (d *DigitalOceanClient) ListAgents() {
	ctx := context.TODO()
	droplets, _, err := d.Client.Droplets.List(ctx, &godo.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing droplets: %v", err)
	}

	for _, droplet := range droplets {
		fmt.Printf("ID: %d, Name: %s\n", droplet.ID, droplet.Name)
	}
}
