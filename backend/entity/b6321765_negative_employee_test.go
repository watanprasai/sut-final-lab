package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func Test_negative_employee(t *testing.T) {
	g := NewGomegaWithT(t)

	Employee := Employee{
		Name:       "Watan Prasai",
		Email:      "watanp2012@gmail.com",
		EmployeeID: "B6321765",
	}

	ok, err := govalidator.ValidateStruct(Employee)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("EmployeeID invalid format"))
}
