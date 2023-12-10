package env_test

import (
	"od-task/cmd/env"
	"os"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	const (
		invalidPort = "invalid_port"
		port        = 8080
		portEnv     = "PORT"

		host    = "host"
		hostEnv = "HOST"

		DBHost    = "dbhost"
		DBHostEnv = "DB_HOST"

		DBPort    = 1234
		DBPortEnv = "DB_PORT"

		DBUser    = "dbuser"
		DBUserEnv = "DB_USER"

		DBPass    = "dbpass"
		DBPassEnv = "DB_PASSWORD"

		DBName    = "dbname"
		DBNameEnv = "DB_NAME"
	)

	When("all needed env variables are set properly", func() {
		var expectedConfig = env.AppConfig{
			Host:   host,
			Port:   port,
			DBHost: DBHost,
			DBPort: DBPort,
			DBUser: DBUser,
			DBPass: DBPass,
			DBName: DBName,
		}
		BeforeEach(func() {
			Expect(os.Setenv(hostEnv, host)).To(Succeed())
			Expect(os.Setenv(portEnv, strconv.Itoa(port))).To(Succeed())
			Expect(os.Setenv(DBHostEnv, DBHost)).To(Succeed())
			Expect(os.Setenv(DBPortEnv, strconv.Itoa(DBPort))).To(Succeed())
			Expect(os.Setenv(DBUserEnv, DBUser)).To(Succeed())
			Expect(os.Setenv(DBPassEnv, DBPass)).To(Succeed())
			Expect(os.Setenv(DBNameEnv, DBName)).To(Succeed())
		})

		AfterEach(func() {
			Expect(os.Unsetenv(hostEnv)).To(Succeed())
			Expect(os.Unsetenv(portEnv)).To(Succeed())
			Expect(os.Unsetenv(DBHostEnv)).To(Succeed())
			Expect(os.Unsetenv(DBPortEnv)).To(Succeed())
			Expect(os.Unsetenv(DBUserEnv)).To(Succeed())
			Expect(os.Unsetenv(DBPassEnv)).To(Succeed())
			Expect(os.Unsetenv(DBNameEnv)).To(Succeed())

		})

		It("should bind to AppConfig properly", func() {
			appConfig, err := env.LoadAppConfig()
			Expect(err).ToNot(HaveOccurred())
			Expect(appConfig).To(Equal(expectedConfig))
		})
	})

	When("there is error in env variable", func() {
		BeforeEach(func() {
			Expect(os.Setenv(hostEnv, host)).To(Succeed())
			Expect(os.Setenv(portEnv, invalidPort)).To(Succeed())
			Expect(os.Setenv(DBHostEnv, DBHost)).To(Succeed())
			Expect(os.Setenv(DBPortEnv, strconv.Itoa(DBPort))).To(Succeed())
			Expect(os.Setenv(DBUserEnv, DBUser)).To(Succeed())
			Expect(os.Setenv(DBPassEnv, DBPass)).To(Succeed())
			Expect(os.Setenv(DBNameEnv, DBName)).To(Succeed())
		})

		AfterEach(func() {
			Expect(os.Unsetenv(hostEnv)).To(Succeed())
			Expect(os.Unsetenv(portEnv)).To(Succeed())
			Expect(os.Unsetenv(DBHostEnv)).To(Succeed())
			Expect(os.Unsetenv(DBPortEnv)).To(Succeed())
			Expect(os.Unsetenv(DBUserEnv)).To(Succeed())
			Expect(os.Unsetenv(DBPassEnv)).To(Succeed())
			Expect(os.Unsetenv(DBNameEnv)).To(Succeed())
		})

		It("should return an empty app config and an error", func() {

			appConfig, err := env.LoadAppConfig()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("failed to parse configuration from environment"))
			Expect(appConfig).To(Equal(env.AppConfig{}))
		})
	})
})
