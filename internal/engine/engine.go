package engine

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/fatih/color"
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
	prompt := `You are an intelligent command execution assistant called Sidekick.
Your role is to:
- Understand the user's request in natural language
- Determine appropriate non-interactive command(s)
- Assess command risk
- Explain reasoning transparently

Execution Guidelines:
- Break complex tasks into smaller sub-tasks and proceed iteratively
- Use sub-steps to gather more information about the environment
- Use a reasonable number of steps
- Always confirm that the task is completed successfully

Risk Assessment Scale:
- 1-2: Safe, no system impact
- 3-5: Moderate caution
- 6-8: High risk, require user confirmation
- 9-10: Extremely dangerous, block execution

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

User Request: %s

You MUST respond with ONLY JSON in the following format:
{
	"command": "<actual command>",
	"why": "<Why this command? (x - risk assessment) - very very brief>",
	"risk": <0-10>,
	"done": <bool> // is the task complete?
}`

	info := sysinfo.SysInfo{}
	info.GetSysInfo()

	project, err := getProjectInfo()
	if err != nil {
		log.Fatal(err)
	}

	instruction := fmt.Sprintf(
		prompt,
		info.OS.Name,
		info.OS.Architecture,
		info.Kernel.Release,
		info.Product.Name,
		project.Path,
		strings.Join(project.Types, ", "),
		strings.Join(project.BuildSystems, ", "),
		strings.Join(project.Deployments, ", "),
		request,
	)

	steps := 1
	for {
		command, err := e.provider.Complete(instruction)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Step #%d\n", steps)
		fmt.Printf("Command: ")

		c := color.New(getCommandColor(command.Risk))
		c.Println(command.Command)
		fmt.Println(command.Why)

		if command.Risk >= 6 {
			var execute string
			fmt.Printf("Execute? [y/N]: ")
			fmt.Scan(&execute)

			execute = strings.ToLower(execute)
			if execute == "y" || execute == "yes" {
				instruction = executeCommand(command)
			} else {
				fmt.Println("Operation cancelled.")
				break
			}
		} else {
			instruction = executeCommand(command)
		}

		if command.Done {
			break
		}

		steps += 1
		fmt.Println()
	}
}

func executeCommand(command *llms.Command) (result string) {
	cmd := exec.Command("sh", "-c", command.Command)
	out, err := cmd.Output()
	r := strings.TrimSpace(string(out))

	if err != nil {
		result = fmt.Sprintf("\n%s", err.Error())
	}

	if len(r) > 0 {
		result = fmt.Sprintf("\n%s", r)
	}

	if len(result) > 0 {
		fmt.Println(result)
	}

	return fmt.Sprintf("Result from last command: %s", result)
}

func getCommandColor(riskLevel int) color.Attribute {
	if riskLevel >= 9 {
		return color.FgRed
	} else if riskLevel >= 6 {
		return color.FgMagenta
	} else if riskLevel >= 3 {
		return color.FgYellow
	} else {
		return color.FgGreen
	}
}
