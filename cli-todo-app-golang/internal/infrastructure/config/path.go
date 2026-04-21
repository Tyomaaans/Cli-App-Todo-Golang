package config

import "path/filepath"

type PathConfig struct {
    BaseDir string
}

func NewPathConfig(baseDir string) *PathConfig {
    return &PathConfig{BaseDir: baseDir}
}

func (p *PathConfig) TodoFile() string {
    return filepath.Join(p.BaseDir, "todos.json")
}

func (p *PathConfig) UserFile() string {
    return filepath.Join(p.BaseDir, "users.json")
}

func (p *PathConfig) SessionFile() string {
    return filepath.Join(p.BaseDir, "session.json")
}

func (p *PathConfig) UserLogFile() string {
    return filepath.Join(p.BaseDir, "user_log.json")
}