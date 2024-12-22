package engine

import (
	"fmt"
	"os"
)

var (
	projectTypes = map[string]string{
		"CMakeLists.txt":     "C/C++",
		"go.mod":             "Go",
		"go.sum":             "Go",
		"build.gradle":       "Java",
		"pom.xml":            "Java",
		"package.json":       "JavaScript",
		"package-lock.json":  "JavaScript",
		"yarn.lock":          "JavaScript",
		"composer.json":      "PHP",
		"composer-lock.json": "PHP",
		"Pipfile":            "Python",
		"Pipfile.lock":       "Python",
		"pyproject.toml":     "Python",
		"poetry.lock":        "Python",
		"requirements.txt":   "Python",
		"setup.cfg":          "Python",
		"setup.py":           "Python",
		"Gemfile":            "Ruby",
		"Gemfile.lock":       "Ruby",
		"Rakefile":           "Ruby",
		"Cargo.toml":         "Rust",
		"Cargo.lock":         "Rust",
		"tsconfig.json":      "TypeScript",
	}

	buildSystems = map[string]string{
		"Gemfile":           "bundler",
		"Cargo.toml":        "cargo",
		"CMakeLists.txt":    "cmake",
		"composer.json":     "composer",
		"build.gradle":      "gradle",
		"Makefile":          "make",
		"pom.xml":           "maven",
		"meson.build":       "meson",
		"ninja.build":       "ninja",
		"package.json":      "npm",
		"package-lock.json": "npm",
		"Pipfile":           "pip",
		"requirements.txt":  "pip",
		"poetry.lock":       "poetry",
		"pyproject.toml":    "poetry",
		"pnpm-lock.yaml":    "pnpm",
		"Rakefile":          "rake",
		"setup.cfg":         "setuptools",
		"setup.py":          "setuptools",
		"yarn.lock":         "yarn",
	}

	deploymentSystems = map[string]string{
		"ansible.yml":         "Ansible",
		"playbook.yml":        "Ansible",
		"Dockerfile":          "Docker",
		".dockerfile":         "Docker",
		"docker-compose.yaml": "Docker Compose",
		"docker-compose.yml":  "Docker Compose",
		"compose.yaml":        "Docker Compose",
		"compose.yml":         "Docker Compose",
		"helm.yaml":           "Helm",
		"kubernetes.yaml":     "Kubernetes",
		"kubernetes.yml":      "Kubernetes",
		"k8s.yaml":            "Kubernetes",
		"k8s.yml":             "Kubernetes",
		"terraform.tf":        "Terraform",
		"main.tf":             "Terraform",
	}
)

type ProjectInfo struct {
	Path        string
	Type        string
	BuildSystem string
	Deployment  string
}

func getProjectInfo() (*ProjectInfo, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("cannot get current directory: %w", err)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("cannot read current directory: %w", err)
	}

	info := &ProjectInfo{
		Path:        dir,
		Type:        "N/A",
		BuildSystem: "N/A",
		Deployment:  "N/A",
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if projectType, ok := projectTypes[name]; ok {
			info.Type = projectType
		}

		if buildSystem, ok := buildSystems[name]; ok {
			info.BuildSystem = buildSystem
		}

		if deployment, ok := deploymentSystems[name]; ok {
			info.Deployment = deployment
		}
	}

	return info, nil
}
