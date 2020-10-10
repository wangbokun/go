package tools

type  Autolower struct{}

func (Autolower)MonitorLevelP(level string)(l string) {
	switch level {
	case "P0":
		return "p1"
	case "P1":
		return "p2"
	case "P2":
		return "p3"
	case "P3":
		return "p4"
	default:
		return level
	}
}