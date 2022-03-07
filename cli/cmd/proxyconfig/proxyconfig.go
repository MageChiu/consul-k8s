package proxyconfig

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/hashicorp/consul-k8s/cli/common"
	"github.com/hashicorp/consul-k8s/cli/common/flag"
	"github.com/hashicorp/consul-k8s/cli/common/terminal"
	"github.com/hashicorp/consul-k8s/cli/format"
	"k8s.io/client-go/kubernetes"
)

// Command is the proxy-config command.
type Command struct {
	*common.BaseCommand

	set *flag.Sets

	flagPodName     string
	flagNamespace   string
	flagFullConfig  bool
	flagFormat      string
	flagKubeConfig  string
	flagKubeContext string

	kubernetes kubernetes.Interface

	once sync.Once
	help string
}

// Run queries the Kubernetes Pod specified and returns its proxy configuration.
func (c *Command) Run(args []string) int {
	c.once.Do(c.init)
	c.Log.ResetNamed("proxy-config")
	defer common.CloseWithError(c.BaseCommand)

	if err := c.set.Parse(args); err != nil {
		c.UI.Output("Error parsing flags: "+err.Error(), terminal.WithErrorStyle())
		return 1
	}

	if err := c.validateFlags(); err != nil {
		c.UI.Output("Error validating flags: "+err.Error(), terminal.WithErrorStyle())
		return 1
	}

	if err := c.setupKubernetes(); err != nil {
		c.UI.Output("Error setting up Kubernetes client: "+err.Error(), terminal.WithErrorStyle())
		return 1
	}

	config, err := c.fetchConfig()
	if err != nil {
		c.UI.Output("Error fetching configuration for " + c.flagPodName + ": " + err.Error())
		return 1
	}

	c.UI.Output("Proxy configuration for "+c.flagPodName+"in namespace "+c.flagNamespace, terminal.WithHeaderStyle())
	c.outputConfig(config)

	return 0
}

func (c *Command) Help() string {
	c.once.Do(c.init)
	return c.Synopsis() + "\n\nUsage: consul-k8s proxy-config [flags]\n\n" + c.help
}

func (c *Command) Synopsis() string {
	return "Get the proxy configuration for a Kubernetes Pod."
}

func (c *Command) init() {
	// Set up the flags.
	c.set = flag.NewSets()

	f := c.set.NewSet("Command Options")
	f.StringVar(&flag.StringVar{
		Name:    "pod",
		Usage:   "The name of the Kubernetes Pod to query.",
		Aliases: []string{"p"},
		Target:  &c.flagPodName,
	})
	f.StringVar(&flag.StringVar{
		Name:    "namespace",
		Usage:   "The Namespace of the Kubernetes Pod to query.",
		Aliases: []string{"n"},
		Target:  &c.flagNamespace,
		Default: "default",
	})
	f.BoolVar(&flag.BoolVar{
		Name:   "full-config",
		Usage:  "Return the full proxy configuration.",
		Target: &c.flagFullConfig,
	})
	f.StringVar(&flag.StringVar{
		Name:    "format",
		Usage:   "The output format (JSON, YAML).",
		Aliases: []string{"o"},
		Target:  &c.flagFormat,
	})

	f = c.set.NewSet("Global Options")
	f.StringVar(&flag.StringVar{
		Name:    "kubeconfig",
		Usage:   "The path to the Kubernetes config file.",
		Aliases: []string{"c"},
		Target:  &c.flagKubeConfig,
	})
	f.StringVar(&flag.StringVar{
		Name:   "context",
		Usage:  "The name of the Kubernetes context to use.",
		Target: &c.flagKubeContext,
	})

	c.help = c.set.Help()

	c.Init()
}

func (c *Command) validateFlags() error {
	if (len(c.set.Args())) > 0 {
		return fmt.Errorf("non-flag arguments given: %s", strings.Join(c.set.Args(), ", "))
	}

	if c.flagPodName == "" {
		return fmt.Errorf("pod must be specified (e.g. -pod podname)")
	}

	return nil
}

func (c *Command) setupKubernetes() error {
	if c.kubernetes != nil {
		return nil
	}

	var err error
	c.kubernetes, err = common.CreateKubernetesClient(c.flagKubeConfig, c.flagKubeContext)
	return err
}

func (c *Command) fetchConfig() (string, error) {
	// This will use the Kubernetes API in the final version.
	output, err := exec.Command(
		"kubectl", "exec", c.flagPodName, "--namespace", c.flagNamespace,
		"-c", "envoy-sidecar", "--", "wget", "-qO-", "127.0.0.1:19000/config_dump",
	).Output()

	if err != nil {
		return "", err
	}

	return string(output), nil
}

func (c *Command) outputConfig(config string) {
	if !c.flagFullConfig {
		c.UI.Output(format.FormatEnvoyConfig(config))
		return
	}

	c.UI.Output(config)
}
