package confredis

import "fmt"

func keyWithPrefix(p, k string) string {
	return fmt.Sprintf("%s_%s", p, k)
}
