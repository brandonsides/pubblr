package main

import (
	"encoding/json"
	"fmt"

	"github.com/brandonsides/pubblr/activitystreams"
	"github.com/brandonsides/pubblr/util"
)

func main() {
	subj := util.Left[activitystreams.ObjectIface, activitystreams.Link]((&activitystreams.Place{
		Object: activitystreams.Object{
			Id: "http://example.org/~jane",
		},
		Accuracy: 23.7,
	}))
	r := activitystreams.TopLevelObject{
		Context: "https://www.w3.org/ns/activitystreams",
		ObjectIface: &activitystreams.Relationship{
			Object: activitystreams.Object{
				Id: "http://example.org/~john",
			},
			Subject: subj,
		},
	}
	jsonR, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonR))
}
