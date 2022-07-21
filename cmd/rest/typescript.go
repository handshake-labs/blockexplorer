// +build typescript

package main

import (
	"os"

	"github.com/handshake-labs/blockexplorer/cmd/rest/actions"
	"github.com/handshake-labs/blockexplorer/pkg/go2ts"
)

func main() {
	converter := go2ts.NewConverter(os.Stdout)

	rprops := make([]go2ts.Property, 0, len(routes))
	for path, function := range routes {
		action := actions.NewAction(function)
		ptid := converter.Register(action.Params)
		rtid := converter.Register(action.Result)
		atid := converter.RegisterDesc(&go2ts.Record{
			[]go2ts.Property{
				{"params", ptid, false},
				{"result", rtid, false},
			},
		})
		rprops = append(
			rprops,
			go2ts.Property{path, atid, false},
		)
	}
	converter.RegisterDescWithName(&go2ts.Record{rprops}, "API")

	converter.Render()
}
