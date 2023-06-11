package mypackage

import "fmt"

// the uppercase means the access is public
// the lowercase means the access is private

// CarPublic car publico
type CarPublic struct {
	Brand string
	Year  int
}

func (myCar CarPublic) String() string {
	return fmt.Sprintf("Soy un %s modelo %d", myCar.Brand, myCar.Year)
}

type carPrivate struct {
	brand string
	year  int
}
