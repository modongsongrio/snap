package wmap

import (
	"fmt"
)

func (w *WorkflowMap) String() string {
	var out string
	out += "Workflow\n"
	out += "\tCollect:\n"
	if w.CollectNode != nil {
		out += w.CollectNode.String("\t\t")
	} else {
		out += "\n"
	}

	return out
}

func (c *CollectWorkflowMapNode) String(pad string) string {
	var out string
	out += pad + "Metric Namespaces:\n"
	for _, x := range c.MetricsNamespaces {
		out += pad + "\t" + x + "\n"
	}
	out += "\n"
	out += pad + "Config:\n"
	for k, v := range c.Config {
		out += pad + "\t" + k + "\n"
		for x, y := range v {
			out += pad + "\t\t" + fmt.Sprintf("%s=%+v\n", x, y)
		}
	}
	out += "\n"
	out += pad + "Process Nodes:\n"
	for _, pr := range c.ProcessNodes {
		out += pr.String(pad)
	}
	out += "\n"
	out += pad + "Publish Nodes:\n"
	for _, pu := range c.PublishNodes {
		out += pu.String(pad) + "\n"
	}
	out += "\n"
	return out
}

func (p *ProcessWorkflowMapNode) String(pad string) string {
	var out string
	out += pad + fmt.Sprintf("\tName: %s\n", p.Name)
	out += pad + fmt.Sprintf("\tVersion: %d\n", p.Version)

	out += pad + "\tProcess Nodes:\n"
	for _, pr := range p.ProcessNodes {
		out += pr.String(pad + "\t")
	}
	out += pad + "\tPublish Nodes:\n"
	for _, pu := range p.PublishNodes {
		out += pu.String(pad + "\t")
	}
	return out
}

func (p *PublishWorkflowMapNode) String(pad string) string {
	var out string
	out += pad + fmt.Sprintf("\tName: %s\n", p.Name)
	out += pad + fmt.Sprintf("\tVersion: %d\n", p.Version)

	out += pad + "\tConfig:\n"
	for k, v := range p.Config {
		out += pad + "\t\t" + fmt.Sprintf("%s=%+v\n", k, v)
	}
	return out
}
