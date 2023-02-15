package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func Test_positive_test(t *testing.T) {
	g := NewGomegaWithT(t)

	Employee := Employee{
		Name:       "Watan Prasai",
		Email:      "watanp2012@gmail.com",
		EmployeeID: "J63217655",
	}

	ok, err := govalidator.ValidateStruct(Employee)

	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}
