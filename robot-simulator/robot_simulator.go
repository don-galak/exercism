package robot

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

type Action byte

func StartRobot(command chan Command, action chan Action) {
	for c := range command {
		action <- Action(c)
	}
	action <- 0
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
		case 0:
			report <- robot
		}
	}
}

type Action3 struct {
	robotName string
	command   rune
}

func StartRobot3(name, script string, action chan Action3, log chan string) {
	for _, command := range script {
		action <- Action3{name, command}
	}
	action <- Action3{command: 0}
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {
	for a := range action {
		index := findRobotIndex(robots, a.robotName)
		switch a.command {
		case 'A':
			switch robots[index].Dir {
			case N:
				if robots[index].Northing+1 <= extent.Max.Northing {
					robots[index].Northing++
				} else {
					log <- "bump"
				}
			case S:
				if robots[index].Northing-1 >= extent.Min.Northing {
					robots[index].Northing--
				} else {
					log <- "bump"
				}
			case W:
				if robots[index].Easting-1 >= extent.Min.Easting {
					robots[index].Easting--
				} else {
					log <- "bump"
				}
			case E:
				if robots[index].Easting+1 <= extent.Max.Easting {
					robots[index].Easting++
				} else {
					log <- "bump"
				}
			}
		case 'R':
			if robots[index].Dir == W {
				robots[index].Dir = N
			} else {
				robots[index].Dir++
			}
		case 'L':
			if robots[index].Dir == N {
				robots[index].Dir = W
			} else {
				robots[index].Dir--
			}
		case 0:
			rep <- robots
		}
	}
}

// func advance(direction *RU, border RU, step RU, log chan string) {
// 	switch step {
// 	case 1:
// 		if (*direction)+step <= border {
// 			*direction = *direction + step
// 			// robots[index].Northing++
// 		} else {
// 			log <- "bump"
// 		}
// 	}
// }

func findRobotIndex(robots []Step3Robot, name string) int {
	for i := range robots {
		if robots[i].Name == name {
			return i
		}
	}
	return -1
}
