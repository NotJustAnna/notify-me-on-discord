package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"io"
	"math/big"
	"net/http"
	"os"
	"strings"
)

func main() {
	router := gin.Default()
	router.POST("/qbittorrent", qbittorrent)
	router.Run("0.0.0.0:8080")
}

const QBittorrentIconUrl = "https://upload.wikimedia.org/wikipedia/commons/thumb/6/66/New_qBittorrent_Logo.svg/240px-New_qBittorrent_Logo.svg.png"

func qbittorrent(ctx *gin.Context) {
	b, _ := io.ReadAll(ctx.Request.Body)

	information := strings.Split(string(b), "|/|")

	size := new(big.Int)
	size.SetString(information[1], 10)

	postBody, _ := json.Marshal(map[string]interface{}{
		"content": fmt.Sprintf("%s A torrent finished downloading!", os.Getenv("USER_TO_MENTION")),
		"embeds": []map[string]interface{}{
			{
				"title": information[0],
				"description": fmt.Sprintf(
					"Size: **%s**\n\n%s\n%s\n%s",
					humanize.BigIBytes(size),
					fmt.Sprintf("· Go to [qBittorrent](%s)", os.Getenv("URL_QBITTORRENT")),
					fmt.Sprintf("· Go to [FileBrowser](%s)", os.Getenv("URL_FILEBROWSER")),
					fmt.Sprintf("· Go to [Jellyfin](%s)", os.Getenv("URL_JELLYFIN")),
				),
				"color": 5814783,
			},
		},
		"username":   "qBittorrent",
		"avatar_url": QBittorrentIconUrl,
	})

	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(os.Getenv("WEBHOOK_URL"), "application/json", responseBody)

	defer resp.Body.Close()

	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	} else {
		ctx.String(http.StatusOK, "OK")
	}
}
