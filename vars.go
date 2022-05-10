package main

import (
	"github.com/valyala/fasthttp"
	"github.com/fatih/color"
	"sync"
)

const (
	AUTOCLAIMED_USERNAME_TEMPLATE string = `{ "content": null, "embeds": [ { "description": "> ` + "`@%s`" +  ` has been successfully autoclaimed!", "color": 65443, "author": { "name": "TikTok Autoclaimer" } } ], "attachments": [] }`
	SWAPPED_USERNAME_TEMPLATE string = `{"content": null,"embeds": [{"description": "> ` + "`@%s`" + ` has been swapped to ` + "`%s`" + `","color": 14498006,"author": {"name": "TikTok Autoclaimer"}}],"attachments": []}`
	SUSPENDED_USERNAME_TEMPLATE string = `{ "content": null, "embeds": [ { "description": "> ` + "`@%s` " + `has been suspended", "color": 14498006, "author": { "name": "TikTok Monitor" } } ], "attachments": [] }`

	AUTOCLAIMED_WEBHOOK string = "https://discord.com/api/webhooks/973384378589798482/0HgMURD4aJCDD0Fkobc2fp8SaoPWNktlDnOwtgfgg_cpsnsMY0A4cAjkOcl1Jg5hlj7F"
	MONITOR_WEBHOOK string = "https://discord.com/api/webhooks/973384721365078046/AkruVQRa20qAzfELVjDmuEO6Y3yex8ZOJNbBIChcs3aO_J971gwvlg2uRLC6VFMaN8xb"

)

var (
	httpClient *fasthttp.Client

	channel chan string
	usernames []string
	
	sessions []string
	session string
	index int

	mutex sync.Mutex
	threads int
	lock bool

	attempts int64
	rs int64
	rl int64

	pink func(a ...interface{}) string = color.New(color.FgHiMagenta).SprintFunc()
	green func(a ...interface{}) string = color.New(color.FgHiGreen).SprintFunc()
)