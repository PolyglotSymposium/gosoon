package gosoon_test

import (
    . "gosoon"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

type Blank struct {}

type OneAttribute struct {
    GilliRocks bool
}

type OneStringAttribute struct {
    Phrase string
}

type TwoAttributes struct {
    KeithWouldLikeThisNameCuzCSharp int
    Phrase string
}

type MockHasPhraseString struct {
}

func (self MockHasPhraseString) AttributeValue(foo string) string {
    if foo == "Phrase" {
        return "Phrase's value"
    }
    return ""
}

type MockEmptyParsedJson struct {
}

func (self MockEmptyParsedJson) AttributeValue(foo string) string {
    return ""
}

var _ = Describe("Gosoon", func() {
    Describe(".Deserialize", func() {
        Context("When given an empty JSON array", func() {
            var (
                subject TwoAttributes
            )

            BeforeEach(func() {
                subject = Deserialize(MockEmptyParsedJson{}, TwoAttributes{}).(TwoAttributes)
            })

            It("Should have the default value for its integer field", func() {
                Expect(subject.KeithWouldLikeThisNameCuzCSharp).To(Equal(0))
            })

            It("Should have the default value for its string field", func() {
                Expect(subject.Phrase).To(Equal(""))
            })
        })

        Context("When given a JSON object, but serializing to an object with no attributes", func() {
            var (
                subject Blank
            )

            BeforeEach(func() {
                subject = Deserialize(MockHasPhraseString{}, Blank{}).(Blank)
            })

            It("Should return a blank object", func() {
                Expect(subject).To(Equal(Blank{}))
            })
        })

        Context("When none of the properties on the JSON match the object's properties", func() {
            var (
                subject OneAttribute
            )

            BeforeEach(func() {
                subject = Deserialize(MockHasPhraseString{}, OneAttribute{}).(OneAttribute)
            })

            It("Should return a blank object", func() {
                Expect(subject).To(Equal(OneAttribute{}))
            })
        })


        Context("When given a JSON object with a string field (1 char), whose attribute matches the databag's", func() {
            var (
                subject OneStringAttribute
            )

            BeforeEach(func() {
                subject = Deserialize(MockHasPhraseString{}, OneStringAttribute{}).(OneStringAttribute)
            })

            It("Should have the JSON value for the string field", func() {
                Expect(subject.Phrase).To(Equal("Phrase's value"))
            })
        })
    })
})