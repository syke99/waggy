package json

import "fmt"

func BuildJSONStringFromWaggyError(typ string, ttl string, dtl string, sts int, inst string, fld string) string {

	errStr := "{"

	if typ != "" {
		errStr = fmt.Sprintf("%[1]s \"type\": \"%[2]s\",", errStr, typ)
	}

	if ttl != "" {
		if errStr[:1] != "{" {
			errStr = fmt.Sprintf("%[1]s,", errStr)
		}

		errStr = fmt.Sprintf("%[1]s \"title\": \"%[2]s\",", errStr, ttl)
	}

	if dtl != "" {
		if errStr[:1] != "{" {
			errStr = fmt.Sprintf("%[1]s,", errStr)
		}

		errStr = fmt.Sprintf("%[1]s \"detail\": \"%[2]s\",", errStr, dtl)
	}

	if sts != 0 {
		if errStr[:1] != "{" {
			errStr = fmt.Sprintf("%[1]s,", errStr)
		}

		errStr = fmt.Sprintf("%[1]s \"status\": \"%[2]d\",", errStr, sts)
	}

	if inst != "" {
		if errStr[:1] != "{" {
			errStr = fmt.Sprintf("%[1]s,", errStr)
		}

		errStr = fmt.Sprintf("%[1]s \"instance\": \"%[2]s\",", errStr, inst)
	}

	if fld != "" {
		if errStr[:1] != "{" {
			errStr = fmt.Sprintf("%[1]s,", errStr)
		}

		errStr = fmt.Sprintf("%[1]s \"field\": \"%[2]s\"", errStr, fld)
	}

	if errStr[len(errStr)-1:] == "," {
		errStr = errStr[:len(errStr)-1]
	}

	return fmt.Sprintf("%[1]s }", errStr)
}
