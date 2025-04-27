package ramdomCode

// "github.com/Grs2080w/grp_server/core/utils/ramdomCode"

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomCode() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
