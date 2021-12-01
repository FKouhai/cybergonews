package bot

import (
	"fmt"
	"go-cyberbot/config"
	"math/rand"
	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly"
	//"github.com/mmcdole/gofeed"
	//"golang.org/x/text/message"
)
var BotId string
var goBot *discordgo.Session

func Start() {
  goBot, err := discordgo.New("Bot " + config.Token)
  if err != nil {
    fmt.Println(err.Error(), "Error at startup")
    return
  }
  u, err := goBot.User("@me")
  if err != nil {
    fmt.Println(err.Error(), "Error in @me ")
    return
  }
  BotId = u.ID
  goBot.AddHandler(messageHandler)
  err = goBot.Open()
  if err != nil {
    fmt.Println(err.Error(), "Error in gobot open")
    return
  }
  fmt.Println("Bot is running")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
  if m.Author.ID == BotId {
    return
  }
  if m.Content == "~"{
    _, _ = s.ChannelMessageSend(m.ChannelID, "in order to run the bot you need to specify a media outlet(cyberhoot)")
  }
  if m.Content == "~ cyberhoot" {
    Artis()
    _, _ = s.ChannelMessageSend(m.ChannelID, Artis())
  }
}

func Artis() string{
  var link string
  var allLinks[]string
  c := colly.NewCollector()
  colly.AllowedDomains("cyberhoot.com")
  colly.MaxDepth(2)
  colly.Async(true)
  c.OnRequest(func(r *colly.Request) {
})
  c.OnHTML(".entry-title", func(e *colly.HTMLElement){
    link = e.ChildAttr("a", "href")
    fmt.Println(link)
    allLinks = append(allLinks, link)
 //   link = e.Request.AbsoluteURL(e.Attr("href"))
})
  c.Visit("https://cyberhoot.com/category/blog/")
  r := rand.Intn(len(allLinks))
  return allLinks[r]
}
