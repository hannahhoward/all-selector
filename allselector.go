package allselector

import (
	"github.com/ipld/go-ipld-prime"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/ipld/go-ipld-prime/traversal/selector/builder"
	"github.com/ipld/go-ipld-prime/traversal/selector"

)

// AllSelector is a selector for a whole dag traversal
var AllSelector ipld.Node

func init() {
	ssb := builder.NewSelectorSpecBuilder(basicnode.Style.Any)
	AllSelector = ssb.ExploreRecursive(selector.RecursionLimitNone(), ssb.ExploreAll(ssb.ExploreRecursiveEdge())).Node()
}
