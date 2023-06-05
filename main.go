package main

import (
	"context"
	"fmt"
	"github.com/hetznercloud/hcloud-go/hcloud"
	"log"
)

const apiToken = ""

func main() {

	//	client := hcloud.NewClient(hcloud.WithToken(apiToken))
	//	ctx := context.Background()
	//	GetServerTypes(ctx, client)
	//	result := CreateNewServer(ctx, client, "testServer")
	//	fmt.Println(result)
	//	time.Sleep(100 * time.Second)
	//	result := GetAllRunningServers(ctx, client)
	//	RestartServer(ctx, client, result.Server)
	//	time.Sleep(100 * time.Second)
	//	ShutDownServer(ctx, client, result.Server)
	//	time.Sleep(100 * time.Second)
	//	DeleteServer(ctx, client, result[0])
	//	time.Sleep(100 * time.Second)

}

func CreateNewServer(ctx context.Context, client *hcloud.Client, clientEmail string) hcloud.ServerCreateResult {
	// this alows you to load a snap shot if you want you can remove the type snapshot and use deafault
	result, _, err := client.Server.Create(ctx, hcloud.ServerCreateOpts{
		Location:   &hcloud.Location{ID: 1},
		Name:       clientEmail,
		Image:      &hcloud.Image{Name: "docker-ce", Type: "snapshot"},
		ServerType: &hcloud.ServerType{ID: 1},
	})
	if err != nil {
		log.Fatalf("error retrieving server: %s\n", err)

	}

	return result
}

func GetAllRunningServers(ctx context.Context, client *hcloud.Client) []*hcloud.Server {

	result, err := client.Server.All(ctx)
	if err != nil {
		log.Fatalf("error Deleteing  server: %s\n", err)

	}

	for i := 0; i < len(result); i++ {
		fmt.Println("Running server details machine number: %d", i)
		fmt.Println(result[i].Name)
		fmt.Println(result[i].ID)
		fmt.Println(result[i].Labels)
		fmt.Println("/////////end ////////// ")
	}
	fmt.Println(result)
	return result
}

func DeleteServer(ctx context.Context, client *hcloud.Client, server *hcloud.Server) *hcloud.ServerDeleteResult {

	result, _, err := client.Server.DeleteWithResult(ctx, server)
	if err != nil {
		log.Fatalf("error Deleteing  server: %s\n", err)

	}

	return result
}
func RestartServer(ctx context.Context, client *hcloud.Client, server *hcloud.Server) *hcloud.Action {

	result, _, err := client.Server.Reset(ctx, server)
	if err != nil {
		log.Fatalf("error Restarting  server: %s\n", err)

	}

	return result
}

func ShutDownServer(ctx context.Context, client *hcloud.Client, server *hcloud.Server) *hcloud.Action {

	result, _, err := client.Server.Shutdown(ctx, server)
	if err != nil {
		log.Fatalf("error shuting down  server: %s\n", err)

	}

	return result
}

func GetSSHKeys(ctx context.Context, client *hcloud.Client) {

	result, err := client.SSHKey.All(ctx)
	if err != nil {

		panic(err)

	}

	for i := 0; i < len(result); i++ {

		fmt.Println(result[i].Name)
		// use dot notation
	}

}

func GetServerTypes(ctx context.Context, client *hcloud.Client) {

	result, err := client.ServerType.All(ctx)
	if err != nil {

		panic(err)

	}

	for i := 0; i < len(result); i++ {

		fmt.Println(result[i].Name)
		// use dot notation
	}

}

func GetDataCenters(ctx context.Context, client *hcloud.Client) {

	result, err := client.Datacenter.All(ctx)
	if err != nil {

		panic(err)

	}

	for i := 0; i < len(result); i++ {

		fmt.Println(result[i].Name)
		// use dot notation
	}

}

func GetImages(ctx context.Context, client *hcloud.Client) {

	result, err := client.Image.All(ctx)
	if err != nil {

		panic(err)

	}

	for i := 0; i < len(result); i++ {

		fmt.Println(result[i].Name)
		// use dot notation
	}

}

func GetLocationOfServers(ctx context.Context, client *hcloud.Client) {

	result, err := client.Location.All(ctx)
	if err != nil {

		panic(err)

	}

	for i := 0; i < len(result); i++ {

		fmt.Println(result[i].Name)
		// use dot notation
	}

}
