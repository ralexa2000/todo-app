package help

import (
	"fmt"
	commandPkg "gitlab.ozon.dev/ralexa2000/todo-bot/internal/pkg/bot/command"
	"strings"
)

func New(extendedMap map[string][2]string) commandPkg.Interface {
	if extendedMap == nil {
		extendedMap = map[string][2]string{}
	}
	return &command{
		extendedMap: extendedMap,
	}
}

type command struct {
	extendedMap map[string][2]string
}

func (c *command) Name() string {
	return "help"
}

func (c *command) Arguments() string {
	return ""
}

func (c *command) Description() string {
	return "list all commands"
}

func cmdHelpBuild(name string, arguments string, description string) string {
	if arguments == "" {
		return fmt.Sprintf("/%s - %s", name, description)
	}
	return fmt.Sprintf("/%s %s - %s", name, arguments, description)
}

func (c *command) Process(_ string, _ string) string {
	result := []string{
		cmdHelpBuild(c.Name(), c.Arguments(), c.Description()),
	}
	for cmdName, cmdValue := range c.extendedMap {
		cmdArguments, cmdDescription := cmdValue[0], cmdValue[1]
		result = append(result, cmdHelpBuild(cmdName, cmdArguments, cmdDescription))
	}
	return strings.Join(result, "\n")
}
