package commands

import (
	"github.com/MathisBurger/AnalAsia/pkg/colors"
	"github.com/MathisBurger/AnalAsia/pkg/system"
	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
	"strconv"
)

// Information command
// Returns some information about
// the bot and its system
func InfoCommand(s *discordgo.Session, m *discordgo.MessageCreate) {

	emb := embed.NewEmbed()
	emb.SetTitle("Information")
	emb.SetColor(colors.Blue)
	emb.SetDescription("Information about the bot")
	emb.AddField("Version:", "v0.0.1")

	sysInfo := system.GetSystemInformation()
	emb.AddField("CPU:", sysInfo.CPU)
	emb.AddField("RAM:", strconv.Itoa(int(sysInfo.RAM_USED))+"/"+strconv.Itoa(int(sysInfo.RAM_All))+"MB")
	emb.AddField("SWAP::", strconv.Itoa(int(sysInfo.SWAP_USED))+"/"+strconv.Itoa(int(sysInfo.SWAP_All))+"MB")
	emb.AddField("disk:", strconv.Itoa(int(sysInfo.Disk))+"GB")
	emb.AddField("OS:", sysInfo.Platform)

	emb.SetURL("https://github.com/MathisBurger/AnalAsia")
	emb.SetFooter("©2021 Mathis Burger")

	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, emb.MessageEmbed)
}
