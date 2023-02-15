package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func Test_negative_name_test(t *testing.T) {
	g := NewGomegaWithT(t)

	Employee := Employee{
		Name:       "",
		Email:      "watanp2012@gmail.com",
		EmployeeID: "J63217655",
	}

	ok, err := govalidator.ValidateStruct(Employee)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("Please enter your name"))
}
