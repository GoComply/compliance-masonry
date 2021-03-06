/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package info_test

import (
	. "github.com/opencontrol/compliance-masonry/pkg/cli/info"

	"errors"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
)

var _ = Describe("Info", func() {
	Describe("Searching Implementation Status", func() {
		var (
			workingDir string
		)
		BeforeEach(func() {
			workingDir, _ = os.Getwd()
		})
		Describe("bad inputs", func() {
			Context("When no certification is specified", func() {
				It("should return an empty slice and an error", func() {
					config := Config{}
					i, err := FindImplementationStatus(config, "partial")
					assert.Equal(GinkgoT(), []error{errors.New("Error: Missing Certification Argument")}, err)
					assert.Equal(GinkgoT(), 0, len(i.SatisfiesMap))
				})
			})
			Context("When bad / no folder location is given", func() {
				It("should return an empty slice and an error", func() {
					config := Config{Certification: "LATO"}
					i, err := FindImplementationStatus(config, "partial")
					assert.Equal(GinkgoT(), []error{errors.New("Error: `certifications` directory does exist")}, err)
					assert.Equal(GinkgoT(), 0, len(i.SatisfiesMap))
				})
			})
		})
		Context("When we search for an implementation_status", func() {
			It("should find at least one component in our test data", func() {
				config := Config{
					OpencontrolDir: filepath.Join(workingDir, "..", "..", "..", "test", "fixtures", "opencontrol_fixtures"),
					Certification:  "LATO",
				}
				i, err := FindImplementationStatus(config, "partial")
				assert.Nil(GinkgoT(), err)
				assert.Greater(GinkgoT(), len(i.ComponentList), 0)
			})
		})
		Context("When we search for the 'partial' implementation_status", func() {
			It("should find more than one in our test data", func() {
				config := Config{
					Certification:  "LATO",
					OpencontrolDir: filepath.Join(workingDir, "..", "..", "..", "test", "fixtures", "opencontrol_fixtures"),
				}
				l, err := FindImplementationStatus(config, "partial")
				assert.Nil(GinkgoT(), err)
				assert.Greater(GinkgoT(), len(l.SatisfiesMap), 1)
			})
		})
		Context("When we search for the 'planned' implementation_status", func() {
			It("should only find one in our test data", func() {
				config := Config{
					OpencontrolDir: filepath.Join(workingDir, "..", "..", "..", "test", "fixtures", "opencontrol_fixtures"),
					Certification:  "LATO",
				}
				i, err := FindImplementationStatus(config, "planned")
				assert.Nil(GinkgoT(), err)
				assert.Equal(GinkgoT(), 1, len(i.SatisfiesMap))
			})
		})
	})
})
