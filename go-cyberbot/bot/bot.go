package bot

import (
	"fmt"
	"go-cyberbot/config"
	"math/rand"
	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly"
)
var BotId string
var goBot *discordgo.Session
type Site struct {
  domain string
  pattern string
  visit string
}
func Start() {
  goBot, err := discordgo.New("Bot " + config.Token) //starting the bot
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
  cyberhoot := new(Site)
  cyberhoot.domain = "cyberhoot.com"
  cyberhoot.pattern = ".entry-title"
  cyberhoot.visit = "https://cyberhoot.com/category/blog/"
  darkReading := new(Site)
  darkReading.domain = "www.darkreading.com"
  darkReading.pattern = ".article-title"
  darkReading.visit = "https://www.darkreading.com/attacks-breaches"
  techRep := new(Site)
  techRep.domain = "www.techrepublic.com"
  //techRep.pattern = ".title"
  techRep.pattern = ".title"
  techRep.visit = "https://www.techrepublic.com/topic/security/"
  fmt.Println(cyberhoot)
  if m.Author.ID == BotId {
    return
  }
  if m.Content == "~ help"{
    _, _ = s.ChannelMessageSend(m.ChannelID, "in order to run the bot you need to specify a media outlet(cyberhoot, dark-reader, tech-rep)")
  }
  if m.Content == "~ cyberhoot" {
    _, _ = s.ChannelMessageSend(m.ChannelID, Artis(cyberhoot.domain, cyberhoot.pattern, cyberhoot.visit))
  }
  if m.Content == "~ dark-reader" {
    _, _ = s.ChannelMessageSend(m.ChannelID, Artis(darkReading.domain, darkReading.pattern, darkReading.visit))
  }
  if m.Content == "~ tech-rep" {
    _, _ = s.ChannelMessageSend(m.ChannelID, Artis(techRep.domain, techRep.pattern, techRep.visit))
  }
}

func Artis(domain string, pattern string, visit string ) string{
  var link string
  var allLinks[]string
  c := colly.NewCollector()
  colly.AllowedDomains(domain)
  colly.MaxDepth(2)
  colly.Async(true)
  c.OnRequest(func(r *colly.Request) {
})
  c.OnHTML(pattern, func(e *colly.HTMLElement){
    if domain == "www.darkreading.com" {
      fmt.Println(domain)
      link = e.Attr("href")
      fmt.Println(link)
    }else {
      link = e.ChildAttr("a", "href")
    }
    fmt.Println(link)
    allLinks = append(allLinks, link)
})
  c.Visit(visit)
  r := rand.Intn(len(allLinks))
  //r := rand.Intn(6) + 1
  fmt.Println(r)
  return allLinks[r]
}
