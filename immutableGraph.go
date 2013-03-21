package gonigsberg

import (
	"os"
	"errors"
	"bufio"
	"strings"
	"strconv"
	"stringSet"
)
/*
 Immutable graph
*/
type ImmutableGraph struct {
	adj [][]int
	nodes map[string]int
	idxToID []string
}
/*
 Creates a new Immutable graph from an edge list where edge list has the format:
       # comment
       nodeid   nodeid
       nodeid   nodeid
       ...

 returns an error if path is invalid
*/
func NewImmutableGraphFromEdgeList(path string) (*ImmutableGraph, error) {
	g := new(ImmutableGraph)
	// open file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make([]string, 0)
	// create reader
	r := bufio.NewReader(file)
	// read fist line
	line, err := r.ReadString('\n')
	// while there is no error
	for err == nil {
		if (line[0] != '\n' && line[0] != '#'){
			lines = append(lines, line)
		}
		line, err = r.ReadString('\n')
	}

	//create set of node ids
	nodes := stringSet.NewStringSet()

	for _, line := range lines{
		vals := strings.Fields(line)
		if len(vals) != 2 {
			return nil, errors.New("Invalid file format")
		}
		nodes.Add(vals[0])
		nodes.Add(vals[1])
	}
	// allocate data for graph
	g.idxToID = make([]string, len(nodes))
	g.nodes = make(map[string]int, len(nodes))
	g.adj = make([][]int, len(nodes))
	for i := 0; i < len(nodes); i++{
		g.adj[i] = make([]int, 0)
	}
	// create maps for conversion from id to idx and the reverse
	i := 0
	for k, _ := range nodes{
		g.idxToID[i] = k
		g.nodes[k] = i
		i++
	}
	
	for _, line := range lines{
                fields := strings.Fields(line)
		src := g.nodes[fields[0]]
		dest := g.nodes[fields[1]]
		g.adj[src] = append(g.adj[src], dest)
		g.adj[dest] = append(g.adj[dest],src) 
	}

	return g, nil
}
/*
 to be implemented
*/
func NewImmutableGraphFromGraphML(path string) *ImmutableGraph {
	return new(ImmutableGraph)
}
/*
 to be implemented
*/
func NewImmutableGraphFromDot(path string) *ImmutableGraph {
	return new(ImmutableGraph)
}

func (g *ImmutableGraph) nbrsFromIdx(idx int) []int {
	return g.adj[idx]
} 

/*
 returns neighbors ids of id as a slice of strings
*/
func (g *ImmutableGraph) Neighbors(id string) []string{
	idx := g.nodes[id]
		nbrs := make([]string, len(g.adj[idx]))
	for i, v := range g.adj[idx]{
		nbrs[i] = g.idxToID[v]
	}
	return nbrs
}
/*
 counts number of paths of length depth from source to each vertex reachable from 
 source in depth steps
*/
func (g *ImmutableGraph) CountPaths(source string, depth int) map[string] int {
	// get source idx from source id
	sourceIdx := g.nodes[source]
	// visited nodes for each level
	levels := make([][]int, depth)
	// set for marking edges
	traversedEdges := stringSet.NewStringSet()
	levels[0] = make([]int, len(g.adj[sourceIdx]))
	copy(levels[0], g.adj[sourceIdx])
	// for each level
	for i := 1; i < depth; i++ {
		levels[i] = make([]int, 0, 16)
		for _, v := range levels[i-1]{
			for _, val := range g.adj[v]{
				// create forward and reverse edge keys
				edgeKey1 := strconv.Itoa(v) + strconv.Itoa(val)
				edgeKey2 := strconv.Itoa(val) + strconv.Itoa(v)
				// if we have not traversed this edge
				if (!traversedEdges.Contains(edgeKey1)) || (!traversedEdges.Contains(edgeKey2)){
					// mark edge as traversed
					traversedEdges.Add(edgeKey1)
					traversedEdges.Add(edgeKey2)
					// visit node
					levels[i] = append(levels[i], val)
				}
			}
		}
	}
	// create return map 
	retMap := make(map[string]int)
	// get counts of paths
	for _, v := range levels[depth -1]{
		retMap[g.idxToID[v]]++
	}
	return retMap
}

/*
 returns connected components for immutable graph g
*/
func (g *ImmutableGraph) GetConnectedComponents() [][]string{
	components := make([][]string, 0)
	return components 
}