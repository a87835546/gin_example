package controllers

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"golang.org/x/oauth2"
	"google.golang.org/api/youtube/v3"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
)

type ApiCtl struct {
}

//const missingClientSecretsMessage = `
//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
//AIzaSyCHV5I6BklqfUGo3rCeuHAsrqqynmmswuU -- API 密钥 1
//`

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile, err := tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}
	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(cacheFile, tok)
	}
	return config.Client(ctx, tok)
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape("youtube-go-quickstart.json")), err
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func handleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}

func channelsListByUsername(service *youtube.Service, part string, forUsername string) {
	list := make([]string, 0, 0)
	list = append(list, part)
	call := service.Channels.List(part)
	call.MaxResults(100)
	call = call.ForUsername(forUsername)
	response, err := call.Do()
	handleError(err, "")
	if len(response.Items) > 0 {
		fmt.Println(fmt.Sprintf("This channel's ID is %s. Its title is '%s', "+
			"and it has %d views.",
			response.Items[0].Id,
			response.Items[0].Snippet.Title,
			response.Items[0].Statistics.ViewCount))
	}
}
func search(service *youtube.Service, part string, forUsername string) {
	list := make([]string, 0, 0)
	list = append(list, part)
	call := service.VideoCategories.List(part)
	call = call.Hl("")
	call = call.RegionCode("US")
	//call = call.ForUsername(forUsername)
	response, err := call.Do()
	handleError(err, "")
	if len(response.Items) > 0 {
		fmt.Println(fmt.Sprintf("This channel's ID is %s. Its title is '%s', "+
			"and it has views.",
			response.Items[0].Id,
			response.Items[0].Snippet.Title))
	}
}

//func main() {
//	ctx := context.Background()
//
//	//b, err := ioutil.ReadFile("./client_secret.json")
//	//if err != nil {
//	//	log.Fatalf("Unable to read client secret file: %v", err)
//	//}
//	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey("AIzaSyCHV5I6BklqfUGo3rCeuHAsrqqynmmswuU"))
//
//	// If modifying these scopes, delete your previously saved credentials
//	// at ~/.credentials/youtube-go-quickstart.json
//	//config, err := google.ConfigFromJSON(b, youtube.YoutubeReadonlyScope)
//	//if err != nil {
//	//	log.Fatalf("Unable to parse client secret file to config: %v", err)
//	//}
//	//client := getClient(ctx, config)
//	//service, err := youtube.New(client)
//
//	handleError(err, "Error creating YouTube client")
//
//	channelsListByUsername(youtubeService, "snippet,contentDetails,statistics", "GoogleDevelopers")
//}
