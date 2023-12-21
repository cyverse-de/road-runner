package main

import (
	"context"
	"os/exec"

	"github.com/spf13/viper"
)

// Creates a command that can be used to run docker-compose.
func DockerComposeCommand(cfg *viper.Viper, args ...string) *exec.Cmd {
	return DockerComposeCommandContext(cfg, context.Background(), args...)
}

// Creates a command that can be used to run docker.
func DockerCommand(cfg *viper.Viper, args ...string) *exec.Cmd {
	return DockerCommandContext(cfg, context.Background(), args...)
}

// Creates a command that can be used to run docker-compose in a context.
func DockerComposeCommandContext(cfg *viper.Viper, ctx context.Context, args ...string) *exec.Cmd {
	dockerComposePath := cfg.GetString("docker-compose.path")
	if dockerComposePath != "" {
		return exec.CommandContext(ctx, dockerComposePath, args...)
	}
	return DockerCommandContext(cfg, ctx, append([]string{"compose"}, args...)...)
}

// Creates a command that can be used to run docker in a context.
func DockerCommandContext(cfg *viper.Viper, ctx context.Context, args ...string) *exec.Cmd {
	dockerPath := cfg.GetString("docker.path")
	return exec.CommandContext(ctx, dockerPath, args...)
}
