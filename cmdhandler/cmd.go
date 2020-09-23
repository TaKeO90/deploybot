package cmdhandler

import (
	"log"
	"strings"

	"github.com/TaKeO90/deploybot/botio"
)

const (
	ERRCMDNOTFOUND string = "sorry this command is not available"
)

type Info struct {
	fstname string
	lstname string
	id      int
}

type CMD struct {
	LeadCmd    string
	CmdArgs    []string
	SenderInfo Info
}

func NewCmd(cmd, sFstname, sLstname string, sId int) *CMD {
	var (
		leadcmd string
		cmdargs []string
	)
	if cmd != "" && sId != 0 {
		leadcmd = strings.Trim(strings.Split(cmd, " ")[0], "/")
		cmdargs = strings.Split(cmd, " ")[1:]
		info := &Info{sFstname, sLstname, sId}
		return &CMD{leadcmd, cmdargs, *info}
	}
	return nil
}

func (c *CMD) HandleCmd() {
	b := new(botio.SendMessage)
	b.Chatid = c.SenderInfo.id
	switch c.LeadCmd {
	case "deploy":
		b.Text = "Comming soon"
		_, err := b.Send()
		if err != nil {
			log.Fatal(err)
		}
	case "getstat":
		//TODO: ...
	default:
		b.Text = ERRCMDNOTFOUND
		b.Send()
	}
}
