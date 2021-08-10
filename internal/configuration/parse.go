package configuration

import "fmt"

type lm []map[string]string

// Parse costam
func convertConfig(c interface{}) lm {
	var ret_val lm
	ai, ok := c.([]interface{})
	if ok {
		for _, ai_v := range ai {
			msi, ok := ai_v.(map[string]interface{})
			if ok {
				el := make(map[string]string)
				for msi_k, msi_v := range msi {
					s, ok := msi_v.(string)
					if ok {
						el[msi_k] = s
					}
				}
				ret_val = append(ret_val, el)
			}
		}
	}
	return ret_val
}

// Parse costam
func (c Configuration) Parse() (b error) {
	for k, v := range c.ch {
		if k != "conf" {
			logger.Printf("Key [%s] is not a configuration value - ignoring", k)
			continue
		} else {
			fmt.Println(v)
			ln := convertConfig(v)
			for ln_i, ln_v := range ln {
				fmt.Println(ln_i, ln_v)

			}

		}
	}
	return nil
}
