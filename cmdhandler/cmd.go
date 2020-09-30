package cmdhandler

import (
	"strings"

	"github.com/TaKeO90/deploybot/botio"
	"github.com/TaKeO90/deploybot/msg"
)

const (
	ErrCmdNotFound string = "sorry this command is not available"
)

// Info structure holds element that we might use to reply to a command.
type Info struct {
	fstname string
	lstname string
	id      int
}

// CMD structure describes the body of a command.
type CMD struct {
	LeadCmd    string
	CmdArgs    []string
	SenderInfo Info
}

// NewCmd
func NewCmd(cmd string, data msg.SenderData) *CMD {
	var (
		leadcmd string
		cmdargs []string
	)
	if cmd != "" && data.Id != 0 {
		leadcmd = strings.Trim(strings.Split(cmd, " ")[0], "/")
		cmdargs = strings.Split(cmd, " ")[1:]
		info := &Info{data.Firstname, data.Lastname, data.Id}
		return &CMD{leadcmd, cmdargs, *info}
	}
	return nil
}

// HandleCmd
func (c *CMD) HandleCmd() error {
	b := new(botio.SendMessage)
	b.Chatid = c.SenderInfo.id
	switch c.LeadCmd {
	case "deploy":
		b.Text = "Comming soon"
		_, err := b.Send()
		if err != nil {
			return err
		}
	case "getstat":
		//TODO: ...
	default:
		b.Text = ErrCmdNotFound
		b.Send()
	}
	return nil
}
