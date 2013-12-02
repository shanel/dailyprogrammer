package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "sort"
    "strconv"
    "strings"
)

type RoomData struct {
    number  int
    avg     int
    v_count int
}

type RoomDataArray []RoomData

func (rd RoomDataArray) Len() int {
    return len(rd)
}

func (rd RoomDataArray) Swap(i, j int) {
    rd[i], rd[j] = rd[j], rd[i]
}

func (rd RoomDataArray) Less(i, j int) bool {
    return rd[i].number < rd[j].number
}

func main() {
    // room_number: {visitor_number: [time_in, time_out]}
    var rooms = map[int]map[int][2]int{}
    fh, err := os.Open("rooms.txt")
    if err != nil {
        panic("Couldn't open file")
    }
    scanner := bufio.NewScanner(fh)
    for scanner.Scan() {
        line := scanner.Text()
        chunks := strings.Split(line, " ")
        visitor, _ := strconv.Atoi(chunks[0])
        room, _ := strconv.Atoi(chunks[1])
        status := chunks[2]
        timestamp, _ := strconv.Atoi(chunks[3])
        if _, prs := rooms[room]; prs {
            if _, v_prs := rooms[room][visitor]; v_prs {
                if rooms[room][visitor][1] == 0 {
                    if status == "I" {
                        panic("A visitor can't be in two places at once!")
                    } else {
                        rooms[room][visitor] = [2]int{rooms[room][visitor][0], timestamp}
                    }
                } else {
                    // someone has come back to the same room
                    t := rooms[room][visitor][1] - rooms[room][visitor][0]
                    rooms[room][visitor] = [2]int{timestamp - t, 0}
                }
            } else {
                rooms[room][visitor] = [2]int{timestamp, 0}
            }
        } else {
            rooms[room] = map[int][2]int{visitor: [2]int{timestamp, 0}}
        }
    }
    var rda RoomDataArray
    for room, _ := range rooms {
        v_count := 0
        time := 0
        for _, visitor_data := range rooms[room] {
            v_count++
            time = time + (visitor_data[1] - visitor_data[0] + 1)
        }
        avg := int(math.Floor(float64(time) / float64(v_count)))
        rda = append(rda, RoomData{number: room, avg: avg, v_count: v_count})
    }
    sort.Sort(rda)
    for _, val := range rda {
        fmt.Printf("Room %d, %d minute average visit, %d visitors total\n", val.number, val.avg, val.v_count)
    }
}
