package nozzle_test

import (
	"errors"

	"github.com/cloudfoundry/sonde-go/events"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cf-platform-eng/splunk-firehose-nozzle/nozzle"
)

var _ = Describe("Nozzle", func() {
	var (
		eventChannel chan *events.Envelope
		errorChannel chan error
	)

	BeforeEach(func() {
		eventChannel = make(chan *events.Envelope)
		errorChannel = make(chan error)
	})

	It("returns error on error channel", func() {
		nozzle := NewSplunkForwarder(nil, errorChannel)
		go func() {
			errorChannel <- errors.New("Fail")
		}()
		err := nozzle.Forward()

		Expect(err).To(Equal(errors.New("Fail")))
	})
})