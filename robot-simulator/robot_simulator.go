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
type Action byte

func StartRobot(command chan Command, action chan Action) {
	for c := range command {
		action <- Action(c)
	}
}

func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	for a := range action {
		switch a {
		case 'A':
			switch robot.Dir {
			case N:
				if robot.Northing+1 <= extent.Max.Northing {
					robot.Northing++
				}
			case S:
				if robot.Northing-1 >= extent.Min.Northing {
					robot.Northing--
				}
			case W:
				if robot.Easting-1 >= extent.Min.Easting {
					robot.Easting--
				}
			case E:
				if robot.Easting+1 <= extent.Max.Easting {
					robot.Easting++
				}
			}
		case 'R':
			if robot.Dir == W {
				robot.Dir = N
			} else {
				robot.Dir++
			}
		case 'L':
			if robot.Dir == N {
				robot.Dir = W
			} else {
				robot.Dir--
			}
		}
	}
	report <- robot
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
