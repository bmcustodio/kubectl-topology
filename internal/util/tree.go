package util

import (
	"fmt"
	"sort"
	"strings"
)

const (
	firstElemPrefix  = `├─`
	middleElemPrefix = `│ `
	lastElemPrefix   = `└─`
	indent           = "  "
)

type treeElement struct {
	name string
	subs []*treeElement
}

func newTreeElement(name string) *treeElement {
	return &treeElement{
		name: name,
		subs: make([]*treeElement, 0),
	}
}

func (te *treeElement) addSub(sub *treeElement) {
	te.subs = append(te.subs, sub)
	sort.SliceStable(te.subs, func(i, j int) bool {
		return te.subs[i].name < te.subs[j].name
	})
}

func (te *treeElement) treeString(linePrefix string) string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%s\n", te.name))
	if te.subs == nil || len(te.subs) == 0 {
		return sb.String()
	}
	subSb := strings.Builder{}
	for i, sub := range te.subs {
		last := i == len(te.subs)-1
		var subLinePrefix string
		if !last {
			subSb.WriteString(firstElemPrefix)
			subLinePrefix = middleElemPrefix
		} else {
			subSb.WriteString(lastElemPrefix)
			subLinePrefix = indent
		}
		subSb.WriteString(sub.treeString(subLinePrefix))
	}
	sb.WriteString(linesPrefix(subSb.String(), linePrefix))
	return sb.String()
}

func PrintNodeTree(nodes []Node) error {
	regionMap := groupNodes(nodes, func(node Node) string {
		return node.Region
	})
	te := newTreeElement("cluster")
	for region, nodes := range regionMap {
		zoneMap := groupNodes(nodes, func(node Node) string {
			return node.Zone
		})
		regionTE := newTreeElement(region)
		te.addSub(regionTE)
		for zone, nodes := range zoneMap {
			zoneTE := newTreeElement(zone)
			regionTE.addSub(zoneTE)
			for _, node := range nodes {
				zoneTE.addSub(newTreeElement(node.Name))
			}
		}
	}
	treeStr := te.treeString("")
	fmt.Println(treeStr)
	return nil
}

func PrintPodTree(pods []Pod) error {
	regionMap := groupPods(pods, func(pod Pod) string {
		return pod.Node.Region
	})
	te := newTreeElement("cluster")
	for region, pods := range regionMap {
		zoneMap := groupPods(pods, func(pod Pod) string {
			return pod.Node.Zone
		})
		regionTE := newTreeElement(region)
		te.addSub(regionTE)
		for zone, pods := range zoneMap {
			nodeMap := groupPods(pods, func(pod Pod) string {
				return pod.Node.Name
			})
			zoneTE := newTreeElement(zone)
			regionTE.addSub(zoneTE)
			for node, pods := range nodeMap {
				nodeTE := newTreeElement(node)
				zoneTE.addSub(nodeTE)
				for _, pod := range pods {
					nodeTE.addSub(newTreeElement(fmt.Sprintf("%s - %s", pod.Namespace, pod.Name)))
				}
			}
		}
	}
	treeStr := te.treeString("")
	fmt.Println(treeStr)
	return nil
}

func groupNodes(nodes []Node, keyFunc func(Node) string) map[string][]Node {
	m := make(map[string][]Node)
	for _, node := range nodes {
		key := keyFunc(node)
		n, ok := m[key]
		if !ok {
			n = make([]Node, 0)
		}
		m[key] = append(n, node)
	}
	return m
}

func groupPods(pods []Pod, keyFunc func(Pod) string) map[string][]Pod {
	m := make(map[string][]Pod)
	for _, pod := range pods {
		key := keyFunc(pod)
		n, ok := m[key]
		if !ok {
			n = make([]Pod, 0)
		}
		m[key] = append(n, pod)
	}
	return m
}

func linesPrefix(s string, prefix string) string {
	lines := strings.Split(s, "\n")
	sb := strings.Builder{}
	for _, line := range lines {
		if len(line) != 0 {
			sb.WriteString(prefix + line + "\n")
		}
	}
	return sb.String()
}
