package apis

import (
	"apiServer/utils"
	"fmt"
	"net/http"
)

func getWsUrl(path string) string {
	wsURL := utils.GetEnv("WS_API_URL", "http://localhost:4000")
	return wsURL + path
}

func NotifyResourceBooked(resourceID uint) {
	url := fmt.Sprintf("/resources/%d/book", resourceID)
	url = getWsUrl(url)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("ERR: Cannot notify ws")
		fmt.Println(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("ERR: Cannot notify ws")
		fmt.Println(err)
	}
	defer resp.Body.Close()
}

func NotifyAllResourcesRestored() {
	url := getWsUrl("/resources/restore")

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("ERR: Cannot notify ws")
		fmt.Println(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("ERR: Cannot notify ws")
		fmt.Println(err)
	}
	defer resp.Body.Close()
}
