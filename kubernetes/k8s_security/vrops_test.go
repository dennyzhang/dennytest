// Copyright 2018 VMware, Inc. All Rights Reserved.
//
package pks_vrops_release_test

import (
	"context"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/vmware/pks-ci/tests/integration-tests/vrops-release/helpers"
	"github.com/vmware/pks-ci/tests/test-helpers"

	boshdir "github.com/cloudfoundry/bosh-cli/director"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

const (
	cadvisorServiceName = "vrops-cadvisor"
	cadvisorNamespace   = "kube-system"
	cadvisorNodePort    = "31194"
	cadvisorServicePort = "4194"
	tileClusterPrefix   = "service-instance_"
)

var (
	client           *helpers.Client
	director         boshdir.Director
	runner           *test_helpers.KubectlRunner
	testCluster      *test_helpers.Cluster
	originalManifest string
	deploymentName   string
)

var (
	vropsURL      = test_helpers.GetEnv("VROPS_URL")
	vropsUsername = test_helpers.GetEnv("VROPS_USERNAME")
	vropsPassword = test_helpers.GetEnv("VROPS_PASSWORD")
	testEnv       = test_helpers.GetEnv("TEST_ENVIRONMENT")
)

func TestVrops(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "vRops Suite")
}

func checkDataReceiving() bool {
	id, err := helpers.GetKubernetesAdapterID(context.Background(), client, strings.TrimPrefix(deploymentName, tileClusterPrefix))
	if err != nil {
		GinkgoT().Errorf("Kubernetes adapter %s is not found\n", strings.TrimPrefix(deploymentName, tileClusterPrefix))
		return false
	}
	states, err := helpers.GetInstanceResourceState(context.Background(), client, id)
	if err != nil {
		GinkgoT().Errorf("Failed to get kubernetes adapter states: %s\n", err)
		return false
	}
	if len(states) > 0 {
		GinkgoT().Logf("Data Status: %s\n", states[0].ResourceStatus)
		return states[0].ResourceStatus == "DATA_RECEIVING"
	}
	return false
}

func checkAdapterIsDeleted() bool {
	_, err := helpers.GetKubernetesAdapterID(context.Background(), client, strings.TrimPrefix(deploymentName, tileClusterPrefix))
	if err == nil {
		GinkgoT().Logf("Kubernetes adapter %s still exists\n", strings.TrimPrefix(deploymentName, tileClusterPrefix))
		return false
	}
	if _, ok := err.(helpers.NotFound); ok {
		GinkgoT().Logf("Kubernetes adapter %s is not found\n", strings.TrimPrefix(deploymentName, tileClusterPrefix))
		return true
	}
	GinkgoT().Errorf("Failed to query kubernetes adapters: %s\n", err)
	return false
}

func revertToOriginalManifest() error {
	deployment, err := director.FindDeployment(deploymentName)
	if err != nil {
		GinkgoT().Errorf("Error while finding deployment: %s\n", err)
		return err
	}
	var updateOpts boshdir.UpdateOpts
	err = deployment.Update([]byte(originalManifest), updateOpts)
	if err != nil {
		GinkgoT().Errorf("Error updating deployment %s due to error: %s", deployment.Name(), err)
		return err
	}
	GinkgoT().Logf("Reverted to original Manifest!!")
	return nil
}

func editManifestRegexp(strKey string, strValue string) error {
	deployment, err := director.FindDeployment(deploymentName)
	if err != nil {
		GinkgoT().Errorf("Error while finding deployment: %s\n", err)
		return err
	}
	manifest := originalManifest
	re := regexp.MustCompile(strKey + ": (.*)")
	manifest = re.ReplaceAllString(manifest, strKey+": "+strValue)
	var updateOpts boshdir.UpdateOpts
	err = deployment.Update([]byte(manifest), updateOpts)
	if err != nil {
		GinkgoT().Errorf("Error updating deployment %s due to error: %s", deployment.Name(), err)
		return err
	}
	return nil
}

