package engine

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/st3v3nmw/sidekick/internal/llms"
	"github.com/zcalusic/sysinfo"
)

type Engine struct {
	provider llms.Provider
}

func NewEngine(provider, model, apiKey string) (*Engine, error) {
	switch provider {
	case "openrouter":
		return &Engine{
			provider: llms.NewOpenRouter(model, apiKey),
		}, nil
	default:
		return nil, fmt.Errorf("unknown provider %s", provider)
	}
}

func (e *Engine) Loop(request string) {
	info := sysinfo.SysInfo{}
	info.GetSysInfo()

	prompt := `You are an intelligent command execution assistant called Sidekick.
Your role is to:
- Understand the user's request in natural language
- Determine appropriate system command(s)
- Assess command risk
- Explain reasoning transparently

Execution Guidelines:
- Break complex tasks into smaller sub-tasks and proceed iteratively
- Use sub-step commands to gather more information about the environment when necessary
- Respond with only one command at a time
- Use a reasonable number of commands to complete the task

Risk Assessment Scale:
- 1-2: Safe, no system impact
- 3-5: Moderate caution
- 6-8: High risk, require user confirmation
- 9-10: Extremely dangerous, block execution

User Request: %s

Environment:
- OS: %s
- Architecture: %s
- Kernel: %s
- Device: %s

Working Directory:
- Path: %s
- Project Type: %s
- Build System: %s
- Deployment: %s

You MUST respond with ONLY JSON in the following format:
{
	"command": "<actual command>",
	"reasoning": "<why run this command - brief>",
	"risk": <0-10>,
	"assessment": "<why this risk level - brief>",
	"done": <bool> // is this the last command?
}
`

	project, err := getProjectInfo()
	if err != nil {
		log.Fatal(err)
	}

	instruction := fmt.Sprintf(
		prompt,
		request,
		info.OS.Name,
		info.OS.Architecture,
		info.Kernel.Release,
		info.Product.Name,
		project.Path,
		project.Type,
		project.BuildSystem,
		project.Deployment,
	)

	for {
		fmt.Println(instruction)

		command, err := e.provider.Complete(instruction)
		if err != nil {
			log.Fatal(err)
		}

		cmd := exec.Command("sh", "-c", command.Command)
		out, _ := cmd.Output()
		result := string(out)

		if command.Done {
			fmt.Println(result)
			break
		}

		instruction = fmt.Sprintf("Result:\n%s", result)
		time.Sleep(15)
	}
}
