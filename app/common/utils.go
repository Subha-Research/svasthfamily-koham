package common

import (
	"time"

	"github.com/Subha-Research/svasthfamily-koham/app/constants"
)

type TimeUtil struct {
}

func (tu *TimeUtil) CurrentTimeInTimezone() *time.Time {
	time_location, _ := time.LoadLocation(constants.Timezone)
	time := time.Now().In(time_location)
	return &time
}

func (tu *TimeUtil) CurrentTimeInUTC() *time.Time {
	time := time.Now().UTC()
	return &time
}
