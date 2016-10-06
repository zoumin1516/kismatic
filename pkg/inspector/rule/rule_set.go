package rule

import (
	"io"
	"strings"
)

// DefaultRuleSet is the list of rules that are built into the inspector
const defaultRuleSet = `---
# Python 2.5+ is installed on all nodes
# This is required by ansible
- kind: Python2Version
  when: []
  supportedVersions:
   - Python 2.5
   - Python 2.6
   - Python 2.7

# Executables required by kubelet
- kind: ExecutableInPath
  when: ["worker"]
  executable: iptables
- kind: ExecutableInPath
  when: ["worker"]
  executable: iptables-save
- kind: ExecutableInPath
  when: ["worker"]
  executable: iptables-restore

# Ports used by etcd are available
- kind: TCPPortAvailable
  when: ["etcd"]
  port: 2379
- kind: TCPPortAvailable
  when: ["etcd"]
  port: 6666
- kind: TCPPortAvailable
  when: ["etcd"]
  port: 2380
- kind: TCPPortAvailable
  when: ["etcd"]
  port: 6660

# Ports used by etcd are accessible
- kind: TCPPortAccessible
  when: ["etcd"]
  port: 2379
  timeout: 5s
- kind: TCPPortAccessible
  when: ["etcd"]
  port: 6666
  timeout: 5s
- kind: TCPPortAccessible
  when: ["etcd"]
  port: 2380
  timeout: 5s
- kind: TCPPortAccessible
  when: ["etcd"]
  port: 6660
  timeout: 5s

# Ports used by K8s master are available
- kind: TCPPortAvailable
  when: ["master"]
  port: 6443
- kind: TCPPortAvailable
  when: ["master"]
  port: 8080

# Ports used by K8s master are accessible
- kind: TCPPortAccessible
  when: ["master"]
  port: 6443
  timeout: 5s
- kind: TCPPortAccessible
  when: ["master"]
  port: 8080
  timeout: 5s

# TODO: Add kismatic package checks
`

// DefaultRules returns the list of rules that are built into the inspector
func DefaultRules() []Rule {
	rules, err := UnmarshalRulesYAML([]byte(defaultRuleSet))
	if err != nil {
		// The default rules should not contain errors
		// If they do, panic so that we catch them during tests
		panic(err)
	}
	return rules
}

// DumpDefaultRules writes the default rule set to a file
func DumpDefaultRules(writer io.Writer) error {
	_, err := io.Copy(writer, strings.NewReader(defaultRuleSet))
	if err != nil {
		return err
	}
	return nil
}
