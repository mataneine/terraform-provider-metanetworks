package metanetworks

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTypeSetToStringSlice(s *schema.Set) []string {
	valuesList := s.List()
	values := make([]string, len(valuesList))
	for i := 0; i < len(valuesList); i++ {
		values[i] = fmt.Sprint(valuesList[i])
	}

	return values
}
