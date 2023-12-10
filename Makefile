internal-tests:
	ginkgo -r -race -randomize-all -randomize-suites "./cmd/internal/"

config-tests:
	ginkgo -r -race -randomize-all -randomize-suites "./cmd/env/"

all-tests:
	ginkgo -r -race -randomize-all -randomize-suites .
