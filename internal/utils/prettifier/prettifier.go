package prettifier

import "encoding/json"

func PrettyPrint(i interface{}) ([]byte, error) {
	return json.MarshalIndent(i, "", "\t")
}