func runErrand(errandName string, expectedOutput string, expectedExitCode int, outputType string) error {
	deployment, err := director.FindDeployment(deploymentName)
	if err != nil {
		return err
	}
	Expect(err).NotTo(HaveOccurred())
	result, err := deployment.RunErrand(errandName, true, false, nil)
	Expect(err).NotTo(HaveOccurred())
	Expect(result).NotTo(BeNil())
	Expect(result).To(HaveLen(1))
	GinkgoT().Logf("%+v\n", result)
	Expect(result[0].ExitCode).To(Equal(expectedExitCode))
	if outputType == "stdout" {
		Expect(result[0].Stdout).To(ContainSubstring(expectedOutput))
	}
	if outputType == "stderr" {
		Expect(result[0].Stderr).To(ContainSubstring(expectedOutput))
	}
	return nil
}

var _ = Describe("vrops ci test", func() {
	Context("vrops cadvisor daemonset sanity test", func() {
		It("Should check cadvisor nodePort and port are correct", func() {
			session := runner.RunKubectlCommandInNamespace(cadvisorNamespace, "describe", "service", cadvisorServiceName)
			Eventually(session, "30s").Should(gexec.Exit(0))
			output := session.Out.Contents()
			fields := strings.Fields(string(output))
			var port, nodePort string
			getPort := func(i int, fields []string) string {
				var p string
				if len(fields) > i+2 {
					p = fields[i+2]
				}
				endIdx := strings.Index(p, "/")
				if endIdx > 0 {
					p = p[:strings.Index(p, "/")]
				}
				return p
			}
			for i, field := range fields {
				switch field {
				case "NodePort:":
					nodePort = getPort(i, fields)
				case "Port:":
					port = getPort(i, fields)
				default:
				}
			}
			Expect(nodePort).To(Equal(cadvisorNodePort))
			Expect(port).To(Equal(cadvisorServicePort))
		})

		It("Should check cadvisor daemonset is running", func() {
			stats := runner.CheckDaemonSet(cadvisorNamespace, cadvisorServiceName)
			for _, v := range stats[1:] {
				Expect(v).To(Equal(stats[0]))
			}
		})
	})

	Context("vrops errand e2e sanity test", func() {
		BeforeEach(func() {
			if testEnv != "kubo" {
				Skip("errand test is not supported by PKS tile test")
			}
		})

//		It("Run vrops register errand..", func() {
//			err := runErrand("register", "Started to monitor kubernetes cluster "+strings.TrimPrefix(deploymentName, tileClusterPrefix), 0, "stdout")
//			Expect(err).NotTo(HaveOccurred())
//		})
//
//		It("Validate vrops register errand..", func() {
//			GinkgoT().Log("Verifying the vrops data collection...")
//			Eventually(checkDataReceiving, 120*time.Second, 5*time.Second).Should(BeTrue())
//		})
//
// 		It("Run vrops unregister errand..", func() {
// 			err := runErrand("unregister", "Adapter "+strings.TrimPrefix(deploymentName, tileClusterPrefix)+" is deleted", 0, "stdout")
// 			Expect(err).NotTo(HaveOccurred())
// 		})
// 
// 		It("Validate vrops unregister errand..", func() {
// 			Eventually(checkAdapterIsDeleted, 20*time.Second, 5*time.Second).Should(BeTrue())
// 		})
//
// 		It("Run vrops insecure register errand without cert..", func() {
// 			GinkgoT().Log("Update manifest vrops_insecure to true")
// 			err := editManifestRegexp("vrops_ca", "")
// 			Expect(err).NotTo(HaveOccurred())
// 			err = editManifestRegexp("vrops_insecure", "true")
// 			Expect(err).NotTo(HaveOccurred())
// 			err = runErrand("register", "Started to monitor kubernetes cluster "+strings.TrimPrefix(deploymentName, tileClusterPrefix), 0, "stdout")
// 			Expect(err).NotTo(HaveOccurred())
// 		})
// 
// 		It("Validate vrops register errand..", func() {
// 			GinkgoT().Log("Verifying the vrops data collection...")
// 			Eventually(checkDataReceiving, 120*time.Second, 5*time.Second).Should(BeTrue())
// 		})
// 
// 		It("Run vrops unregister errand..", func() {
// 			err := runErrand("unregister", "Adapter "+strings.TrimPrefix(deploymentName, tileClusterPrefix)+" is deleted", 0, "stdout")
// 			Expect(err).NotTo(HaveOccurred())
// 			GinkgoT().Log("Revert manifest")
// 			err = revertToOriginalManifest()
// 			Expect(err).NotTo(HaveOccurred())
// 		})
// 
// 		It("Validate vrops unregister errand..", func() {
// 			Eventually(checkAdapterIsDeleted, 20*time.Second, 5*time.Second).Should(BeTrue())
// 		})
// 
// 		It("Run vrops register errand with vrops disabled..", func() {
// 			GinkgoT().Log("Update manifest vrops_enabled to Disabled")
// 			err := editManifestRegexp("vrops_enabled", "Disabled")
// 			Expect(err).NotTo(HaveOccurred())
// 			err = runErrand("register", "vROps monitoring is not enabled", 0, "stdout")
// 			Expect(err).NotTo(HaveOccurred())
// 		})
// 
// 		It("Run vrops register errand with invalid vrops url..", func() {
// 			GinkgoT().Log("Update manifest vrops_url to fake url")
// 			err := editManifestRegexp("vrops_url", "fakeurl")
// 			Expect(err).NotTo(HaveOccurred())
// 			err = runErrand("register", "unable to load certificate", 1, "stderr")
// 			Expect(err).NotTo(HaveOccurred())
// 		})

 		It("Run vrops register errand with invalid vrops username..", func() {
 			GinkgoT().Log("Update manifest vrops_admin to wrong name")
 			err := editManifestRegexp("vrops_admin", "administrator")
 			Expect(err).NotTo(HaveOccurred())
 			err = runErrand("register", "is either invalid or has expired", 1, "stdout")
 			Expect(err).NotTo(HaveOccurred())
 		})
 
 		It("Run vrops register errand with invalid vrops password..", func() {
 			GinkgoT().Log("Update manifest vrops_admin_pass to wrong one")
 			err := editManifestRegexp("vrops_admin_pass", "Alfred!23")
 			Expect(err).NotTo(HaveOccurred())
 			err = runErrand("register", "is either invalid or has expired", 1, "stdout")
 			Expect(err).NotTo(HaveOccurred())
 		})
 
 		It("Run vrops register errand with invalid vrops cert..", func() {
 			GinkgoT().Log("Update manifest vrops_ca to fake cert")
 			err := editManifestRegexp("vrops_ca", "fake cert")
 			Expect(err).NotTo(HaveOccurred())
 			err = runErrand("register", "error setting certificate verify locations", 77, "stderr")
 			Expect(err).NotTo(HaveOccurred())
 		})
 
// 		It("Run vrops register errand..", func() {
// 			GinkgoT().Log("Revert manifest")
// 			err := revertToOriginalManifest()
// 			Expect(err).NotTo(HaveOccurred())
// 			err = runErrand("register", "Started to monitor kubernetes cluster "+strings.TrimPrefix(deploymentName, tileClusterPrefix), 0, "stdout")
// 			Expect(err).NotTo(HaveOccurred())
// 		})
// 
// 		It("Run vrops unregister errand with invalid vrops url..", func() {
// 			GinkgoT().Log("Update manifest vrops_url to fake url")
// 			err := editManifestRegexp("vrops_url", "fakeurl")
// 			Expect(err).NotTo(HaveOccurred())
// 			err = runErrand("unregister", "unable to load certificate", 1, "stderr")
// 			Expect(err).NotTo(HaveOccurred())
// 		})
// 
// 		It("Run vrops unregister errand with invalid vrops username..", func() {
// 			GinkgoT().Log("Update manifest vrops_admin to wrong name")
// 			err := editManifestRegexp("vrops_admin", "administrator")
// 			Expect(err).NotTo(HaveOccurred())
// 			err = runErrand("unregister", "is either invalid or has expired", 1, "stdout")
// 			Expect(err).NotTo(HaveOccurred())
// 		})
// 
// 		It("Run vrops unregister errand with invalid vrops password..", func() {
// 			GinkgoT().Log("Update manifest vrops_admin_pass to wrong one")
// 			err := editManifestRegexp("vrops_admin_pass", "Alfred!23")
// 			Expect(err).NotTo(HaveOccurred())
// 			err = runErrand("unregister", "is either invalid or has expired", 1, "stdout")
// 			Expect(err).NotTo(HaveOccurred())
// 		})
// 
// 		It("Run vrops unregister errand with invalid vrops cert..", func() {
// 			GinkgoT().Log("Update manifest vrops_ca to fake cert")
// 			err := editManifestRegexp("vrops_ca", "fake cert")
// 			Expect(err).NotTo(HaveOccurred())
// 			err = runErrand("unregister", "error setting certificate verify locations", 77, "stderr")
// 			Expect(err).NotTo(HaveOccurred())
// 		})
// 
// 		It("Run vrops unregister errand..", func() {
// 			GinkgoT().Log("Revert manifest")
// 			err := revertToOriginalManifest()
// 			Expect(err).NotTo(HaveOccurred())
// 			err = runErrand("unregister", "Adapter "+strings.TrimPrefix(deploymentName, tileClusterPrefix)+" is deleted", 0, "stdout")
// 			Expect(err).NotTo(HaveOccurred())
// 		})
// 
// 		It("Validate vrops unregister errand..", func() {
// 			Eventually(checkAdapterIsDeleted, 20*time.Second, 5*time.Second).Should(BeTrue())
// 		})

	})

	Context("vrops tile sanity test", func() {
		BeforeEach(func() {
			if testEnv != "PKS" {
				Skip("vrops tile test is covered by e2e test in kubo test")
			}
		})

// 		// TODO: remove this errand execution after nsx-t nat rule is created during cluster creation
// 		It("Run vrops unregister errand..", func() {
// 			GinkgoT().Log("Running unregister errand...")
// 			err := runErrand("unregister", "Adapter "+strings.TrimPrefix(deploymentName, tileClusterPrefix)+" is deleted", 0, "stdout")
// 			Expect(err).NotTo(HaveOccurred())
// 		})
// 		It("Run vrops register errand..", func() {
// 			GinkgoT().Log("Running register errand...")
// 			err := runErrand("register", "Started to monitor kubernetes cluster "+strings.TrimPrefix(deploymentName, tileClusterPrefix), 0, "stdout")
// 			Expect(err).NotTo(HaveOccurred())
// 		})
// 
// 		It("Validate vrops registration", func() {
// 			GinkgoT().Log("Verifying the vrops data collection...")
// 			Eventually(checkDataReceiving, 120*time.Second, 5*time.Second).Should(BeTrue())
// 		})
	})

	var _ = BeforeSuite(func() {
		GinkgoT().Log("Pre suite setup...")

		director = test_helpers.NewDirector()
		// Fetch information about the Director.
		info, err := director.Info()
		if err != nil {
			Expect(err).NotTo(HaveOccurred())
		}
		GinkgoT().Logf("Director: %s\n", info.Name)

		client, err = helpers.NewClient(vropsURL, vropsUsername, vropsPassword, true)
		if err != nil {
			Expect(err).NotTo(HaveOccurred())
		}
		GinkgoT().Log("vrops client is initialized")

		switch testEnv {
		case "PKS":
			testCluster = test_helpers.NewCluster()
			err = testCluster.InitCluster(2)
			Expect(err).To(BeNil())
			deploymentName = testCluster.DeploymentName
		case "kubo":
			deploymentName = test_helpers.GetEnv("DEPLOYMENT_NAME")
			deployment, err := director.FindDeployment(deploymentName)
			Expect(err).NotTo(HaveOccurred())

			originalManifest, _ = deployment.Manifest()
			GinkgoT().Logf("Saving Original Manifest...")
		default:
			GinkgoT().Fatalf("\nunknown testEnv: %s\n", testEnv)
		}

		runner = test_helpers.NewKubectlRunner()
		Expect(runner).NotTo(BeNil())

	})
	var _ = AfterSuite(func() {
		GinkgoT().Logf("Post suite setup...")
		switch testEnv {
		case "PKS":
			err := testCluster.DeleteCluster()
			Expect(err).To(BeNil())
			Eventually(checkAdapterIsDeleted, 120*time.Second, 5*time.Second).Should(BeTrue())
		case "kubo":
			err := revertToOriginalManifest()
			Expect(err).NotTo(HaveOccurred())
		default:
			GinkgoT().Fatalf("\nunknown testEnv: %s\n", testEnv)
		}
		if client != nil {
			client.Logout(context.Background())
		}
	})
})
