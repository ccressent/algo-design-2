package graph

type Edge struct {
	To   *Vertex
	From *Vertex
	Cost int
}

type Vertex struct {
	Name  string
	Edges []*Edge
}

func (v *Vertex) ConnectTo(u *Vertex, cost int) {
	e := &Edge{To: u, From: v, Cost: cost}
	v.Edges = append(v.Edges, e)
}

type UndirectedGraph struct {
	hashmap  map[string]int
	vertices []*Vertex
}

func NewUndirectedGraph() *UndirectedGraph {
	return &UndirectedGraph{hashmap: make(map[string]int)}
}

func (g *UndirectedGraph) Add(v *Vertex) {
	g.vertices = append(g.vertices, v)
	g.hashmap[v.Name] = len(g.vertices) - 1
}

func (g UndirectedGraph) Get(name string) *Vertex {
	index, ok := g.hashmap[name]
	if !ok {
		return nil
	}

	return g.vertices[index]
}

func (g UndirectedGraph) Connect(u, v *Vertex, cost int) {
	u.ConnectTo(v, cost)
	v.ConnectTo(u, cost)
}

func (g UndirectedGraph) MST_Prim() []*Edge {
	var x []*Vertex
	var t []*Edge

	if len(g.vertices) == 0 {
		return nil
	}

	x = append(x, g.vertices[0])

	bestEdge := func() *Edge {
		var best *Edge

		in := func(vs []*Vertex, v *Vertex) bool {
			for i := range vs {
				if vs[i] == v {
					return true
				}
			}
			return false
		}

		for _, v := range x {
			for _, e := range v.Edges {
				// Skip edges to vertices already in x
				if in(x, e.To) {
					continue
				}

				if best == nil {
					best = e
					continue
				}

				if e.Cost < best.Cost {
					best = e
				}
			}
		}

		return best
	}

	// While all the vertices of g are not in x
	for len(x) != len(g.vertices) {
		e := bestEdge()
		t = append(t, e)
		x = append(x, e.To)
	}

	return t
}
