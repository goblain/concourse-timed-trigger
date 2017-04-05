package resource

import (
	"fmt"
	"time"
)

type Trigger struct {
	Id        int `json:"id"`
	CreatedAt time.Time `json:"createdTimestamp"`
	FiresAt   time.Time `json:"fireTimestamp"`
	Comment   string `json:"comment"`
	State     string `json:"state"`
}

type Triggers struct {
	triggers []Trigger
}

// Satisfy sort.Interface with Len, Less and Swap
func (t *Triggers) Len() int {
	return len(t.triggers)
}

func (t *Triggers) Less(i, j int) {

}

func (t *Triggers) Swap() {

}

func (t *Triggers) Load() {

}

func (t *Triggers) Store() {

}

func (t *Triggers) addTrigger(firesAt time.Time, comment string, reset bool) {
	trigger := Trigger{Id: len(t.triggers)+1, CreatedAt: time.Now(), FiresAt: firesAt, Comment: comment, State: "waiting"}
	t.triggers = append(t.triggers, trigger)
}

func (t *Triggers) getLatestFiredTrigger() {

}

// edges - if true, include start and end trigger in result
func (t *Triggers) getTriggersBetweenIds(start string, end string, edges bool) (Triggers, error) {
	startError := fmt.Errorf("Start of trigger range not found")
	endError := fmt.Errorf("End of trigger range not found")
	var list Triggers
	if startError != nil {
		return nil, startError
	}
	if endError != nil {
		return nil, endError
	}

	return list, nil
}
