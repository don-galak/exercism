package robot

// See defs.go for other definitions

// Step 1
// Define N, E, S, W here.
const (
	N Dir = iota
	E
	S
	W
)

func Right() {
	if Step1Robot.Dir == W {
		Step1Robot.Dir = N
		return
	}
	Step1Robot.Dir++
}

func Left() {
	if Step1Robot.Dir == N {
		Step1Robot.Dir = W
		return
	}
	Step1Robot.Dir--
}

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case S:
		Step1Robot.Y--
	case W:
		Step1Robot.X--
	case E:
		Step1Robot.X++
	}
}

var directions = map[Dir]string{N: "North", E: "East", S: "South", W: "West"}

func (d Dir) String() string {
	return directions[d]
}

// Step 2
// Define Action type here.
type Action int

func StartRobot(command chan Command, action chan Action) {
	panic("Please implement the StartRobot function")
}

func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	panic("Please implement the Room function")
}

// Step 3
// Define Action3 type here.
type Action3 int

func StartRobot3(name, script string, action chan Action3, log chan string) {
	panic("Please implement the StartRobot3 function")
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {
	panic("Please implement the Room3 function")
}
