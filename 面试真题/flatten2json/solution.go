package solution

import (
	"encoding/json"
	"fmt"
	"strings"
)

/*
problem:
Given a “flatten” dictionary object, whose keys are dot-separated.
For example, { ‘A’: 1, ‘B.A’: 2, ‘B.B’: 3, ‘CC.D.E’: 4, ‘CC.D.F’: 5}.
Implement a function in any language to transform it to a “nested” dictionary object.
In the above case, the nested version is like:

{
  ‘A’: 1,
  ‘B’: {
    ‘A’: 2,
    ‘B’: 3,
  },
  ‘CC’: {
    ‘D’: {
      ‘E’: 4,
      ‘F’: 5,
    }
  }
}
It’s guaranteed that no keys in dictionary are prefixes of other keys.
*/

func unFlatten(flat map[string]interface{}) string {
	m, err := flatten2Map(flat)
	if err != nil {
		fmt.Println(err)
	}

	return map2Json(m)
}

func flatten2Map(flat map[string]interface{}) (map[string]interface{}, error) {
	unflat := map[string]interface{}{}

	for key, value := range flat {
		keyParts := strings.Split(key, ".")

		// Walk the keys until we get to a leaf node.
		m := unflat
		for i, k := range keyParts[:len(keyParts)-1] {
			v, exists := m[k]
			if !exists {
				newMap := map[string]interface{}{}
				m[k] = newMap
				m = newMap
				continue
			}

			innerMap, ok := v.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("key=%v is not an object", strings.Join(keyParts[0:i+1], "."))
			}
			m = innerMap
		}

		// if input allows duplicate prefix, uncomment below codes
		/*
			leafKey := keyParts[len(keyParts)-1]
			if _, exists := m[leafKey]; exists {
				return nil, fmt.Errorf("key=%v already exists", key)
			}
		*/

		m[keyParts[len(keyParts)-1]] = value
	}

	return unflat, nil
}

func map2Json(m map[string]interface{}) string {
	jsonBytes, _ := json.MarshalIndent(m, "", "  ")
	return string(jsonBytes)
}
