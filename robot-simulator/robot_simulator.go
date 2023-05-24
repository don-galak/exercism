package robot

import "fmt"

type Action byte
type Action3 struct {
	name   string
	action Action
}

const move, left, right, stop = 'A', 'L', 'R', 'S'
const N, E, S, W Dir = 0, 1, 2, 3

var directions = [...]string{"north", "east", "south", "west"}

func (r Rect) contains(p Pos) bool {
	return p.Northing <= r.Max.Northing && p.Northing >= r.Min.Northing &&
		p.Easting <= r.Max.Easting && p.Easting >= r.Min.Easting
}
func (d Dir) String() string { return directions[d] }
func (d Dir) right() Dir     { return (d + 1) % 4 }
func (d Dir) left() Dir      { return (d + 3) % 4 }
func Right()                 { Step1Robot.Dir = Step1Robot.right() }
func Left()                  { Step1Robot.Dir = Step1Robot.left() }
func Advance() {
	p := advance(Pos{RU(Step1Robot.X), RU(Step1Robot.Y)}, Step1Robot.Dir)
	Step1Robot.X, Step1Robot.Y = int(p.Easting), int(p.Northing)
}
func StartRobot(command chan Command, action chan Action) {
	for c := range command {
		action <- Action(c)
	}
	close(action)
}
func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	for a := range action {
		switch a {
		case move:
			p := advance(robot.Pos, robot.Dir)
			if extent.contains(p) {
				robot.Pos = p
			}
		case right:
			robot.Dir = robot.right()
		case left:
			robot.Dir = robot.left()
		}
	}
	report <- robot
	close(report)
}

var posMap = map[Dir]Pos{N: {0, 1}, E: {1, 0}, S: {0, -1}, W: {-1, 0}}

func advance(p Pos, d Dir) Pos {
	v := posMap[d]
	return Pos{p.Easting + v.Easting, p.Northing + v.Northing}
}
func StartRobot3(name, script string, action chan Action3, _ chan string) {
	if name == "" {
		return
	}
	for _, c := range script {
		action <- Action3{name, Action(c)}
	}
	action <- Action3{name, stop}
}
func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {
	defer close(rep)
	positions, occupancies := map[string]Step2Robot{}, map[Pos]string{}
	for _, r := range robots {
		if r.Name == "" {
			log <- "empty robot name"
			return
		} else if _, ok := positions[r.Name]; ok {
			log <- fmt.Sprintf("duplicate name: %s", r.Name)
			return
		} else if _, ok := occupancies[r.Pos]; ok {
			log <- fmt.Sprintf("already occupied: %v", r.Pos)
			return
		} else if !extent.contains(r.Pos) {
			log <- fmt.Sprintf("%s outside room", r.Name)
			return
		}
		positions[r.Name] = r.Step2Robot
		occupancies[r.Pos] = r.Name
	}
	var report []Step3Robot
	for len(positions) > 0 {
		a := <-action
		r, ok := positions[a.name]
		if !ok {
			log <- "unknown robot: " + a.name
			return
		}
		switch a.action {
		case move:
			p := advance(r.Pos, r.Dir)
			if !extent.contains(p) {
				log <- "bumped into a wall at " + r.Dir.String()
				continue
			}
			if occupant, occupied := occupancies[p]; occupied {
				log <- fmt.Sprintf("%v already occupied by %s", p, occupant)
				continue
			}
			delete(occupancies, r.Pos)
			r.Pos = p
			occupancies[r.Pos] = a.name
		case right:
			r.Dir = r.right()
		case left:
			r.Dir = r.left()
		case stop:
			delete(positions, a.name)
			report = append(report, Step3Robot{a.name, r})
			continue
		default:
			log <- "unknown action: " + string(a.action)
			return
		}
		positions[a.name] = r
	}
	rep <- report
}
