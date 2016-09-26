package graph

import (
	"testing"
)

func TestVertex(t *testing.T) {
	t.Run("ConnectTo", func(t *testing.T) {
		u := &Vertex{}
		v := &Vertex{}
		cost := 0

		u.ConnectTo(v, cost)

		if len(u.Edges) != 1 {
			t.Errorf("Expected 1 edge, got %d", len(u.Edges))
		}

		from := u.Edges[0].From
		if from != u {
			t.Errorf("Expected u.edges[0].from = %p, got %p", u, from)
		}

		to := u.Edges[0].To
		if to != v {
			t.Errorf("Expected u.edges[0].to = %p, got %p", v, u)
		}

		c := u.Edges[0].Cost
		if c != cost {
			t.Errorf("Expected u.edges[0].cost = %d, got %d", cost, c)
		}
	})
}

func TestUndirectedGraph(t *testing.T) {
	g := NewUndirectedGraph()

	vertex1 := &Vertex{Name: "1"}
	vertex2 := &Vertex{Name: "2"}

	g.Add(vertex1)
	g.Add(vertex2)

	t.Run("Add", func(t *testing.T) {
		if len(g.vertices) != 2 {
			t.Errorf("Expected 2 vertices, got %d", len(g.vertices))
		}
	})

	t.Run("Get", func(t *testing.T) {
		r := g.Get("1")
		if g.Get("1") != vertex1 {
			t.Errorf("Expected vertex1@%p, got %p", vertex1, r)
		}

		r = g.Get("2")
		if g.Get("2") != vertex2 {
			t.Errorf("Expected vertex2@%p, got %p", vertex2, r)
		}
	})

	t.Run("Connect", func(t *testing.T) {
		g.Connect(vertex1, vertex2, 100)

		v1 := g.Get("1")
		v2 := g.Get("2")

		r := len(v1.Edges)
		if r != 1 {
			t.Errorf("Expected vertex '1' to have 1 edge, got %d", r)
		}

		r = v1.Edges[0].Cost
		if r != 100 {
			t.Errorf("Expected edge from '1' to '2' to have cost 100, got %d", r)
		}

		r = len(v2.Edges)
		if r != 1 {
			t.Errorf("Expected vertex '2' to have 1 edge, got %d", r)
		}

		r = v2.Edges[0].Cost
		if r != 100 {
			t.Errorf("Expected edge from '2' to '1' to have cost 100, got %d", r)
		}
	})
}
