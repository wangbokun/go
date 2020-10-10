package tools

type  Autolower struct{}

func (Autolower)MonitorLevelP(level string)(l string) {
	switch level {
	case "P0":
		return "P1"
	case "P1":
		return "P2"
	case "P2":
		return "P3"
	case "P3":
		return "P4"
	default:
		return level
	}
}
