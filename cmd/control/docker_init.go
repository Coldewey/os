package control

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"syscall"
	"time"

	"github.com/codegangsta/cli"
	"github.com/rancher/os/log"
	"github.com/rancher/os/util"
)

const (
	dockerConf = "/var/lib/rancher/conf/docker"
	dockerDone = "/run/docker-done"
	dockerLog  = "/var/log/docker.log"
)

func dockerInitAction(c *cli.Context) error {
	// TODO: this should be replaced by a "Console ready event watcher"
	for {
		if _, err := os.Stat(consoleDone); err == nil {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}

	dockerBin := ""
	dockerPaths := []string{
		"/usr/bin",
		"/opt/bin",
		"/usr/local/bin",
		"/var/lib/rancher/docker",
	}
	for _, binPath := range dockerPaths {
		if util.ExistsAndExecutable(path.Join(binPath, "dockerd")) {
			dockerBin = path.Join(binPath, "dockerd")
			break
		}
	}
	if dockerBin == "" {
		for _, binPath := range dockerPaths {
			if util.ExistsAndExecutable(path.Join(binPath, "docker")) {
				dockerBin = path.Join(binPath, "docker")
				break
			}
		}
	}
	if dockerBin == "" {
		err := fmt.Errorf("Failed to find either dockerd or docker binaries")
		log.Error(err)
		return err
	}
	log.Infof("Found %s", dockerBin)

	if err := syscall.Mount("", "/", "", syscall.MS_SHARED|syscall.MS_REC, ""); err != nil {
		log.Error(err)
	}
	if err := syscall.Mount("", "/run", "", syscall.MS_SHARED|syscall.MS_REC, ""); err != nil {
		log.Error(err)
	}

	mountInfo, err := ioutil.ReadFile("/proc/self/mountinfo")
	if err != nil {
		return err
	}

	for _, mount := range strings.Split(string(mountInfo), "\n") {
		if strings.Contains(mount, "/var/lib/docker /var/lib/docker") && strings.Contains(mount, "rootfs") {
			os.Setenv("DOCKER_RAMDISK", "1")
		}
	}

	args := []string{
		"bash",
		"-c",
		fmt.Sprintf(`[ -e %s ] && source %s; exec /usr/bin/dockerlaunch %s %s $DOCKER_OPTS >> %s 2>&1`, dockerConf, dockerConf, dockerBin, strings.Join(c.Args(), " "), dockerLog),
	}

	// TODO: this should be replaced by a "Docker ready event watcher"
	if err := ioutil.WriteFile(dockerDone, []byte(CurrentEngine()), 0644); err != nil {
		log.Error(err)
	}

	return syscall.Exec("/bin/bash", args, os.Environ())
}
