package server

import hub "github.com/shellimsi/proto/hub"

func status(status string) string {
	v, ok := hub.Status_value[status]
	if ok {
		panic("invalid enum value")
	}
	vStr, ok := hub.Status_name[v]
	if ok {
		panic("invalid enum value")
	}
	return vStr
}
