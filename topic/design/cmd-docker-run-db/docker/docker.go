package docker

import (
	"bytes"
	"encoding/json"
	"os/exec"
)

// ================================================================================
// Container
type Container struct {
	id   string
	host string
}

// ================================================================================
// NewContainer
func NewContainer() (*Container, error) {
	var stdout bytes.Buffer
	cmd := exec.Command("docker", "create", "-P", "postgres:11.1-alpine")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	id := stdout.String()[:12]
	stdout.Reset()

	c := Container{
		id: id,
	}

	return &c, nil
}

// ================================================================================
// ID
func (c *Container) ID() string { return c.id }

// ================================================================================
// Host
func (c *Container) Host() string { return c.host }

// ================================================================================
// StartContainer
func (c *Container) StartContainer() error {
	var stdout bytes.Buffer
	cmd := exec.Command("docker", "start", c.id)
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	stdout.Reset()

	// =============================
	cmd = exec.Command("docker", "inspect", c.id)
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return err
	}

	var data []struct {
		NetworkSettings struct {
			Ports struct {
				TCP5432 []struct {
					HostIP   string `json:"HostIp"`
					HostPort string `json:"HostPort"`
				} `json:"5432/tcp"`
			} `json:"Ports"`
		} `json:"NetworkSettings"`
	}

	if err := json.Unmarshal(stdout.Bytes(), &data); err != nil {
		return err
	}

	// =============================
	network := data[0].NetworkSettings.Ports.TCP5432[0]
	c.host = network.HostIP + ":" + network.HostPort
	return nil
}

// ================================================================================
// StopContainer
func (c *Container) StopContainer() error {
	if err := exec.Command("docker", "stop", c.id).Run(); err != nil {
		return err
	}
	return nil
}

// ================================================================================
// RemoveContainer
func (c *Container) RemoveContainer() error {
	if err := exec.Command("docker", "rm", c.id).Run(); err != nil {
		return err
	}
	return nil
}

// ================================================================================
// DumpContainerLogs
func (c *Container) DumpContainerLogs() ([]byte, error) {
	out, err := exec.Command("docker", "logs", c.id).CombinedOutput()
	if err != nil {
		return nil, err
	}
	return out, nil
}
